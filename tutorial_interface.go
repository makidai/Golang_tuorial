package main

import "fmt"

type Person struct {} //Person構造体
type Person2 struct {} //Person2構造体

type People interface{
	intro()
}

func IntroForPerson(arg People) {
	arg.intro();
}

//Person構造体のメソッドintro()
func (p *Person) intro() { 
    fmt.Println("Hello World")
}

//Person2構造体のメソッドintro()
func (p *Person2) intro() {
    fmt.Println("Hello World")
}

//Person構造体のメソッドintro()
// func (p *Person) intro() string{ 
//     return "Hello World"
// }

//Person2構造体のメソッドintro()
// func (p *Person2) intro() string{
//     return "Hello World"
// }

//Person構造体のメソッドintro()を実行するメソッド
// func IntroForPerson(arg *Person){
//     fmt.Println(arg.intro())
// }

//Person2構造体のメソッドintro()を実行するメソッド
// func IntroForPerson2(arg *Person2){
//     fmt.Println(arg.intro())
// }


func main(){
	bob := new(Person)
	mike := new(Person2) 

	IntroForPerson(bob) //=> Hello World
	IntroForPerson(mike) //=> Hello World
}