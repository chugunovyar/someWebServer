package core

import (
	"fmt"
	"strings"
	"time"
)

type PythonDateTime struct {
	time.Time
}

const expiryDateLayout = "2006-01-02 15:04:05"

func (ct *PythonDateTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(expiryDateLayout, s)
	return
}

func (ct *PythonDateTime) MarshalJSON() ([]byte, error) {
	if ct.Time.IsZero() {
		return nil, nil
	}
	return []byte(fmt.Sprintf(`"%s"`, ct.Time.Format(expiryDateLayout))), nil
}

type Article struct {
	PubDate  PythonDateTime `json:"pub_date"`
	Headline string         `json:"headline"`
	Content  string         `json:"content"`
}

type Page struct {
	Title string
	Body  string
}
