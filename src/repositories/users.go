package repositories

import (
	"admin/config"
	db "admin/database"
	"admin/src/queries"
	request "admin/src/schemas/request"
	"admin/src/tools"
	"fmt"
	"log"
	"time"

	response "admin/src/schemas/response"
	"context"

	"golang.org/x/crypto/bcrypt"
)

func Get() ([]response.Users, error) {
	var users []response.Users

	rows, err := db.DB.Query(
		context.Background(), queries.Get_users,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user response.Users
		err := rows.Scan(

			&user.ID,
			&user.First_name,
			&user.Last_name,
			&user.Email,
			&user.Phone,
			&user.Hash_pass,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err != nil {
		return nil, err
	}

	return users, nil
}

func Create_user(user request.Create_users) error {
	hashed_password, _ := bcrypt.GenerateFromPassword([]byte(user.Hash_pass),
		bcrypt.DefaultCost)

	_, err := db.DB.Exec(
		context.Background(), queries.Create_users,
		user.First_name,
		user.Last_name,
		user.Email,
		user.Phone,
		hashed_password,
	)

	if err != nil {
		return err
	}

	return nil
}

func Check_User(user response.User) (request.Finded_user, error) {
	var users request.Finded_user

	err := db.DB.QueryRow(context.Background(),
		queries.Compare, user.First_name,
	).Scan(
		&users.First_name,
		&users.Hash_pass,
	)

	err = bcrypt.CompareHashAndPassword([]byte(users.Hash_pass),
		[]byte(user.Hash_pass))

	if err != nil {
		return users, err
	}
	fmt.Println(err)

	return users, err

}

func Create_token(user request.Finded_user) request.Tokens {
	access_time := time.Now().Add(config.ENV.ACCESS_TIME).Unix()
	refresh_time := time.Now().Add(config.ENV.REFRESH_TIME).Unix()

	var token request.Tokens

	token.Access_token = tools.Create_token(
		user.First_name, user.Hash_pass, access_time, config.ENV.ACCESS_KEY,
	)

	token.Refresh_token = tools.Create_token(
		user.First_name, user.Hash_pass, refresh_time, config.ENV.REFRESH_KEY,
	)

	token.First_name = user.First_name

	return token

}

func Get_book() ([]response.Books, error) {
	var books []response.Books

	rows, err := db.DB.Query(context.Background(),
		queries.Get_books,
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
