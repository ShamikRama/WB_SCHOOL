package repository

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func ConnectToDB(envfile string) error {
	err := godotenv.Load(envfile)
	if err != nil {
		return fmt.Errorf("Error reading .env %s", err)
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	sc, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return fmt.Errorf("Error opening sql %s", err)
	}
	defer sc.Close()

	if err = sc.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Print("Успешное подключение к базе данных")
	return nil

}
