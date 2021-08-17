package main 

import "fmt"

func main(){
    defer fmt.Println("Golang") //defer1
    defer fmt.Println("Ruby") //defer2
    fmt.Println("JS") 
    //=> JS
    //=> Ruby
    //=> Golang
}