package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ================= PAYMENT =================

// POST /api/v1/payments
func InitiatePayment(c *gin.Context) {
	var req interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	// TODO: call PaymentService.InitiatePayment()
	c.JSON(http.StatusCreated, nil)
}

// POST /api/v1/payments/:paymentId/refund
func RefundPayment(c *gin.Context) {
	paymentID := c.Param("paymentId")

	// TODO: call PaymentService.RefundPayment()
	_ = paymentID

	c.JSON(http.StatusOK, nil)
}
