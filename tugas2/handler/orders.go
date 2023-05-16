package handler

import (
	"tugas2/orders"

	"github.com/gin-gonic/gin"
)

type ordersHandler struct {
	service orders.Service
}

func NewOrdersHandler(service orders.Service) *ordersHandler {
	return &ordersHandler{service}
}

func (h *ordersHandler) FindAll(c *gin.Context) {
	orders, err := h.service.FindAll()
	if err != nil {
		c.JSON(500, gin.H{"message": "error"})
		return
	}

	c.JSON(200, orders)
}

func (h *ordersHandler) FindByID(c *gin.Context) {
	var input orders.FindOrderInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		c.JSON(400, gin.H{"message": "error"})
		return
	}

	order, err := h.service.FindByID(input)
	if err != nil {
		c.JSON(500, gin.H{"message": "error"})
		return
	}

	c.JSON(200, order)
}

func (h *ordersHandler) Save(c *gin.Context) {
	var input orders.SaveOrderInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"message": "error"})
		return
	}

	order, err := h.service.Save(input)
	if err != nil {
		c.JSON(500, gin.H{"message": "error"})
		return
	}

	c.JSON(200, order)
}

func (h *ordersHandler) Update(c *gin.Context) {
	var inputID orders.FindOrderInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	var inputData orders.UpdateOrderInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		c.JSON(400, gin.H{"message": "error"})
		return
	}

	order, err := h.service.Update(inputID, inputData)
	if err != nil {
		c.JSON(500, gin.H{"message": "error"})
		return
	}

	c.JSON(200, order)
}

func (h *ordersHandler) Delete(c *gin.Context) {
	var input orders.FindOrderInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		c.JSON(400, gin.H{"message": "error"})
		return
	}

	order, err := h.service.Delete(input)
	if err != nil {
		c.JSON(500, gin.H{"message": "error"})
		return
	}

	c.JSON(200, order)
}
