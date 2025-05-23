package model

import "github.com/Binit-Dhakal/e-fit/users/gen"

func ProtoToUser(u *gen.User) *User {
	return &User{
		ID:            u.Id,
		Username:      u.Username,
		Password:      password{hash: &u.Password},
		Email:         u.Email,
		EmailVerified: u.EmailVerified,
		IsGuest:       u.IsGuest,
	}
}

func UserToProto(u *User) *gen.User {
	return &gen.User{
		Id:            u.ID,
		Username:      u.Username,
		Password:      *u.Password.hash,
		Email:         u.Email,
		EmailVerified: u.EmailVerified,
		IsGuest:       u.IsGuest,
	}
}
