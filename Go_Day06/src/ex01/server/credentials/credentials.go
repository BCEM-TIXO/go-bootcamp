package credentials

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type AdminContent struct {
	Login, Password string
}

type DBContent struct {
	Login, Password string
}

func GetAdminCredentials() (AdminContent, error) {
	file, err := os.Open("credentials/admin_credentials.txt")
	admin := AdminContent{}
	if err != nil {
		return admin, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			return admin, errors.New("wrong admin credentials format")
		}
		admin.Login = parts[0]
		admin.Password = parts[1]
	}
	return admin, nil
}

func GetDBCredentials() (DBContent, error) {
	file, err := os.Open("credentials/db_credentials.txt")
	db := DBContent{}
	if err != nil {
		return db, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			return db, errors.New("wrong db credentials format")
		}
		db.Login = parts[0]
		db.Password = parts[1]
	}
	return db, nil
}
