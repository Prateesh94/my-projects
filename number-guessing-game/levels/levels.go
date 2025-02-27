package levels

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"text/tabwriter"
	"time"
)

type Score struct {
	Mode string `json:"mode"`
	Num  int    `json:"attempts"`
}

func Easy(a int) (int, time.Duration, error) {
	j := rand.Intn(100)
	t1 := time.Now()
	var k int
	for m := range a {
		fmt.Printf("Enter your Guess:- ")
		fmt.Scanf("%d\n", &k)
		if k == j {
			t2 := time.Since(t1)
			writ(a, m+1)
			return m + 1, t2, nil
		} else if k < j {
			fmt.Printf("Incorrect! The number is greater than %d\n", k)
		} else {
			fmt.Printf("Incorrect! The number is less than %d\n", k)
		}
	}
	er := errors.New("lost")
	t2 := time.Since(t1)
	return j, t2, er
}

func writ(a int, b int) {
	var s string
	switch a {
	case 10:
		s = "Easy"

	case 5:
		s = "Medium"

	case 3:
		s = "Hard"

	}
	//fmt.Println(s)
	_, er := os.ReadFile("score.json")
	if er != nil {
		dt := []Score{{"Easy", 100}, {"Medium", 100}, {"Hard", 100}}
		fil, _ := json.Marshal(dt)
		os.WriteFile("score.json", fil, 0644)
	}
	d, _ := os.ReadFile("score.json")
	var dt []Score
	json.Unmarshal(d, &dt)
	for r := range dt {
		fmt.Println(dt[r], s)
		if dt[r].Mode == s {
			//fmt.Println("matched")
		}
		if dt[r].Mode == s && dt[r].Num > b {
			dt[r].Num = b
			fmt.Println("New High-Score")
		}

	}
	fil, _ := json.Marshal(dt)
	os.WriteFile("score.json", fil, 0644)
}

func Read() {
	_, er := os.ReadFile("score.json")
	if er != nil {
		fmt.Println("No records exist")
		return
	}
	d, _ := os.ReadFile("score.json")
	var dt []Score
	json.Unmarshal(d, &dt)
	w := tabwriter.NewWriter(os.Stdout, 1, 2, 5, ' ', 0)
	fmt.Fprintln(w, "Mode", " Score")
	for r := range dt {

		fmt.Fprintln(w, dt[r].Mode, dt[r].Num)
	}
	w.Flush()
}
