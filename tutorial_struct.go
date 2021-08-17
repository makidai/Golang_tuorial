package main

import "fmt"

// ex1: typeとstructを使用して定義する。
type Person struct {
	name string 
	age int
}

//ex3:構造体内にメソッドを定義できる
//普通の関数と違うのはレシーバ引数(下記の「(ele Person)」)の部分だけ。
func(ele Person)intro(arg string) string {
    return arg + " I am" + " " + ele.name
}

//ex4:継承に似た機能として構造体の埋め込みが可能
//構造体UserにPersonを組み込む。
type User struct {
    Person
}

func(ele User)intro(arg string) string { //Userのメソッド
    return "User No." + arg + " " + ele.name
}


func main(){
//ex2:複数の初期化方法が存在する
	//初期化方法①:変数定義後にフィールドを設定する
	var mike Person
	mike.name = "Mike"
	mike.age = 23

	//初期化方法②: {} で順番にフィールドの値を渡す
	var bob = Person{"Bob", 35}

	//初期化方法③:フィールド名を ： で指定する方法
	var sam = Person{age:89, name: "Sam"}

	fmt.Println(mike, bob, sam) //=> {Mike 23} {Bob 35} {Sam 89}

	fmt.Println(bob.intro("Hello!")) //=>Hello! I am Bob

	var user1 User

	//組み込みによりnameの定義が可能
	user1.name="Bob"

	fmt.Println(user1.intro("1")) //=> User No.1 Bob
}