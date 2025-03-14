package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func getConnectionString() (string, error) {
	godotenv.Load()
	var dbuser string = os.Getenv("DBUSER")
	var dbpass string = os.Getenv("DBPASS")
	var dbhost string = os.Getenv("DBHOST")
	var dbport string = os.Getenv("DBPORT")
	var dbname string = os.Getenv("DBNAME")
	if dbuser == "" || dbpass == "" || dbhost == "" || dbport == "" || dbname == "" {
		return "", fmt.Errorf("error al obtener variables de entorno")
	}
	var connectionString string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbuser, dbpass, dbhost, dbport, dbname)
	return connectionString, nil
}

func ConnectDB() error {
	connectionString, connectionStringError := getConnectionString()
	if connectionStringError != nil {
		return connectionStringError
	}
	var dbConnectionError error
	db, dbConnectionError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbConnectionError != nil {
		return dbConnectionError
	}
	db.AutoMigrate(
	//colocar modelos
	)
	fmt.Println("Conectado a base de datos")
	return nil
}

func GetInstance() *gorm.DB {
	return db
}
