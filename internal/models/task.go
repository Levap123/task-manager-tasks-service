package models

type Task struct {
	ID     string `bson:"_id,omitempty"`
	Title  string `bson:"title,omitempty"`
	Body   string `bson:"body,omitempty"`
	UserID int64  `bson:"user_id,omitempty"`
}
