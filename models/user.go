package models

import (
	"../auth"
	"time"
)



type User struct {
	Id        uint64    `json:"id,omitempty"`
	Nickname  string    `form:"nickname" json:"nickname,omitempty"`
	Password  string    `form:"password" json:"-"`
	Gender    int64     `json:"gender,omitempty"`
	Email	  string    `json:email,omitempty`
	Birthday  time.Time `json:"birthday,omitempty"`
	CreatedAt time.Time `gorm:"column:created_time" json:"created_time,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_time" json:"updated_time,omitempty"`

	authenticated bool `form:"-" db:"-" json:"-"`
}

func GenerateAnonymousUser() auth.User {
	return &User{}
}

func (u User) TableName() string {
	return "user"
}

// Login will preform any actions that are required to make a user model
// officially authenticated.
func (u *User) Login() {
	// Update last login time
	// Add to logged-in user's list
	// etc ...
	u.authenticated = true
}

// Logout will preform any actions that are required to completely
// logout a user.
func (u *User) Logout() {
	// Remove from logged-in user's list
	// etc ...
	u.authenticated = false
}

func (u *User) IsAuthenticated() bool {
	return u.authenticated
}

func (u *User) UniqueId() interface{} {
	return u.Id
}

// GetById will populate a user object from a database model with
// a matching id.
func (u *User) GetById(id interface{}) error {
	u.Id = id.(uint64)
	// err := dbmap.SelectOne(u, "SELECT * FROM users WHERE id = $1", id)
	// if err != nil {
	// 	return err
	// }

	return nil
}
