package main

import (
	"fmt"
	"log"
	"time"
	"html/template"
	"net/http"
	
)


func main(){
	fileServer := http.FileServer(http.Dir("./static")) 
	http.Handle("/", fileServer)
	
	http.HandleFunc("/",HomePage)
	http.HandleFunc("/weather", WeatherPage)
	log.Fatal(http.ListenAndServe(":8000",nil))

}

func HomePage(w http.ResponseWriter, r *http.Request){
	fmt.Println("Home page accessed")

	now:= time.Now()

	fmt.Fprintf(w,"Welcome!\n")
	fmt.Fprintf(w,"It is %s\n",now)
	fmt.Fprint(w, "What would you like to do\n")

}

func WeatherPage(w http.ResponseWriter, r *http.Request){

	fmt.Println("Weather page accessed")

	fmt.Fprintf(w,"Where would you like to know the weather of?\n")

}
