package orders

type SaveOrderInput struct {
	CustomerName string          `json:"customerName" binding:"required"`
	Items        []SaveItemInput `json:"items" binding:"required"`
}

type SaveItemInput struct {
	Code        string `json:"itemCode" binding:"required"`
	Description string `json:"description" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
}

type UpdateOrderInput struct {
	CustomerName string            `json:"customerName" binding:"required"`
	Items        []UpdateItemInput `json:"items" binding:"required"`
}

type UpdateItemInput struct {
	ID          int    `json:"lineItemId" binding:"required"`
	Code        string `json:"itemCode" binding:"required"`
	Description string `json:"description" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
}

type FindOrderInput struct {
	ID int `uri:"id" binding:"required"`
}
