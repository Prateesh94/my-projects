package backups

import (
	"fmt"
	"os/exec"
)

func BackupPostgres() {
	// Implement PostgreSQL backup logic here
	var con DbConfig
	var command string
	dir := "/data/postgres/"
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
	command = fmt.Sprintf("/usr/bin/pg_dump -h %s -p %d -U %s -W %s > %s", con.Host, con.Port, con.User, con.DBName, dir+filename)

	cmd := exec.Command("bash", "-c", command)
	cmd.Env = append(cmd.Env, "PGPASSWORD="+con.Password)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error executing command:", err)
	} else {
		fmt.Println("Backup successful, file saved at:", dir+filename)
	}
}

func RestorePostgres() {
	// Implement PostgreSQL restore logic here
	var con DbConfig
	var command string
	var filename string
	dir := "/data/postgres/"
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

	command = fmt.Sprintf("/usr/bin/pg_restore --clean --if-exists -h %s -p %d -U %s -d %s -v %s", con.Host, con.Port, con.User, con.DBName, dir+filename)
	cmd := exec.Command("bash", "-c", command)
	cmd.Env = append(cmd.Env, "PGPASSWORD="+con.Password)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error executing command:", err)
	} else {
		fmt.Println("Restore successful from file:", dir+filename)
	}
}
