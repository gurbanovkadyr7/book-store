package schemas

type Create_users struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Hash_pass  string `json:"hash_pass"`
}

type Finded_user struct {
	First_name string `json:"first_name"`
	Hash_pass  string `json:"hash_pass"`
}

type Tokens struct {
	Access_token  string `json:"access_token"`
	Refresh_token string `json:"refresh_token"`
	First_name    string `json:"first_name"`
}
