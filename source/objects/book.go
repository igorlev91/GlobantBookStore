package objects

type Book struct {
	BookID  uint    `json:"book_id" gorm:"primaryKey"`
	Name    string  `json:"name" gorm:"unique;size:100;not null;"`
	Price   float32 `json:"price" gorm:"not null"`
	GenreID uint    `json:"genre_id" gorm:"not null"`
	Amount  uint    `json:"amount" gorm:"not null"`
	Genre   Genre   `gorm:"foreignKey:GenreID"`
}

type Genre struct {
	Id   uint   `json:"id" gorm:"gorm:primaryKey"`
	Name string `json:"name" gorm:"unique;type:varchar(100);not null"`
}

const (
	BOOK  string = "Book"
	GENRE string = "Genre"
)
