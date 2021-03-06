package model

import (
	"time"
)

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `json:"-"`
}

type Teacher struct {
	Model
	UserID   uint    `json:"userId"`
	Nid      *string `gorm:"unique" json:"nid"`
	Fullname *string `json:"fullname"`
	Country  *string `json:"country"`
	State    *string `json:"state"`
	City     *string `json:"city"`
	Phone    *string `json:"phone"`
	Email    *string `json:"email" gorm:"unique"`
	User     User    `json:"-"`
}

type TeacherBTUser struct {
	Teacher
	User User `json:"user"`
}

func (TeacherBTUser) TableName() string {
	return "teachers"
}

/* -------------------------------------------------------------------------- */
/*                                  REGISTER                                  */
/* -------------------------------------------------------------------------- */

// TeacherRegisterParams ...
type TeacherRegisterParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
}
