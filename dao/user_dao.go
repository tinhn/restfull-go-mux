package dao

import (
	. "../model"
	. "../mongodb"

	"gopkg.in/mgo.v2/bson"
)

const (
	COLLECTION = "users"
)

// Find list of users
func FindAll() ([]User, error) {
	var users []User
	err := DB.C(COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}

// Find a user by its id
func FindById(id string) (User, error) {
	var user User
	err := DB.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// Insert a user into database
func Insert(user User) error {
	err := DB.C(COLLECTION).Insert(&user)
	return err
}

// Delete an existing user
func Delete(user User) error {
	err := DB.C(COLLECTION).Remove(&user)
	return err
}

// Update an existing user
func Update(user User) error {
	err := DB.C(COLLECTION).UpdateId(user.ID, &user)
	return err
}
