package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDataBaseDriver() (*gorm.DB, error) {
	dsn := "root:0000@tcp(127.0.0.1:3306)/controlEscolar"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error al conectarse a la base de datos", err)
	} else {
		fmt.Println("Conexión a la base de datos exitosa")
	}
	return db, nil
}
