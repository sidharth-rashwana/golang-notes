package model

type User struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Email    string `json:"email" gorm:"unique;no null" binding:"required"`
	Password string   `json:"password"`
}

func (User) TableName() string {
	return "user"
}
