package schemas

type Create_books struct {
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Price       float32 `json:"price"`
	About       string  `json:"about"`
	Category    int     `json:"category"`
	Language_id int     `json:"language_id"`
}
