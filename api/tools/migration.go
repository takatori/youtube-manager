package main

import (
	"github.com/sirupsen/logrus"
	"github.com/takatori/youtube-manager/api/databases"
	"github.com/takatori/youtube-manager/api/models"
)

func main() {
	db, err := databases.Connect()
	defer db.Close()
	if err != nil {
		logrus.Fatal(err)
	}
	db.Debug().AutoMigrate(&models.User{})
	db.Debug().AutoMigrate(&models.Favorite{})
}
