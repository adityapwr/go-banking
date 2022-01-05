package app

import (
	"net/http"

	"github.com/adityapwr/go-banking/domain"
)

type Quote struct {
	Author string   `json:"author_name"`
	Quote  string   `json:"quote"`
	Tags   []string `json:"tags"`
	Id     string   `json:"id"`
	Image  string   `json:"image_url"`
	Length int      `json:"length"`
}

type content struct {
	Quotes []Quote `json:"quotes"`
}

type QoD struct {
	Sucess  string  `json:"success"`
	Content content `json:"content"`
}

type QoDHandler struct {
	repository domain.QoDRepository
}

func (qh *QoDHandler) GetQoD(w http.ResponseWriter, r *http.Request) {
	qod, err := qh.repository.GetQod()
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, nil)
	} else {
		writeResponse(w, http.StatusOK, qod)
	}
}
