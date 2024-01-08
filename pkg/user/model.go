package user

import (
	pb_unit_user "github.com/byeol-i/firebase-auth-module/pb/unit/user"
)

type User struct {
	UserInterface
	UserImpl
	UserCredential
}

type UserImpl struct {
	Name string `validate:"required" json:"name" example:"gil dong"`
	Email string `validate:"required" json:"email" example:"example@example.com"`
}

type UserCredential struct {
	Uid string
	Token string
}

type UserInterface interface {
	GetCredential() string
	SetEmail(string) 
	SetName(string)
	GetName(string)
	GetEmail(string)
	ToProto() (*pb_unit_user.User, error)
}

func NewUser() *User {
	newUser := &User{}
	return newUser
}

func NewUserFromProto(pbUser *pb_unit_user.User) (*User, error) {
	newUser := &User{
		UserImpl: UserImpl{
			Email: pbUser.Email,
			Name: pbUser.Name,
		},
		UserCredential: UserCredential{
			Uid: pbUser.UserCredential.Uid,
		},	
	}

	err := UserValidator(&newUser.UserImpl)
	if err != nil {
		return nil, err
	}
	
	return newUser, nil
}

func (u *User) GetUserCredential() string {
	return u.Uid
}

func (u *User) SetUserCredential(userCredential UserCredential) {
	u.UserCredential = userCredential
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) SetName(uName string) {
	u.Name = uName
}

func (u *User) ToProto() (*pb_unit_user.User, error) {
	pbUnit := &pb_unit_user.User{}

	id := &pb_unit_user.UserCredential{
		Uid: u.Uid,
	}

	pbUnit.UserCredential = id
	pbUnit.Name = u.Name

	return pbUnit, nil
}