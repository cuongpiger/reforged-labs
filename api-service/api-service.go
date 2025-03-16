package api_service

import (
	lctx "context"
	lfmt "fmt"
	lhttp "net/http"

	lgin "github.com/gin-gonic/gin"
	lzap "go.uber.org/zap"
	lgorm "gorm.io/gorm"

	lsconfig "github.com/cuongpiger/reforged-labs/configuration/api-service"
	lspostgres "github.com/cuongpiger/reforged-labs/infra/postgres"
	lsmdw "github.com/cuongpiger/reforged-labs/middleware"
	lsmdl "github.com/cuongpiger/reforged-labs/models"
	lsadshdl "github.com/cuongpiger/reforged-labs/services/domain/advertisement/delivery/http"
	lsadsuc "github.com/cuongpiger/reforged-labs/services/domain/advertisement/usecase"
	lsrepo "github.com/cuongpiger/reforged-labs/services/repository"
)

func NewAPIService(pconfig *lsconfig.APIServiceConfiguration) (*APIService, error) {
	router := lgin.New()
	return &APIService{
		router: router,
		config: pconfig,
	}, nil
}

type APIService struct {
	router     *lgin.Engine
	httpServer *lhttp.Server
	config     *lsconfig.APIServiceConfiguration
}

func (s *APIService) WarmUp() error {
	ctx := lctx.Background()

	// Configure the database
	db, err := s.setupDatabase(s.config.APIService.Database.URI)
	if err != nil {
		lzap.L().Error("Failed to configure database", lzap.Error(err))
		return err
	}

	repo := lsrepo.NewRepository(db)
	domains := s.setupDomains(ctx, repo)

	s.setupMiddlewares()
	s.setupRoutes(domains)
	s.setupHealthCheckRoute()
	return nil
}

func (s *APIService) Stop() error {
	lzap.L().Info("Stop the server")
	return nil
}

func (s *APIService) ServeHTTPService() error {
	address := lfmt.Sprintf("%s:%d", s.config.APIService.Host, s.config.APIService.Port)
	s.httpServer = &lhttp.Server{
		Handler: s.router,
		Addr:    address,
	}

	lzap.S().Infof("The API service is running on %s", address)
	return s.httpServer.ListenAndServe()
}

func (s *APIService) setupMiddlewares() {
	// Add middlewares here
	s.router.Use(lsmdw.GenerateRequestID())
}

func (s *APIService) setupRoutes(pdomains *Domains) {
	// Add routes here

	// The group of API v1
	apiV1Group := s.router.Group("api/v1")
	lsadshdl.NewAdvertisementHandler(pdomains.advertisement).Route(apiV1Group.Group("ads")) // ads
}

func (s *APIService) setupHealthCheckRoute() {
	s.router.GET("/healthz", func(pctx *lgin.Context) {
		pctx.JSON(lhttp.StatusOK, lgin.H{
			"status": "ok",
		})
	})
}

func (s *APIService) setupDomains(pctx lctx.Context, prepo lsrepo.IRepository) *Domains {
	return &Domains{
		advertisement: lsadsuc.NewAdvertisementUseCase(prepo),
	}
}

func (s *APIService) setupDatabase(puri string) (*lgorm.DB, error) {
	client, err := lspostgres.InitPostgreSQL(puri)
	if err != nil {
		return nil, err
	}

	if err = client.AutoMigrate(&lsmdl.Advertisement{}); err != nil {
		lzap.L().Error("Failed to auto migrate Advertisement model", lzap.Error(err))
		return nil, err
	}

	return client, nil
}
