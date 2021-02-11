package api

import "mserv/database"

// GetPosts from database
func GetPosts() ([]Post, error) {
	var Posts []Post
	db, err := database.Connect()

	rows, err := db.Query("SELECT * FROM posts")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post Post

		// for each row, scan the result into our tag composite object
		err = rows.Scan(&post.ID, &post.Body, &post.Date)

		if err != nil {
			return nil, err
		}

		Posts = append(Posts, post)
	}

	return Posts, nil
}
