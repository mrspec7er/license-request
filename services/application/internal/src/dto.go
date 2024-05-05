package src

type ChangeStatusInput struct {
	Number string `json:"number"`
	Status string `json:"status"`
	Note   string `json:"note"`
}
