package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"modules/internal/db"
	"strings"
	_ "github.com/go-sql-driver/mysql"
)

var userStruct struct{
	ID int32
	Email string	
}

func DatabaseConnect() *db.Queries{
	dbconn, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/NEWSLETTER")
	if err != nil {
		log.Fatal(err)
		
	}
	database := db.New(dbconn);
	return database;
}



func GetLastID() string {
	database := DatabaseConnect();
	ID, err := database.GetLastID(context.Background());
	if err != nil {
		log.Fatal(err)
	}
	value := fmt.Sprintf("%d", ID)
	return value;
}

func AddNewUser(ID int32, email string) {
	database := DatabaseConnect();
	newUser, err_ := database.GetEmail(context.Background(), email)
	if err_ != nil {
		fmt.Println(err_)
	}
	if (strings.EqualFold(newUser.Email, email)) {
		fmt.Println("A user with that email already exists")
	} else {
		database.InsertUser(context.Background(), db.InsertUserParams{ID, email});
		fmt.Println("User registred")
	}
	
}

