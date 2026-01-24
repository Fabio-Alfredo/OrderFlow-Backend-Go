package entities

type User struct {
	Id       string `gorm:"primary_key;column:id"`
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
	Status   string `gorm:"column:status"`
}

func (User) TableName() string { return "users_tb" }
