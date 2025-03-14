package db

var Db []Book

type Book struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Author  Author  `json:"author"`
	Price   float32 `json:"price"`
	YearPub int     `json:"year_pub"`
}

type Author struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func init() {
	book1 := Book{
		ID:      1,
		Name:    "Война и мир",
		Author:  Author{Name: "Лев", Surname: "Толстой"},
		Price:   870.3,
		YearPub: 1869,
	}
	book2 := Book{
		ID:      2,
		Name:    "Преступление и наказание",
		Author:  Author{Name: "Фёдор", Surname: "Достоевский"},
		Price:   360.5,
		YearPub: 1870,
	}
	Db = append(Db, book1, book2)
}

func FindBookById(id int) (*Book, bool) {
	for _, book := range Db {
		if book.ID == id {
			return &book, true
		}
	}
	return &Book{}, false
}

func DeleteBookById(id int) bool {
	for _, v := range Db {
		if v.ID == id {
			Db = append(Db[:v.ID], Db[v.ID+1:]...)
			return true
		}
	}
	return false
}
