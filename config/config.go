package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type DatabaseInfo struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

type configuration struct {
	OffsetUpdate int
	Timeout      int
	ApiToken     string
	Debug        bool
	Database     DatabaseInfo
}

var instance *configuration

func NewInstance() *configuration {
	if instance == nil {
		instance = configuration{}.readConfig()
	}
	return instance
}

func (c configuration) readConfig() *configuration {
	file, _ := os.Open("./config/config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return &configuration
}
