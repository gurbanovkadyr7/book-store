package schemas

type Books struct {
	Title    string  `json:"title"`
	Author   string  `json:"author"`
	Price    float32 `json:"price"`
	About    string  `json:"about"`
	Category string  `json:"category"`
}

type Book_ByName struct {
	Title string `json:"title"`
}

type Book_ByAuthor struct {
	Author string `json:"author"`
}

type Book_ByCategory struct {
	Category string `json:"category"`
}
