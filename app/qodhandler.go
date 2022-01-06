package app

import (
	"fmt"
	"net/http"

	"github.com/adityapwr/go-banking/domain"
)

type Quote struct {
	Author     string
	Quote      string
	Tags       []string
	Id         string
	Background string
	Length     string
	Category   string
	Language   string
	Date       string
	Permalink  string
	Title      string
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

type QoDHandler struct {
	repository domain.QoDRepository
}

func (qh *QoDHandler) GetQoD(w http.ResponseWriter, r *http.Request) {
	qod, err := qh.repository.GetQod()
	fmt.Println(qod)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, nil)
	} else {
		writeResponse(w, http.StatusOK, qod)
	}
}
