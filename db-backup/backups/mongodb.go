package backups

import (
	"fmt"
	"os/exec"

	_ "go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo/options"
)

func BackupMongo() {
	var con DbConfig
	var command string
	dir := "/data/mongo/"
	fmt.Print("Enter Host: ")
	fmt.Scanln(&con.Host)
	fmt.Print("Enter Port: ")
	fmt.Scanln(&con.Port)
	fmt.Print("Enter User: ")
	fmt.Scanln(&con.User)
	fmt.Print("Enter Password: ")
	fmt.Scanln(&con.Password)
	fmt.Print("Enter Database Name: ")
	fmt.Scanln(&con.DBName)
	filename := FileName(con.DBName)
	command = fmt.Sprintf("/usr/bin/mongodump --host %s --port %d --username %s --password %s --db %s --out %s", con.Host, con.Port, con.User, con.Password, con.DBName, dir+filename)
	cmd := exec.Command("bash", "-c", command)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error executing command:", err)
	} else {
		fmt.Println("Backup successful, file saved at:", dir+filename)
	}
}

func RestoreMongo() {
	var con DbConfig
	var command string
	var filename string
	dir := "/data/mongo/"
	fmt.Print("Enter Host: ")
	fmt.Scanln(&con.Host)
	fmt.Print("Enter Port: ")
	fmt.Scanln(&con.Port)
	fmt.Print("Enter User: ")
	fmt.Scanln(&con.User)
	fmt.Print("Enter Password: ")
	fmt.Scanln(&con.Password)
	fmt.Print("Enter Database Name: ")
	fmt.Scanln(&con.DBName)
	fmt.Print("Enter File Name: ")
	fmt.Scanln(&filename)
	if CheckFileNameExist(filename, dir) {
		fmt.Println("File does not exist, please check the file name and try again.")
		return
	}
	command = fmt.Sprintf("/usr/bin/mongorestore --host %s --port %d --username %s --password %s --db %s %s", con.Host, con.Port, con.User, con.Password, con.DBName, dir+filename)
	cmd := exec.Command("bash", "-c", command)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error executing command:", err)
	} else {
		fmt.Println("Restore successful from file:", dir+filename)
	}
}
