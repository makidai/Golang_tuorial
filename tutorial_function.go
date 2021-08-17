package main

import "fmt"

func main() {
    fmt.Println(hello("Hello World")) //=> Hello World
    fmt.Println(hello2("Hello", "World")) //=> Hello World
    fmt.Println(multipleArgs("Hello", "World")) //=> World Hello
}

//ex1:引数の戻り値の型を書く必要がある
func hello(arg string) string{
	return arg
}
//ex2:関数が２つ以上の同種類の引数を伴う際、型の省略する事が可能
func hello2(arg1, arg2 string) string {
	return arg1 + " " + arg2
}
//ex3:関数は複数の値を返すことができる
func multipleArgs(arg1, arg2 string)(string, string){
	return arg2, arg1
}
