package todo

type Book struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	AuthorId int    `json:"authorid"`
	Year     int    `json:"year"`
	ISBN     string `json:"isbn"`
}

type Author struct {
	Id          int    `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Description string `json:"description"`
	Birthday    string `json:"birthday"`
}

type BookAuthor struct {
	Book   Book   `json:"book"`
	Author Author `json:"author"`
}
