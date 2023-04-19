package models

type User struct {
	ID        int64  `json:id`
	Name      string `json:name`
	Email     string `json:email`
	Password  string `json:password`
	createdAt int32  `json:create_at`
	updatedAt int32  `json:updated_at`
}

type UserLogin struct {
	Email    string `json:email`
	Password string `json:password`
}
