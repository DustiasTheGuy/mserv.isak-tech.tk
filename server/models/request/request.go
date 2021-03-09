package request

import (
	"log"
	"paste/database"
	"time"
)

type Request struct {
	ID      int64     `json:"id"`
	Created time.Time `json:"created"`
	IP      int64     `json:"ip"`
	Href    string    `json:"href"`
}

func (req *Request) SaveRequest() error {
	db, err := database.Connect("isak_tech_paste")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	return nil
}

func GetAllRequests() []Request {
	db, err := database.Connect("isak_tech_paste")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	return nil
}

func (req *Request) GetSingleRequest() *Request {
	db, err := database.Connect("isak_tech_paste")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	return nil
}
