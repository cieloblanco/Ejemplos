package main

import (
	"net/http"
	"os/exec"
	"log"
	"html/template"
	"io/ioutil"
	"strings"
	"bytes"
	"fmt"
)

func main() {
	http.HandleFunc("/code", codeHandler)
	http.HandleFunc("/out", outHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Data struct {
	Code []byte
	Out string
}

func codeHandler(w http.ResponseWriter, r *http.Request) {

	//code.py a estructura Data
	code, _ := ioutil.ReadFile("python3.py")
	d := &Data{Code: code}

	//html
	t, _ := template.ParseFiles("code.html")
	t.Execute(w, d)
}

func outHandler(w http.ResponseWriter, r *http.Request) {

	code := r.FormValue("code")
	ioutil.WriteFile("python3.py", []byte(code), 0600)

	//output
	var out = ""
	cmd := exec.Command("python3", "python3.py")
	stdoutStderr, _ := cmd.CombinedOutput()
	out += string(stdoutStderr)

	//html

	d := &Data{Code: []byte(code), Out: out}
	buf := &bytes.Buffer{}

	t, _ := template.ParseFiles("out.html")
	t.Execute(buf, d)

	body := buf.String()
	body = strings.Replace(body, "\n", "<br>", -1)
	fmt.Fprint(w, body)
}















