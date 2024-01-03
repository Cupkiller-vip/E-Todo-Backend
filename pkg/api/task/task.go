package task

type CreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Level       uint   `json:"level" binding:"required"`
	Timestamp   uint   `json:"timestamp" binding:"required"`
}

type DeleteRequest struct {
	Id uint `json:"id" binding:"required"`
}

type EditRequest struct {
	Id          uint   `json:"id" binding:"required"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Level       uint   `json:"level"`
}

type ReadRequest struct {
	Id uint `json:"id" binding:"required"`
}
