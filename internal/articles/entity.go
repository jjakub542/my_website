package articles

type Image struct {
	Id          string
	Name        string
	Description string
}

type Article struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Public      bool   `json:"public"`
	Cover       Image
	Gallery     []Image
}
