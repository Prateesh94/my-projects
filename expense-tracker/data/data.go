package data

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"
)

type Table struct {
	Id     int       `json:"id"`
	Date   time.Time `json:"date"`
	Desc   string    `json:"description"`
	Amount float64   `json:"amount"`
}
type Budget struct {
	Month int     `json:"month"`
	Bud   float64 `json:"budget"`
}

func Bud(a int) {
	var b float64
	fmt.Println("Set budget for month ", a)
	fmt.Scanf("%v\n", &b)
	var bud []Budget
	fl := 0
	var bd Budget
	bd.Month = a
	bd.Bud = b
	rd, err := os.ReadFile("budget.json")
	if err != nil {
		bud = append(bud, bd)
		dt, _ := json.Marshal(bud)
		os.WriteFile("budget.json", dt, 0644)
		fmt.Printf("Budget for %d set to %v\n", a, b)
		return
	} else {
		json.Unmarshal(rd, &bud)

		for r := range bud {
			if bud[r].Month == a {
				bud[r].Bud = b
				fl = 1
			}
		}
	}
	if fl == 0 {
		bud = append(bud, bd)
		dt, _ := json.Marshal(bud)
		os.WriteFile("budget.json", dt, 0644)
		fmt.Printf("Budget for %d set to %v\n", a, b)
		return
	}
	dt, _ := json.Marshal(bud)
	os.WriteFile("budget.json", dt, 0644)
	fmt.Printf("Budget for %d set to %v\n", a, b)
}
func chkbud(a int) float64 {
	rd, err := os.ReadFile("budget.json")
	var bd []Budget
	var db []Table
	var rem float64
	var cmp float64
	if err != nil {
		return 0
	} else {
		json.Unmarshal(rd, &bd)
		dc, _ := os.ReadFile("expense.json")
		json.Unmarshal(dc, &db)
		for r := range bd {
			if bd[r].Month == a {
				rem = bd[r].Bud
			}
		}
		for r := range db {
			if int(db[r].Date.Month()) == a {
				cmp += db[r].Amount
			}
		}
	}
	if rem < cmp {
		fmt.Println("Going over budget by:- ", cmp-rem)
	}
	return rem
}
func writefil(a []Table) {
	w, _ := json.Marshal(a)
	os.WriteFile("expense.json", w, 0644)
}

func indexgen() int {
	a, er := os.ReadFile("index.txt")
	if er != nil {
		b := []byte(strconv.Itoa(1))
		os.WriteFile("index.txt", b, 0644)
		return 1
	} else {
		b, _ := strconv.Atoi(string(a))
		b++
		os.WriteFile("index.txt", []byte(strconv.Itoa(b)), 0644)
		return b
	}
}
func Add(s []string) {
	var b float64
	var er error
	var Tb Table
	var Db []Table
	var fl int
	rd, err := os.ReadFile("expense.json")
	if err != nil {
		writefil(Db)
	} else {
		json.Unmarshal(rd, &Db)
	}
	for r := range s {
		if s[r] == "--amount" {
			fl = 2
			b, er = strconv.ParseFloat(s[r+1], 64)
			if er == nil && b > 0 {
				Tb.Amount = b
			} else {
				fmt.Println("Enter valid amount")
				return
			}

			break
		}

		Tb.Desc += s[r] + " "
	}
	if fl == 0 {
		fmt.Println("Please use --amount keyword followed by a valid amount")
		return
	}
	Tb.Id = indexgen()
	Tb.Date = time.Now()
	Db = append(Db, Tb)
	writefil(Db)
	fmt.Printf("Expense added succesfully (ID: %d)\n", Tb.Id)
	chkbud(int(Tb.Date.Month()))
}

func Del(a int) {
	d, er := os.ReadFile("expense.json")
	if er != nil {
		fmt.Println("No records Exist")
		return
	}
	var dt []Table
	var bt []Table
	var tb Table
	fl := 0
	json.Unmarshal(d, &dt)
	for r := range dt {
		if dt[r].Id == a {
			fl = 1
			tb = dt[r]
			continue
		} else {
			bt = append(bt, dt[r])
		}
	}
	if fl == 0 {
		fmt.Println("ID not found")
	} else {
		writefil(bt)
		fmt.Printf("Deleted expense (ID:-%d) successfully\n", a)
		chkbud(int(tb.Date.Month()))
	}
}

func Update(a int, s []string) {
	var in1, in2 int
	var s1 []string
	var r int
	var amt float64
	var tb Table
	d, er := os.ReadFile("expense.json")
	if er != nil {
		fmt.Println("No records Exist")
		return
	}
	for r = range s {
		if s[r] == "--description" {
			in1 = r
			s1 = s[in1+1:]
		}
		if s[r] == "--amount" {
			in2 = r
			amt, _ = strconv.ParseFloat(s[in2+1], 64)
			if amt <= 0 {
				fmt.Println("Enter valid amount")
				return
			}
		}
	}
	if s1 != nil && in2 != 0 {
		s1 = s[in1+1 : in2]
	}
	var ss string
	for ll := range s1 {
		ss += s1[ll] + " "
	}
	var dt []Table
	fl := 0
	json.Unmarshal(d, &dt)
	for r := range dt {
		if dt[r].Id == a {
			fl = 2
			if s1 != nil {
				dt[r].Desc = ss

			}
			if amt > 0 {
				dt[r].Amount = amt
				tb = dt[r]
			}
			break
		}
	}
	if fl == 0 {
		fmt.Printf("ID:- %d not found\n", a)
		return
	}
	fil, _ := json.Marshal(dt)
	os.WriteFile("expense.json", fil, 0644)
	fmt.Printf("ID:- %d updated\n", a)
	chkbud(int(tb.Date.Month()))
}

func View(a int) {
	d, er := os.ReadFile("expense.json")
	if er != nil {
		fmt.Println("No records Exist")
		return
	}
	w := tabwriter.NewWriter(os.Stdout, 1, 2, 5, ' ', 0)
	fmt.Fprintln(w, "ID\t DATE\t DESCRIPTION\t AMOUNT")
	var dt []Table
	json.Unmarshal(d, &dt)
	if a == 0 {
		for r := range dt {
			fmt.Fprintln(w, dt[r].Id, "\t", dt[r].Date.Format("02-01-2006"), "\t", dt[r].Desc, "\t", dt[r].Amount)
		}

	} else {
		for r := range dt {

			if int(dt[r].Date.Month()) == a {
				fmt.Fprintln(w, dt[r].Id, "\t", dt[r].Date.Format("02-01-2006"), "\t", dt[r].Desc, "\t", dt[r].Amount)
			}

		}

	}
	w.Flush()
}

func Summary(a int) {
	d, er := os.ReadFile("expense.json")
	var fl int
	if er != nil {
		fmt.Println("No records Exist")
		return
	}
	var dt []Table
	var sum float64
	json.Unmarshal(d, &dt)
	if a == 0 {
		for r := range dt {
			sum += dt[r].Amount
		}
		fmt.Println("Total Expenses:- ", sum)
		return

	} else {
		var mon string
		for r := range dt {

			if int(dt[r].Date.Month()) == a {
				fl = 1
				sum += dt[r].Amount
				mon = dt[r].Date.Month().String()
			}

		}
		if fl != 0 {
			fmt.Printf("Expenses for month of %s is:- %v \n", mon, sum)

		} else {
			fmt.Println("No records found")
		}
	}
}

func Export(a int) {
	var name string
	fl := 0
	if a == 0 {
		name = "All Summary.csv"
	} else {
		name = "Summary for Month-" + strconv.Itoa(a) + ".csv"
	}
	d, er := os.ReadFile("expense.json")
	if er != nil {
		fmt.Println("No Records Exist")
		return
	}
	head := []string{"ID", "DATE", "DESCRIPTION", "AMOUNT"}
	var row []string
	var dt []Table
	var bt []Table
	var col [][]string
	var sum float64
	json.Unmarshal(d, &dt)
	bt = dt
	if a == 0 && len(bt) != 0 {
		dd, _ := os.Create(name)
		ww := csv.NewWriter(dd)
		ww.Write(head)
		for r := range dt {

			row = []string{strconv.Itoa(dt[r].Id), dt[r].Date.Format("02-01-2006"), dt[r].Desc, strconv.FormatFloat(dt[r].Amount, 'f', -1, 64)}
			ww.Write(row)
			sum += dt[r].Amount
		}
		row = []string{"Total Expense:-", "", "", strconv.FormatFloat(sum, 'f', -1, 64)}
		ww.Write(row)
		ww.Flush()
		dd.Close()
		fmt.Println("Exported ", name)
	} else if len(bt) != 0 {
		for r := range dt {
			if int(dt[r].Date.Month()) == a {
				fl = 1
				row = []string{strconv.Itoa(dt[r].Id), dt[r].Date.Format("02-01-2006"), dt[r].Desc, strconv.FormatFloat(dt[r].Amount, 'f', -1, 64)}
				col = append(col, row)
				sum += dt[r].Amount
			}

		}
	} else {
		fmt.Println("No Records to Export")
	}
	if fl == 1 {
		dd, _ := os.Create(name)
		ww := csv.NewWriter(dd)
		ww.Write(head)
		ww.WriteAll(col)
		row = []string{"Total Expense:-", "", "", strconv.FormatFloat(sum, 'f', -1, 64)}
		ww.Write(row)
		su := chkbud(a)
		row = []string{"Total Budget-", "", "", strconv.FormatFloat(su, 'f', -1, 64)}
		ww.Write(row)
		row = []string{"Budget Remaining-", "", "", strconv.FormatFloat(su-sum, 'f', -1, 64)}
		ww.Write(row)
		ww.Flush()
		dd.Close()
		fmt.Println("Exported ", name)
	} else {
		fmt.Println("No Records to Export")
	}
}
