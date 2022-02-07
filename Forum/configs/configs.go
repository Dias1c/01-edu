package configs

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// setting up environment
func SetEnv() {
	readEnv()
	value, found := os.LookupEnv("DB_PATH")
	if found && len(value) > 0 {
		DB_PATH = value
	}
	value, found = os.LookupEnv("DB_NAME")
	if found && len(value) > 0 {
		DB_NAME = value
	}
	value, found = os.LookupEnv("DB_USERNAME")
	if found && len(value) > 0 {
		DB_USERNAME = value
	}
	value, found = os.LookupEnv("DB_PASSWORD")
	if found && len(value) > 0 {
		DB_PASSWORD = value
	}
}

// read the env file
// for future: maybe add overwriting env
func readEnv() {
	file, err := os.Open("./configs/.env")
	if err != nil {
		log.Printf("Error while opening .env file\nDescription: %s", err.Error())
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		envs := strings.Split(text, "=")
		if len(envs) != 2 {
			// can be adding logging
			continue
		}
		// check for already existing env
		if _, found := os.LookupEnv(envs[0]); found {
			continue
		}
		os.Setenv(envs[0], envs[1])
	}
	file.Close()
}
