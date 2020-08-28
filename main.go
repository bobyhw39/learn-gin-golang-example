package main

import (
	"fmt"
	"github.com/bobyhw39/rest_mvc/Config"
	"github.com/bobyhw39/rest_mvc/Entity"
	"github.com/bobyhw39/rest_mvc/Routes"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var err error

func main(){
	Config.DB, err = gorm.Open("postgres", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Entity.Account{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
