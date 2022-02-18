package handlers

import (
	
	"github.com/zhansul19/qltIntern/pkg/service"

	"github.com/gin-gonic/gin"
)
type Handler struct{
	services *service.Service
}

func NewHandler(services *service.Service)*Handler{
	return &Handler{services: services}
}

func (h *Handler)InitRoutes()*gin.Engine{
	router := gin.New()


		category := router.Group("category")
		{
			category.POST("/", h.CreateCategory)
			category.GET("/", h.GetCategory)
			category.GET("/:id", h.GetCategoryById)
			category.DELETE("/:id", h.DeleteCategory)
			category.PUT("/:id", h.UpdateCategory)

			payments := category.Group(":id/payments")
			{
				payments.POST("/", h.CreatePayments)
				payments.GET("/", h.GetPayments)
			}
		}
		payments:=router.Group("payments")
		{
			payments.GET("/:id", h.GetPaymentsById)
			payments.DELETE("/:id", h.DeletePayments)
			payments.PUT("/:id", h.UpdatePayments)
		}
	
	return router
}
