package schemas

type Users struct {
	ID         int    `json:"id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Hash_pass  string `json:"hash_pass"`
}

type User struct {
	First_name string `json:"first_name"`
	Hash_pass  string `json:"hash_pass"`
}
