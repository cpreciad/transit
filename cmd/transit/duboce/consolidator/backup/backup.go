package backup

import (
	"os"
)

// this should be a path that gets saved to var/, save to local for now
const dirPath = "backups/"

func StoreBackup(fileName string, body []byte) error {
	// set up the path for writing backup data
	setup()

	filePath := dirPath + fileName
	err := os.WriteFile(filePath, body, 0777)
	return err
}

func LoadBackup(fileName string) ([]byte, error) {
	filePath := dirPath + fileName
	body, err := os.ReadFile(filePath)
	return body, err
}

func setup() error {
	if err := os.MkdirAll(dirPath, 0777); err != nil {
		return err
	}
	if err := os.Chmod(dirPath, 0777); err != nil {
		return err
	}
	return nil
}
