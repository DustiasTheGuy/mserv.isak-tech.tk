package api

import (
	"database/sql"
	"fmt"
	"math"
	"paste/database"
)

// SQL struct contains the state of an SQL connection attempt
type SQL struct {
	Error       error
	Established bool
	Connection  *sql.DB
}

type Tag struct {
	ID     int64
	Tag    string
	PostID int64
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
func (sql *SQL) savePost(post *Post) (int64, error) {
	fmt.Printf("%s\n", post.Title)
	fmt.Printf("%s\n", post.Body)
	fmt.Printf("%s\n", post.IP)

	res, err := sql.Connection.Exec(
		"INSERT INTO posts (title, body, ip) VALUES (?, ?, ?)",
		post.Title, post.Body, post.IP)

	if err != nil {
		return 0, err
	}

	lastID, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

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

	if err := row.Scan(&post.ID, &post.Body, &post.Date, &post.IP, &post.Title); err != nil {
		return post, err
	}

	tags, err := sql.getTags(ID)

	if err != nil {
		return post, err
	}

	post.Tags = tags
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
		if err := rows.Scan(&post.ID, &post.Body, &post.Date, &post.IP, &post.Title); err != nil {
			return nil, err
		}

		tags, err := sql.getTags(post.ID)

		if err != nil {
			return nil, err
		}

		post.Tags = tags
		Posts = append(Posts, post)
	}

	return Posts, nil
}

func (sql *SQL) getTags(postID int64) ([]string, error) {
	var tags []string

	rows, err := sql.Connection.Query("SELECT * FROM tags WHERE postid=? ORDER BY id DESC LIMIT 5", postID)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tag Tag

		if err := rows.Scan(&tag.ID, &tag.Tag, &tag.PostID); err != nil {
			return nil, err
		}

		tags = append(tags, tag.Tag)
	}

	return tags, nil
}

func (sql *SQL) insertTags(postID int64, tags []string) error {
	var sqlString string = "INSERT INTO tags (tag, postid) VALUES"

	for i := 0; i < len(tags); i++ {
		if i == len(tags)-1 {
			sqlString += fmt.Sprintf("('%s', %d);", tags[i], postID)
		} else {
			sqlString += fmt.Sprintf("('%s', %d), ", tags[i], postID)
		}
	}

	_, err := sql.Connection.Exec(sqlString)

	if err != nil {
		return err
	}

	return nil
}

func (sql *SQL) deleteOne(postID int64) error {
	r, err := sql.Connection.Exec("DELETE FROM posts WHERE id=?", postID)

	if err != nil {
		return err
	}

	n, _ := r.RowsAffected()

	fmt.Printf("Rows Affected: %d\n", n)

	return nil
}

func (sql *SQL) updateOne(post *Post) error {

	_, err := sql.Connection.Exec(
		"UPDATE posts SET title=?, body=? WHERE id=?",
		post.Title, post.Body, post.ID)

	if err != nil {
		return err
	}

	if err := sql.insertTags(post.ID, post.Tags); err != nil {
		return err
	}

	return nil
}

func (sql *SQL) paginate(limit, start int64) (map[string]interface{}, error) {
	rows, err := sql.Connection.Query("SELECT * FROM posts LIMIT ? OFFSET ?", limit, start)

	if err != nil {
		return nil, err
	}

	row, err := sql.Connection.Query("SELECT COUNT(*) FROM posts")

	if err != nil {
		return nil, err
	}

	var count int64

	for row.Next() {
		if err := row.Scan(&count); err != nil {
			return nil, err
		}
	}

	var posts []Post
	for rows.Next() {
		var post Post

		if err := rows.Scan(&post.ID, &post.Body, &post.Date, &post.IP, &post.Title); err != nil {
			return nil, err
		}

		tags, err := sql.getTags(post.ID)

		if err != nil {
			return nil, err
		}

		post.Tags = tags
		posts = append(posts, post)
	}

	return map[string]interface{}{
		"posts": posts,
		"count": count,
		"pages": math.Ceil(float64(count / limit)),
	}, nil
}
