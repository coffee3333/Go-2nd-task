package main

import (
	"html/template"
	"log"
	"net/http"
	// "text/template"
)

func mainPage(w http.ResponseWriter, r *http.Request) { //func for main page
	tmpl, _ := template.ParseFiles("temp/index.html")
	tmpl.Execute(w, nil)
	// fmt.Fprintf(w, "Привет!")
}

func routes() {
	http.HandleFunc("/", mainPage) //Added route for main page
}

func handleFunc() {
	http.Handle("/temp/", http.StripPrefix("/temp/", http.FileServer(http.Dir("./temp/")))) //Here i am showing direction to temp files, especialy for css
	routes()                                                                                //as you know, we can have many routes, i deceided to make func for this
	err := http.ListenAndServe(":8080", nil)                                                // configuring the port
	log.Println("Server is running!")
	if err != nil { //excepting errors
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() { //The entry point
	handleFunc()
}
