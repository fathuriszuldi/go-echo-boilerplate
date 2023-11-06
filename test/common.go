package test

import (
	"github.com/joho/godotenv"
	BlogModels "github.com/triaton/go-echo-boilerplate/blogs/models"
	"github.com/triaton/go-echo-boilerplate/database"
	UserModels "github.com/triaton/go-echo-boilerplate/users/models"
	"log"
	"os"
	"fmt"
)

func LoadTestEnv() error {
	path, errOS := os.Getwd()
	if errOS != nil {
    	log.Println(errOS)
	}
	fmt.Println(path)
	err := godotenv.Load(os.ExpandEnv("$GOPATH/go-echo-boilerplate/go-echo-boilerplate/test.env"))
	if err != nil {
		log.Fatal("failed to load test env config: ", err)
	}
	return err
}

func InitTest() {
	err := LoadTestEnv()
	db := database.GetInstance()
	db.DropTable("migrations")
	db.DropTableIfExists(&UserModels.User{})
	db.DropTableIfExists(&BlogModels.Blog{})
	m := database.GetMigrations(db)
	err = m.Migrate()
	if err != nil {
		log.Fatal("failed to run db migration: ", err)
	}
}
