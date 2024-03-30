package main

type Env struct {
	Db       string
	User     string
	Password string
	Host     string
	Port     string
	Listen   string
}

type Person struct {
	Name    string
	Address string
	Email   string
}
