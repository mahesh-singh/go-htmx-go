package main

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv() Env {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	return Env{
		Db:       os.Getenv("DB"),
		User:     os.Getenv("POSTGRESUSER"),
		Password: os.Getenv("PASSWORD"),
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		Listen:   os.Getenv("LISTEN"),
	}
}
