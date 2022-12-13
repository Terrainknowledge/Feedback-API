package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	Enable VIPER to read Environment Variables
	viper.AutomaticEnv()
	To get the value from the config file using key
	viper package read .env
	viper_user := viper.Get("POSTGRES_USER")
	viper_password := viper.Get("POSTGRES_PASSWORD")
	viper_db := viper.Get("POSTGRES_DB")
	viper_host := viper.Get("POSTGRES_HOST")
	viper_port := viper.Get("POSTGRES_PORT")

	// https://gobyexample.com/string-formatting
	prosgret_conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", viper_host, viper_port, viper_user, viper_db, viper_password)
	fmt.Println("conname is\t\t", prosgret_conname)

	db, err := gorm.Open("postgres", prosgret_conname)
	fmt.Println(db)
	if err != nil {
		fmt.Println("error", err)
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&Feedback{})
	// Initialise value
	m := Feedback{Name: " emd ", Email: "sfsaf", Ftype: "sddsf", Text: "test "}
	db.Create(&m)
	return db
}
