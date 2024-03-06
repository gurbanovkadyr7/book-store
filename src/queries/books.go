package queries

var Create_book = "insert into books (title, author, price, about, category, language_id) values ($1, $2, $3, $4, $5, $6) returning id"

var Get_books = "select b.title, b.author, b.price, b.about, categories.title from books b join categories on categories.id=b.category"

var Get_books_by_name = "select b.title, b.author, b.price, b.about, categories.title from books b join categories on categories.id=b.category where b.title = $1 "

var Get_books_by_author = "select b.title, b.author, b.price, b.about, categories.title from books b join categories on categories.id=b.category where b.author = $1 "

var Many_to_Many = "insert into categories_books (category_id, book_id) values ($1, $2)"

var Get_by_category = "SELECT b.title, b.author, b.price, b.about, c.title AS category_name FROM books AS b INNER JOIN categories_books AS cb ON b.id = cb.book_id INNER JOIN categories AS c ON cb.category_id = c.id WHERE c.title = $1 GROUP BY b.id, b.title, b.author, b.price, b.about, c.title"
