package backups

import (
	"os"
	"strconv"
	"time"
)

func FileName(dbname string) string {
	filename := "backup_" + dbname + "_" + strconv.FormatInt(time.Now().Unix(), 10) + ".sql"
	return filename
}

func CheckFileNameExist(filename, dir string) bool {
	// Check if the file exists in the specified directory
	_, err := os.Stat(dir + filename)
	if err == nil {
		return false // File exists
	} else {
		return true // File does not exist
	}

}
