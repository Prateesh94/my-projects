package blogops

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	username = "root"
	password = "admin"
	hostname = "0.0.0.0:3306"
	dbname   = "blogs"
)

var db *sql.DB

type Blogs struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Tags     string `json:"tags"`
	Cr_at    string `json:"createdAt"`
	Up_at    string `json:"updatedAt"`
}

func init() {
	db, _ = sql.Open("mysql", username+":"+password+"@tcp("+hostname+")/"+dbname)
	conn()
}
func conn() {
	err := db.Ping()
	fmt.Println(err, " is established")
	qry := `create table if not exists myblogs(id int not null primary key auto_increment,title text,content text,category text,tags text,created datetime current_timestamp,updated datetime current_timestamp)`
	db.Exec(qry)
}

func AddData(w http.ResponseWriter, r *http.Request) {
	type tempst struct {
		Title    string `json:"title"`
		Content  string `json:"content"`
		Category string `json:"category"`
		Tags     string `json:"tags"`
	}
	var tmp tempst
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	er := d.Decode(&tmp)

	fmt.Println(er)
	if er != nil {
		http.Error(w, "Incorrect json", http.StatusNotAcceptable)
		return
	}
	if tmp.Title == "" || tmp.Content == "" || tmp.Tags == "" || tmp.Category == "" {
		http.Error(w, "found empty json fields", http.StatusForbidden)
		return
	} else {
		qry := `insert into myblogs(title,content,category,tags) values(?,?,?,?)`
		db.Exec(qry, tmp.Title, tmp.Content, tmp.Category, tmp.Tags)
		json.NewEncoder(w).Encode(tmp)
	}

}

func DelData(w http.ResponseWriter, r *http.Request) {
	a := mux.Vars(r)
	s := a["id"]
	df, _ := strconv.Atoi(s)
	if df == 0 {
		http.Error(w, "Incorrect ID parameter", http.StatusBadRequest)
		return
	}
	qry := `delete from myblogs where id=?`
	rd, _ := db.Exec(qry, s)
	ad, _ := rd.RowsAffected()
	if ad == 0 {
		http.Error(w, "Blog id does not exist", http.StatusBadRequest)
		return
	} else {
		fmt.Fprintf(w, "Deleted blog post ID(%s)", s)
	}

}
func View(w http.ResponseWriter, r *http.Request) {
	a := mux.Vars(r)
	s := a["id"]
	fmt.Println(s)
	if s != "" {
		var tem Blogs
		row := db.QueryRow(`select * from myblogs where id=?`, s)
		err := row.Scan(&tem.Id, &tem.Title, &tem.Content, &tem.Category, &tem.Tags, &tem.Cr_at, &tem.Up_at)
		if err == nil {
			json.NewEncoder(w).Encode(tem)
		} else {
			fmt.Fprintln(w, "Blog does not exist", http.StatusNotFound)

		}

		return
	} else {
		var tem []Blogs
		var dt Blogs
		rows, _ := db.Query(`select * from myblogs`)
		for rows.Next() {
			rows.Scan(&dt.Id, &dt.Title, &dt.Content, &dt.Category, &dt.Tags, &dt.Cr_at, &dt.Up_at)
			tem = append(tem, dt)
		}
		if tem != nil {
			json.NewEncoder(w).Encode(tem)
		} else {
			fmt.Fprintln(w, "No blogs to display", http.StatusNotFound)
		}
	}
}

func UpData(w http.ResponseWriter, r *http.Request) {
	a := mux.Vars(r)
	s := a["id"]

	df, _ := strconv.Atoi(s)
	if df == 0 {
		http.Error(w, "Incorrect ID parameter", http.StatusBadRequest)
		return
	}
	var tem Blogs
	row := db.QueryRow(`select * from myblogs where id=?`, s)
	err := row.Scan(&tem.Id, &tem.Title, &tem.Content, &tem.Category, &tem.Tags, &tem.Cr_at, &tem.Up_at)
	if err != nil {
		http.Error(w, "ID not found", http.StatusNotFound)
		return
	}
	rdjson := json.NewDecoder(r.Body)
	rdjson.DisallowUnknownFields()
	var updat Blogs
	er := rdjson.Decode(&updat)
	if er != nil {
		http.Error(w, "Incorrect json", http.StatusNotAcceptable)
		return
	}
	if updat.Title != "" {
		tem.Title = updat.Title
	}
	if updat.Content != "" {
		tem.Content = updat.Content
	}
	if updat.Category != "" {
		tem.Category = updat.Category
	}
	if updat.Tags != "" {
		tem.Tags = updat.Tags
	}
	qry := `update myblogs set title=?, content=?, category=?, tags=?,updated=CURRENT_TIMESTAMP where id=?`
	rd, er := db.Exec(qry, tem.Title, tem.Content, tem.Category, tem.Tags, s)
	fmt.Println("updated", er)
	ad, _ := rd.RowsAffected()
	if ad == 0 {
		http.Error(w, "Blog id does not exist", http.StatusBadRequest)
		return
	} else {
		row = db.QueryRow(`select * from myblogs where id=?`, s)
		row.Scan(&tem.Id, &tem.Title, &tem.Content, &tem.Category, &tem.Tags, &tem.Cr_at, &tem.Up_at)
		json.NewEncoder(w).Encode(tem)
	}

}

func Blogsearch(w http.ResponseWriter, r *http.Request) {
	a := mux.Vars(r)
	s := a["term"]
	ss := s
	s += "%"
	var dt Blogs
	var dat []Blogs
	rows, _ := db.Query(`select * from myblogs where title like ? or content like ? or category like ?`, s, s, s)
	for rows.Next() {
		rows.Scan(&dt.Id, &dt.Title, &dt.Content, &dt.Category, &dt.Tags, &dt.Cr_at, &dt.Up_at)
		dat = append(dat, dt)
	}
	if dat != nil {
		json.NewEncoder(w).Encode(dat)
	} else {
		fmt.Fprintf(w, "No blogs found containing %s", ss)
	}
}
