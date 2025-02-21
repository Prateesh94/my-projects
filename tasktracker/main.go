package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	Id      int    `json:"id"`
	Desc    string `json:"desc"`
	Stat    string `json:"stat"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}

func rd(s string) {
	var a []Task
	fl := 0
	fil, _ := os.ReadFile("tasks.json")
	json.Unmarshal(fil, &a)
	if s == "rand" {
		for i := range a {

			fmt.Println(a[i].Id, a[i].Desc, a[i].Stat)
		}
		return
	}
	for i := range a {
		if a[i].Stat == s {
			fl = 1
			fmt.Println(a[i].Id, a[i].Desc, a[i].Stat)
		}
	}

	if fl == 0 {
		fmt.Println("No such tasks")
	}

}

func add(s string) {
	f, _ := os.ReadFile("index.log")
	a, _ := strconv.Atoi(string(f))
	a++
	os.WriteFile("index.log", []byte(strconv.Itoa(a)), 0644)
	dat := Task{a, s, "todo", time.Now().Format("02/01/2006 15:04:05"), time.Now().Format("02/01/2006 15:04:05")}
	var tr []Task
	d, _ := os.ReadFile("tasks.json")
	json.Unmarshal(d, &tr)
	fil, _ := os.OpenFile("tasks.json", os.O_APPEND|os.O_WRONLY, 0644)
	tr = append(tr, dat)
	bt, _ := json.Marshal(tr)
	os.WriteFile("tasks.json", bt, 0644)
	fil.Close()
	fmt.Println("Task Added Succesfully, ID:- ", a)
}
func update(c int, s string) {
	var t []Task
	fl := 0
	f, _ := os.ReadFile("tasks.json")
	json.Unmarshal(f, &t)
	for i := range t {
		if t[i].Id == c {
			t[i].Desc = s
			t[i].Updated = time.Now().Format("02/01/2006 15:04:05")
			fl = 1
		}
	}
	if fl != 0 {
		fmt.Println("Task not found")
	}
	bt, _ := json.Marshal(t)
	os.WriteFile("tasks.json", bt, 0644)
}
func del(c int) {
	var t []Task
	var j []Task
	fl := 0
	f, _ := os.ReadFile("tasks.json")
	json.Unmarshal(f, &t)
	for i := range t {

		if t[i].Id == c {
			fl = 1
			continue
		} else {
			j = append(j, t[i])
		}

	}
	if fl == 0 {
		fmt.Println("Task not found")
		return
	}
	bt, _ := json.Marshal(j)
	os.WriteFile("tasks.json", bt, 0644)
	fmt.Println("Status updated")
}
func mark(c int, s string) {
	var t []Task
	f, _ := os.ReadFile("tasks.json")
	fl := 0
	json.Unmarshal(f, &t)
	for i := range t {
		if t[i].Id == c {
			t[i].Stat = s
			t[i].Updated = time.Now().Format("02/01/2006 15:04:05")
			fl = 1
		}
	}
	if fl == 0 {
		fmt.Println("Task not found")
		return
	}
	bt, _ := json.Marshal(t)
	os.WriteFile("tasks.json", bt, 0644)
}
func main() {
	_, er := os.ReadFile("tasks.json")
	arg := os.Args[1:]
	var s string

	if er != nil {
		os.Create("tasks.json")
	}
	switch arg[0] {
	case "add":
		for i := range arg[1:] {
			s += arg[i+1] + " "
		}
		add(s)
	case "update":
		t, _ := strconv.Atoi(arg[1])
		if t != 0 {
			i := 2
			for range arg[2:] {
				s += arg[i] + " "
				i++
			}
			update(t, s)
		} else {
			fmt.Println("Invalid Syntax")
			break
		}
	case "delete":
		t, _ := strconv.Atoi(arg[1])
		if t != 0 {
			del(t)
		} else {
			fmt.Println("Invalid Syntax")
			break
		}
	case "mark-in-progress":
		t, _ := strconv.Atoi(arg[1])
		if t != 0 {
			mark(t, "in-progress")
		} else {
			fmt.Println("Invalid Syntax")
			break
		}
	case "mark-done":
		t, _ := strconv.Atoi(arg[1])
		if t != 0 {
			mark(t, "done")
		} else {
			fmt.Println("Invalid Syntax")
			break
		}
	case "mark-todo":
		t, _ := strconv.Atoi(arg[1])
		if t != 0 {
			mark(t, "todo")
		} else {
			fmt.Println("Invalid Syntax")
			break
		}
	case "list":
		if len(arg) == 1 {
			rd("rand")
		} else if arg[1] == "done" {
			rd("done")
		} else if arg[1] == "todo" {
			rd("todo")
		} else if arg[1] == "in-progress" {
			rd("in-progress")
		} else {
			fmt.Println("Invalid Syntax")
			break
		}
	default:
		fmt.Println("Invalid Syntax")
	}

}
