package models

type User struct {
	ID        int64  `gorm:id`
	Name      string `gorm:name`
	Email     string `gorm:email`
	Password  string `gorm:password`
	createdAt int32  `gorm:create_at`
	updatedAt int32  `gorm:updated_at`
}

type UserLogin struct {
	Email    string `gorm:email`
	Password string `gorm:password`
}
