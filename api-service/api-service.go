package api_service

import (
	lhttp "net/http"

	lgin "github.com/gin-gonic/gin"

	lsconfig "github.com/vngcloud/reforged-labs/configuration/api-service"
)

type APIService struct {
	router     *lgin.Engine
	httpServer *lhttp.Server
	config     *lsconfig.APIServiceConfiguration
}
