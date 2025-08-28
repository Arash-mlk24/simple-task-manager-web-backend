package entity

type Collection struct {
	BaseEntity
	Title string `json:"title"`
	Tasks []Task `json:"tasks"`
}
