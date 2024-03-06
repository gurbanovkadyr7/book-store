package repositories

import (
	db "admin/database"
	"admin/src/queries"
	request "admin/src/schemas/request"
	response "admin/src/schemas/response"
	"context"
	"log"
)

func Create_book(book request.Create_books) error {
	var id int
	db.DB.QueryRow(
		context.Background(), queries.Create_book,
		book.Title,
		book.Author,
		book.Price,
		book.About,
		book.Category,
		book.Language_id,
	).Scan(&id)

	_, err := db.DB.Exec(context.Background(), queries.Many_to_Many,
		book.Category, id,
	)

	if err != nil {
		return err
	}
	return nil

}

func BookByName(book response.Book_ByName) ([]response.Books, error) {
	var books []response.Books
	rows, err := db.DB.Query(context.Background(),
		queries.Get_books_by_name,
		book.Title,
	)

	if err != nil {
		log.Println("Error querying")
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var book response.Books
		err := rows.Scan(

			&book.Title,
			&book.Author,
			&book.Price,
			&book.About,
			&book.Category,
		)
		if err != nil {
			log.Println("Error in for")
			return nil, err
		}

		books = append(books, book)
	}

	if err != nil {
		log.Println("Error in after for")
		return nil, err
	}

	return books, err
}

func BookByAuthor(author response.Book_ByAuthor) ([]response.Books, error) {
	var books []response.Books

	rows, err := db.DB.Query(context.Background(),
		queries.Get_books_by_author,
		author.Author,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var book response.Books
		err := rows.Scan(

			&book.Title,
			&book.Author,
			&book.Price,
			&book.About,
			&book.Category,
		)
		if err != nil {
			log.Println("Error in for")
			return nil, err
		}

		books = append(books, book)
	}

	if err != nil {

		return nil, err
	}

	return books, err

}
func BookByCategory(category response.Book_ByCategory) ([]response.Books, error) {
	var books []response.Books

	rows, err := db.DB.Query(context.Background(), queries.Get_by_category,
		category.Category,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var book response.Books
		err := rows.Scan(

			&book.Title,
			&book.Author,
			&book.Price,
			&book.About,
			&book.Category,
		)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	if err != nil {

		return nil, err
	}

	return books, err

}
