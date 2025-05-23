package backup

import (
	"log"
	"os"
)

const dirPath = "/tmp/transit/"

func StoreBackup(fileName string, body []byte) error {
	if err := os.Mkdir(dirPath, 0644); err != nil {
		return err
	}
	filePath := dirPath + fileName
	log.Println(filePath)
	err := os.WriteFile(fileName, body, 0644)
	return err
}

func LoadBackup(fileName string) ([]byte, error) {
	filePath := dirPath + fileName
	body, err := os.ReadFile(filePath)
	return body, err
}
