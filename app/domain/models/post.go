package models

type Post struct {
	ID     int8   `json:"id"`
	Title  string `json:"title"`
	UserID int8   `json:"user_id"`
}
