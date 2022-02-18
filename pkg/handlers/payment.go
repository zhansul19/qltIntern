package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	finance "github.com/zhansul19/qltIntern"
)

func (h *Handler) CreatePayments(c *gin.Context) {
	categoryId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input finance.Payments
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	dohod := "доход"
	rashod := "расход"
	if input.Type == dohod || input.Type == rashod {
		id, err := h.services.Payment.Create(categoryId, input)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})
	} else {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Type only can be %s or %s ", dohod, rashod))
		return
	}

}

func (h *Handler) GetPayments(c *gin.Context) {
	categoryId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	payments, err := h.services.Payment.GetAll(categoryId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, payments)
}

func (h *Handler) GetPaymentsById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	payment, err := h.services.Payment.GetPaymentById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, payment)
}

func (h *Handler) DeletePayments(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Payment.DeletePaymentById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "Payment deleted",
	})
}
func (h *Handler) UpdatePayments(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var input finance.UpdatePaymentInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	dohod := "доход"
	rashod := "расход"
	if input.Type == dohod || input.Type == rashod {
		err := h.services.Payment.UpdatePayments(id, input)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, statusResponse{
			Status: "Payment Updated",
		})
	} else {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Type only can be %s or %s ", dohod, rashod))
		return
	}
}
