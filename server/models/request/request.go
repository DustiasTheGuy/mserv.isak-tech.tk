package request

import (
	"paste/database"
	"time"
)

type Request struct {
	ID      int64     `json:"id"`
	Created time.Time `json:"created"`
	IP      string    `json:"ip"`
	Href    string    `json:"href"`
}

func (req *Request) SaveRequest() error {
	db, err := database.Connect("isak_tech_analytics")

	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("INSERT INTO requests (ip, href) VALUES (?, ?)",
		req.IP, req.Href)

	return err
}

func GetAllRequests() []Request {
	var requests []Request
	db, err := database.Connect("isak_tech_analytics")

	if err != nil {
		return nil
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM requests")

	if err != nil {
		return nil
	}

	for rows.Next() {
		var request Request

		if err := rows.Scan(
			&request.ID,
			&request.Created,
			&request.IP,
			&request.Href,
		); err != nil {
			return nil
		}

		requests = append(requests, request)
	}

	return requests
}

func (req *Request) GetSingleRequest() *Request {
	db, err := database.Connect("isak_tech_analytics")

	if err != nil {
		return nil
	}

	defer db.Close()

	return &Request{}
}
