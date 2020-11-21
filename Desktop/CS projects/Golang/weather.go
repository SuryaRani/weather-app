package main

import (
	"fmt"
	"log"
	"time"
	//"html/template"
	"net/http"
	
)

 type coordinates struct{
	lat int
	long int
}


func main(){
	fileServer := http.FileServer(http.Dir("./static")) 
	http.Handle("/", fileServer)
	http.Handle("/form.html", fileServer)

	//http.HandleFunc("/",HomePage)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/weather", WeatherPage)
	log.Fatal(http.ListenAndServe(":8000",nil))

}

func HomePage(w http.ResponseWriter, r *http.Request){
	fmt.Println("Home page accessed")

	now:= time.Now()

	fmt.Fprintf(w,"Welcome!\n")
	fmt.Fprintf(w,"It is %s\n",now)
	fmt.Fprint(w, "What would you like to do\n")
	//comment

}

func WeatherPage(w http.ResponseWriter, r *http.Request){

	fmt.Println("Weather page accessed")



	fmt.Fprintf(w,"Where would you like to know the weather of?\n")

}
func formHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("lat")
    address := r.FormValue("long")

    fmt.Fprintf(w, "Name = %s\n", name)
    fmt.Fprintf(w, "Address = %s\n", address)
    
}
