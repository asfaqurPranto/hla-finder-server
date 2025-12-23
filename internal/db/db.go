package db

import (
	"hla_finder/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB


func Connect_MySql_Server(){
	dsn := "root:root@tcp(127.0.0.1:3306)/hla_db2?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		panic("MySQL conncetion failed"+ err.Error())
	}
	DB=db

}

func Create_Schema(){
	DB.AutoMigrate(&models.User{},&models.HLA{})
}