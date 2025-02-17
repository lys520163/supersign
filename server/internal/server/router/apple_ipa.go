package router

import (
	"supersign/internal/api"
	v1 "supersign/internal/api/v1"
	"supersign/internal/model"
	"supersign/internal/server/middleware"
	"supersign/pkg/log"

	"github.com/gin-gonic/gin"
)

func registerAppleIPARouter(store *model.Store, r *gin.RouterGroup) {
	appleIPA := v1.AppleIPA{
		Base: api.New(store, log.New("AppleIPA").L()),
	}
	appleIPAR := r.Group("/ipa", middleware.JWT())

	{
		appleIPAR.POST("", appleIPA.Upload)
		appleIPAR.GET("", appleIPA.List)
		appleIPAR.DELETE("", appleIPA.Del)
	}
}
