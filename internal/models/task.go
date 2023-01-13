package models

type Task struct {
	ID     int    `bson:"_id,omitempty"`
	Title  string `bson:"title,omitempty"`
	Body   string `bson:"body,omitempty"`
	UserID int    `bson:"user_id,omitempty"`
}
