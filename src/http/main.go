package main

import (
	"fmt"
	"log" 
	"net/http"
)

var number float32 = 5.3
var name, email string = "vikram", "v.gopu@sotec.eu"
func main() {
	y:= "This is my first go-project";
	x:=&y;
	fmt.Printf("%v\n", *x);
	fmt.Printf("%T", number);
	fmt.Printf("nameType: %T, emailType: %T", name, email);
}

func initHttpServer() {
	http.HandleFunc("/", requestHandler);
	log.Fatal("Server Already in use", http.ListenAndServe(":8000", nil));
}
func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my website!!");
}