package api_service

import (
	lctx "context"
	lfmt "fmt"
	lhttp "net/http"

	lgin "github.com/gin-gonic/gin"
	lzap "go.uber.org/zap"

	lsconfig "github.com/cuongpiger/reforged-labs/configuration/api-service"
	lsmdw "github.com/cuongpiger/reforged-labs/middleware"
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

	repo := lsrepo.NewRepository()
	domains := s.configureDomains(ctx, repo)

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
	s.router.Use(lsmdw.CheckUserIDHeader())
	s.router.Use(lsmdw.GenerateRequestID())
}

func (s *APIService) setupRoutes(pdomains *Domains) {
	// Add routes here
	//lsnghdl.NewNodeGroupHandler(pdomains.nodegroup).Route(s.router.Group("api/v1/nodegroups"))
}

func (s *APIService) setupHealthCheckRoute() {
	s.router.GET("/healthz", func(pctx *lgin.Context) {
		pctx.JSON(lhttp.StatusOK, lgin.H{
			"status": "ok",
		})
	})
}

func (s *APIService) configureDomains(pctx lctx.Context, prepo lsrepo.IRepository) *Domains {
	return &Domains{}
}
