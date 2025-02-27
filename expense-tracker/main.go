package main

import (
	"expense-tracker/data"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var fl int

	arg := os.Args[1:]

	cd := strings.ToLower(arg[0])
	switch cd {
	case "add":
		if len(arg) > 1 {
			for r := range arg {
				if arg[r] == "--description" {
					data.Add(arg[r+1:])
					fl = 1
					break
				}

			}
		} else {
			fmt.Println("Invalid syntax")
			break
		}
		if fl == 0 {
			fmt.Println("--description not found")
		}
		break
	case "delete":
		if arg[1] == "--id" {
			d, er := strconv.Atoi(arg[2])
			if er != nil {
				fmt.Println("Invalid Id")
			} else {
				data.Del(d)
			}
		} else {
			fmt.Println("missing or misplaced --id tag")
			break
		}
	case "update":
		if arg[1] == "--id" {
			d, er := strconv.Atoi(arg[2])
			if er != nil {
				fmt.Println("Invalid Id")
			} else {
				data.Update(d, arg[3:])
			}
		} else {
			fmt.Println("missing or misplaced --id tag")
			break
		}
	case "list":
		dat := 0
		if len(arg) > 2 && arg[1] == "--month" {
			dat, _ = strconv.Atoi(arg[2])
			if dat > 12 || dat < 0 {
				fmt.Println("Enter correct month")
				break
			}
			data.View(dat)
		} else {
			data.View(dat)
		}
		break

	case "summary":
		dat := 0
		if len(arg) > 2 && arg[1] == "--month" {
			dat, _ = strconv.Atoi(arg[2])
			if dat > 12 || dat < 0 {
				fmt.Println("Enter correct month")
				break
			}
			data.Summary(dat)
		} else {
			data.Summary(dat)
		}
		break
	case "budget":
		dat := 0
		if len(arg) > 2 && arg[1] == "--month" {
			dat, _ = strconv.Atoi(arg[2])
			if dat > 12 || dat < 0 {
				fmt.Println("Enter correct month")
				break
			}
			data.Bud(dat)
		} else {
			fmt.Println("use --month tag to specify month of budget")
		}
		break
	case "export":
		dat := 0
		if len(arg) > 2 && arg[1] == "--month" {
			dat, _ = strconv.Atoi(arg[2])
			if dat > 12 || dat < 0 {
				fmt.Println("Enter correct month")
				break
			}
			data.Export(dat)
		} else {
			data.Export(dat)
		}
		break
	default:
		fmt.Println("Invalid Command")
		break
	}

}
