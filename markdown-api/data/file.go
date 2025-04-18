package data

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/russross/blackfriday/v2"
)

func init() {
	os.Mkdir("file", os.ModePerm)
}
func checkFileExists(filename string) bool {
	fil := filepath.Join("file", filename)
	_, err := os.Stat(fil)
	return !os.IsNotExist(err)
}
func getUniqueFilename(filename string) string {
	ext := filepath.Ext(filename)
	base := filename[:len(filename)-len(ext)]
	counter := 1

	for checkFileExists(filename) {
		filename = fmt.Sprintf("%s_%d%s", base, counter, ext)
		counter++
	}

	return filename
}
func UploadEndpoint(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	file, handle, er := r.FormFile("upload")
	if er != nil {
		fmt.Fprintln(w, "File Size Exceeded than 10MB", http.StatusNotAcceptable)
		return
	}
	ext := filepath.Ext(handle.Filename)
	if ext != ".md" {
		fmt.Fprintln(w, "Not a Mark Down File", http.StatusNotAcceptable)
		return
	}
	defer file.Close()
	filename := getUniqueFilename(handle.Filename)
	filpath := filepath.Join("file", filename)
	dat, _ := io.ReadAll(file)
	os.WriteFile(filpath, dat, 0644)
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintln(w, "Uploaded file with file name:-", filename)

}

func ListFilesEndpoint(w http.ResponseWriter, r *http.Request) {
	dir, er := os.ReadDir("file")
	if er != nil {
		fmt.Fprintln(w, "No records")
		return
	}
	var names []string
	for _, fil := range dir {
		names = append(names, fil.Name())
	}
	fmt.Fprintf(w, `{"files":%v}`, names)
}

func CreateHTMLEndpoint(w http.ResponseWriter, r *http.Request) {
	fil := filepath.Join("file", "Readme.md")
	inpt, _ := os.ReadFile(fil)
	ot := blackfriday.Run(inpt)
	fmt.Fprintln(w, string(ot))
}
