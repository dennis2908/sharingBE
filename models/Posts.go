package models

import "time"

type Posts struct {
	Id           int    `orm:"auto;pk;index"`
	Title        string `orm:"size(200)"`
	Content      string
	Category     string    `orm:"size(100)"`
	Created_date time.Time `orm:"auto_now;type(datetime)"`
	Updated_date time.Time `orm:"auto_now;type(datetime)"`
	Status       string    `orm:"size(100)"`
}

func (a *Posts) TableName() string {
	return "posts"
}
