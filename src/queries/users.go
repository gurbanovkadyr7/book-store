package queries

var Create_users = "insert into users (first_name, last_name, email,phone,hash_pass) values ($1, $2, $3, $4, $5)"

var Get_users = "select * from users "

var Compare = "select first_name, hash_pass from users where first_name = $1"
