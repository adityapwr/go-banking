package domain

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/adityapwr/banking-lib/errs"
)

type QoDRepository interface {
	GetQod() (QoD, *errs.AppError)
}

type RemoteQoDRepository struct {
}

type QoD struct {
	Success struct {
		Total int `json:"total"`
	} `json:"success"`
	Contents struct {
		Quotes []struct {
			Quote  string `json:"quote"`
			Length string `json:"length"`
			Author string `json:"author"`
			Tags   struct {
				Num0 string `json:"0"`
				Num1 string `json:"1"`
				Num2 string `json:"2"`
				Num3 string `json:"3"`
				Num5 string `json:"5"`
			} `json:"tags"`
			Category   string `json:"category"`
			Language   string `json:"language"`
			Date       string `json:"date"`
			Permalink  string `json:"permalink"`
			ID         string `json:"id"`
			Background string `json:"background"`
			Title      string `json:"title"`
		} `json:"quotes"`
	} `json:"contents"`
	Baseurl   string `json:"baseurl"`
	Copyright struct {
		Year int    `json:"year"`
		URL  string `json:"url"`
	} `json:"copyright"`
}

func (r RemoteQoDRepository) GetQod() (QoD, *errs.AppError) {
	// qod := QoD{
	// 	Sucess: "true",
	// 	Content: content{
	// 		Quotes: []Quote{
	// 			{
	// 				Author: "Aditya",
	// 				Quote:  "I am a programmer",
	// 				Tags:   []string{"programming", "go"},
	// 				Id:     "1",
	// 				Image:  "https://cdn.pixabay.com/photo/2016/03/09/09/43/go-lang-1209823_960_720.png",
	// 				Length: 1,
	// 			},
	// 		},
	// 	},
	// }
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://quotes.rest/qod?language=en", nil)
	if err != nil {
		log.Fatalln(err)
		return QoD{}, errs.NewUnexpectedError(err.Error())

	}
	req.Header.Set("accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return QoD{}, errs.NewUnexpectedError(err.Error())
	}
	defer resp.Body.Close()
	log.Println(resp.Body)
	qod := QoD{}

	if err := json.NewDecoder(resp.Body).Decode(&qod); err != nil {
		log.Fatalln(err)
		return QoD{}, errs.NewUnexpectedError(err.Error())
	}
	return qod, nil

}

func NewQoDRepository() RemoteQoDRepository {
	return RemoteQoDRepository{}
}
