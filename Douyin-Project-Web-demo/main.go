package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	http.HandleFunc("/upload", upload)
	http.Handle("/upload/", http.StripPrefix("/upload/", http.FileServer(http.Dir("upload"))))
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		file, _, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		f, err := os.Create("upload/input.mp4")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		io.Copy(f, file)

		cmd := exec.Command("ffmpeg", "-i", "upload/input.mp4", "-vcodec", "libx264", "-acodec", "aac", "-strict", "-2", "upload/output.mp4")
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			return
		}

		http.Redirect(w, r, "/upload/output.mp4", http.StatusFound)
	}
}
