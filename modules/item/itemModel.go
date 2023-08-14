package item

type Item struct {
	Id    string `json:"id,omitempty"`
	Title string `json:"title" validate:"required"`
}
