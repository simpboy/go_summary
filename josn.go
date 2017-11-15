package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	title string
	Year  int	`json:"released"`
	Color bool	`json:"color,omitempty"`
	Actors	[]string
}
var movies = []Movie{
	{title:"Casablanca",Year:1942,Color:false,Actors:[]string{"Humphrey Bogart","Ingrid Bergman"}},
	{title: "Cool Hand Luke", Year: 1967, Color: true,Actors: []string{"Paul Newman"}},
	{title: "Bullitt", Year: 1968, Color: true,Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main() {
	data,err := json.Marshal(movies);
	if(err!=nil){
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	data_indent,err := json.MarshalIndent(movies,""," ")
	if(err!=nil){
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data_indent)


	var movies_Unmarshal []Movie
	if err :=json.Unmarshal([]byte(data_indent),&movies_Unmarshal);err != nil{
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(movies_Unmarshal)
	for _,movie := range movies_Unmarshal{
		fmt.Println(movie)
	}
}
