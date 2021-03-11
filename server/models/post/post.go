package post

import (
	"database/sql"
	"fmt"
	"math"
	"paste/database"
	"time"
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

// Post is a struct for dealing with new Post
type Post struct {
	ID    int64     `json:"id"`
	Title string    `json:"title"`
	Body  string    `json:"body"`
	Date  time.Time `json:"date"`
	IP    string    `json:"-"`
	Tags  []string  `json:"tags"`
}

// save a new post to the database
// first return value defaults to 0 if an error has occured
func (post *Post) SavePost() (int64, error) {
	db, err := database.Connect("isak_tech_paste")

	if err != nil {
		return 0, err
	}

	defer db.Close()

	res, err := db.Exec(
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
func GetPost(ID int64) (Post, error) {
	var post Post

	db, err := database.Connect("isak_tech_paste")

	if err != nil {
		return post, err
	}

	row := db.QueryRow("SELECT * FROM posts WHERE ID=?", ID)

	if err := row.Scan(&post.ID, &post.Body, &post.Date, &post.IP, &post.Title); err != nil {
		return post, err
	}

	tags, err := post.GetTags()

	if err != nil {
		return post, err
	}

	post.Tags = tags
	return post, nil
}

// GetPosts from database
func GetPosts() ([]Post, error) {
	db, err := database.Connect("isak_tech_paste")

	if err != nil {
		return nil, err
	}

	defer db.Close()

	var Posts []Post

	rows, err := db.Query("SELECT * FROM posts")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post Post

		// for each row, scan the result into our tag composite object
		if err := rows.Scan(&post.ID, &post.Body, &post.Date, &post.IP, &post.Title); err != nil {
			return nil, err
		}

		tags, err := post.GetTags()

		if err != nil {
			return nil, err
		}

		post.Tags = tags
		Posts = append(Posts, post)
	}

	return Posts, nil
}

func (post *Post) GetTags() ([]string, error) {
	db, err := database.Connect("isak_tech_paste")

	if err != nil {
		return nil, err
	}

	defer db.Close()

	var tags []string

	rows, err := db.Query("SELECT * FROM tags WHERE postid=? ORDER BY id DESC LIMIT 5", post.ID)

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

func (post *Post) InsertTags() error {
	db, err := database.Connect("isak_tech_paste")

	if err != nil {
		return err
	}

	defer db.Close()

	var sqlString string = "INSERT INTO tags (tag, postid) VALUES"

	for i := 0; i < len(post.Tags); i++ {
		if i == len(post.Tags)-1 {
			sqlString += fmt.Sprintf("('%s', %d);", post.Tags[i], post.ID)
		} else {
			sqlString += fmt.Sprintf("('%s', %d), ", post.Tags[i], post.ID)
		}
	}

	_, err = db.Exec(sqlString)

	if err != nil {
		return err
	}

	return nil
}

func (post *Post) DeleteOne() error {
	db, err := database.Connect("isak_tech_paste")

	if err != nil {
		return err
	}

	defer db.Close()

	r, err := db.Exec("DELETE FROM posts WHERE id=?", post.ID)

	if err != nil {
		return err
	}

	n, _ := r.RowsAffected()

	fmt.Printf("Rows Affected: %d\n", n)

	return nil
}

func (post *Post) UpdateOne() error {
	db, err := database.Connect("isak_tech_paste")

	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec(
		"UPDATE posts SET title=?, body=? WHERE id=?",
		post.Title, post.Body, post.ID)

	if err != nil {
		return err
	}

	if err := post.InsertTags(); err != nil {
		return err
	}

	return nil
}

func Paginate(limit, start int64) (map[string]interface{}, error) {
	db, err := database.Connect("isak_tech_paste")

	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM posts LIMIT ? OFFSET ?", limit, start)

	if err != nil {
		return nil, err
	}

	row, err := db.Query("SELECT COUNT(*) FROM posts")

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

		tags, err := post.GetTags()

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
