package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"MyGeneratorPrime/pirate"
	"MyGeneratorPrime/new_pdf"
	"MyGeneratorPrime/parsecsv"
)

func main() {

	var args []string = os.Args[1:]
	fmt.Println(args)
	if len(args) < 1 {
		log.Fatal("./main -cli / -file [-name -prime -img / -path]")
	}
	if args[0] == "-cli" {
		if len(args) < 7 {
			log.Fatal("not enough arg")
		} else if args[1] != "-name" {
			log.Fatal("error with args order")
		} else if reflect.TypeOf(args[2]) != reflect.TypeOf("test") {
			log.Fatal("error with args type")
		} else if args[3] != "-prime" {
			log.Fatal("error with args order")
		} else if reflect.TypeOf(args[4]) != reflect.TypeOf("test") {
			log.Fatal("error with args type")
		} else if args[5] != "-img" {
			log.Fatal("error with args order")
		} else if reflect.TypeOf(args[6]) != reflect.TypeOf("test") {
			log.Fatal("error with args type")
		} else {
			new_pirate, err_new := pirate.New(args[2], args[4], args[6])
			if err_new != nil {
				fmt.Println(err_new.Error())
				// fmt.Println(new_pirate)
				return
			}
			pdf.New_pdf(new_pirate)
		}
	} else if args[0] == "-file" {
		if len(args) < 3 {
			log.Fatal("not enough arg")
		} else if args[1] != "-path" {
			log.Fatal("wrong args")
		}else{
			parse.Parsecsv(args[2])
		}
	}
}
