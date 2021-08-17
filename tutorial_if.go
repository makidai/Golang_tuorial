package main 

import "fmt"

func condition(arg string)string{
	if v := "Go"; arg == v {
		return "This is Golang"
	}else{
		return "This is not Golang"
	}
}
func main(){
    fmt.Println(condition("Python")) //=> This is not Golang
}
