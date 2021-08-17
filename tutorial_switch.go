package main

import "fmt"

func main(){
    fmt.Println(condition("Go")) //This is Go
}

func condition(arg string) string{
    switch arg {
    case "Ruby":
        return "This is Ruby"
    case "Go": //これ以降のcaseやdefaultは実行されない。
        return "This is Go"
    case "JS":
        return "This is JS"
    default:
        return "I don't know what this is"
    }
}