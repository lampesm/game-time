package forms

type GamePayload struct {
	Name        string `json:"name" xml:"name" binding:"required"`
	Description string `json:"description" xml:"description" binding:"required"`
	Image       string `json:"image" xml:"image" binding:"required"`
	Cover       string `json:"cover" xml:"cover" binding:"required"`
	Category    int    `json:"category" xml:"category" binding:"required"`
}
