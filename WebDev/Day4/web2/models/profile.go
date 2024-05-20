package models

type Profile struct {
	UserID string `json:"user_id"`
	Bio    string `json:"bio"`
	Age    int    `json:"age"`
}
