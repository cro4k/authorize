package model

import "time"

type User struct {

	// columns cannot be edited

	ID        int64
	Username  string
	Secret    string
	Password  string
	UpdatedAt time.Time
	CreatedAt time.Time

	UserColumns // columns could be edited

	shadow *UserColumns // columns before edited
}

func (u *User) GetID() int64 { return u.ID }

// GetUpdate
// get the columns that only edited
func (u *User) GetUpdate() any {
	update := map[string]any{}
	if u.Nickname != u.shadow.Nickname {
		update["nickname"] = u.Nickname
	}
	return update
}

func (u *User) Rebase() {
	if u.shadow == nil {
		u.shadow = new(UserColumns)
	}
	*u.shadow = u.UserColumns
}

type UserColumns struct {
	Nickname string
}
