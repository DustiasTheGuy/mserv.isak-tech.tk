package api

import (
	"database/sql"
	"paste/database"
)

// SQL struct contains the state of an SQL connection attempt
type SQL struct {
	Error       error
	Established bool
	Connection  *sql.DB
}

// CreateConnection attempts to connect to mysql
func CreateConnection() *SQL {
	db, err := database.Connect()

	if err != nil {
		return &SQL{
			Error:       err,
			Established: false,
			Connection:  nil,
		}
	}

	return &SQL{
		Error:       nil,
		Established: true,
		Connection:  db,
	}
}

// save a new post to the database
// first return value defaults to 0 if an error has occured
func (sql *SQL) savePost(body *Post) (int64, error) {
	res, err := sql.Connection.Exec("INSERT INTO posts (body, ip) VALUES (?, ?)", body.Body, body.IP)

	if err != nil {
		return 0, err
	}

	lastID, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	defer sql.Connection.Close()
	return lastID, nil
}

// GetPosts from database
func (sql *SQL) getPost(ID int64) (Post, error) {
	var post Post

	db, err := database.Connect()

	if err != nil {
		return post, err
	}

	row := db.QueryRow("SELECT * FROM posts WHERE ID=?", ID)

	if err := row.Scan(&post.ID, &post.Body, &post.Date, &post.IP); err != nil {
		return post, err
	}

	return post, nil
}

// GetPosts from database
func (sql *SQL) getPosts() ([]Post, error) {
	var Posts []Post

	rows, err := sql.Connection.Query("SELECT * FROM posts")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post Post

		// for each row, scan the result into our tag composite object
		err = rows.Scan(&post.ID, &post.Body, &post.Date, &post.IP)

		if err != nil {
			return nil, err
		}

		Posts = append(Posts, post)
	}

	defer sql.Connection.Close()

	return Posts, nil
}
