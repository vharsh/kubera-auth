package usermanager

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"github.com/mayadata-io/kubera-auth/pkg/errors"
	"github.com/mayadata-io/kubera-auth/pkg/models"
	"github.com/mayadata-io/kubera-auth/pkg/store"
)

// GetUserByUserName get the user information
func GetUserByUserName(userStore *store.UserStore, userName string) (user *models.UserCredentials, err error) {
	query := bson.M{"username": userName}
	user, err = userStore.GetUser(query)
	if err != nil && err == mgo.ErrNotFound {
		err = errors.ErrInvalidUser
	}
	return
}

// GetUserByUID get the user information
func GetUserByUID(userStore *store.UserStore, userID string) (user *models.UserCredentials, err error) {
	query := bson.M{"uid": userID}
	user, err = userStore.GetUser(query)
	if err != nil && err == mgo.ErrNotFound {
		err = errors.ErrInvalidUser
	}
	return
}

// GetAllUsers get the user information
func GetAllUsers(userStore *store.UserStore) ([]*models.PublicUserInfo, error) {
	users, err := userStore.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var allUsers []*models.PublicUserInfo
	for _, user := range users {
		allUsers = append(allUsers, user.GetPublicInfo())
	}
	return allUsers, nil
}
