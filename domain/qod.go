package domain

import "github.com/adityapwr/banking-lib/errs"

type QoDRepository interface {
	GetQod() (QoD, *errs.AppError)
}

type RemoteQoDRepository struct {
}

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

func (r RemoteQoDRepository) GetQod() (QoD, *errs.AppError) {
	qod := QoD{
		Sucess: "true",
		Content: content{
			Quotes: []Quote{
				{
					Author: "Aditya",
					Quote:  "I am a programmer",
					Tags:   []string{"programming", "go"},
					Id:     "1",
					Image:  "https://cdn.pixabay.com/photo/2016/03/09/09/43/go-lang-1209823_960_720.png",
					Length: 1,
				},
			},
		},
	}
	return qod, nil

}

func NewQoDRepository() RemoteQoDRepository {
	return RemoteQoDRepository{}
}
