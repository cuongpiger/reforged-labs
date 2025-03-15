package http

import lgin "github.com/gin-gonic/gin"

func (s *advertisementHandler) Route(prouter *lgin.RouterGroup) {
	prouter.POST("", s.createAdvertisement())
	prouter.GET("/:ad_id", s.getAdvertisement())
}
