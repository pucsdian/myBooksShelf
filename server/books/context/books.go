package context

import (
	"github.com/shubhamjagdhane/myBooksShelf/server/books/util"
)

// Migrate creates the schema of the tables
func (conn *Database) Migrate() {
	conn.DB.AutoMigrate(&Books{}, &UserCart{})
}

// GetUser gives us the user id by specifying username
func (conn *Database) GetBookByTitle(title string) (book Books, err error) {

	if err = conn.DB.Where("title = ?", title).First(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}

// GetDefaultBooks gives us books which containing the search_string
func (conn *Database) GetDefaultBooks(search_string string) (books []Books, err error) {
	search_string = "%" + search_string + "%"

	if err := conn.DB.Where("title LIKE ? or authors LIKE ?", search_string, search_string).Find(&books).Error; err != nil {
		return books, err
	}

	return books, nil
}

// GetBooksByOffset gives us books details with the mentioned offset
func (conn *Database) GetBooksByOffset(search_string string, offset int) (books []Books, err error) {
	search_string = "%" + search_string + "%"

	if err := conn.DB.Limit(52).Offset(offset).Where("title LIKE ? or authors LIKE ?", search_string, search_string).Find(&books).Error; err != nil {
		return books, err
	}

	return books, nil
}

func (conn *Database) GetRandomBooks() (books []Books, err error) {
	limit := 52
	offset := int(util.RandomInt(1, 9326))
	if err := conn.DB.Limit(limit).Offset(offset).Find(&books).Error; err != nil {
		return books, err
	}
	return books, nil
}

func (conn *Database) AddToCart(bookIDs []int, username string) (err error) {

	user := Users{}
	if err = conn.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return err
	}

	if err := conn.DB.Where("user_id = ?", user.ID).Delete(&UserCart{}).Error; err != nil {
		return err
	}

	for _, bookId := range bookIDs {
		if err := conn.DB.Create(&UserCart{BookID: uint64(bookId), UserID: user.ID}).Error; err != nil {
			return err
		}
	}

	return nil
}

func (conn *Database) GetCartItems(username string) (books []Books, err error) {
	user := Users{}
	if err = conn.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return books, err
	}

	var userCartID []UserCart
	if err := conn.DB.Where("user_id = ?", user.ID).Find(&userCartID).Error; err != nil {
		return books, err
	}

	for _, item := range userCartID {
		book := Books{}
		if err := conn.DB.Where("id = ?", item.BookID).First(&book).Error; err != nil {
			return books, err
		}
		books = append(books, book)
	}

	return books, nil
}
