package controller

import (
	"context"
	ent "root/database"

	"github.com/sirupsen/logrus"
)

// Создание пользователя
func CreateUser(client *ent.Client, name, email, password string, log *logrus.Logger) {
	log.Info("creating new user")

	u, err := client.User.
		Create().
		SetName(name).
		SetEmail(email).
		SetPassword(password).
		Save(context.Background())

	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}

	log.Infof("User created: %v", u)
}

// Получить всех пользователей
func GetAllUsers(client *ent.Client, log *logrus.Logger) []*ent.User {
	log.Info("querying for users")

	users, err := client.User.
		Query().
		All(context.Background())
	if err != nil {
		log.Fatalf("failed getting all users: %v", err)
	}

	log.Info("all users returned")
	return users
}
