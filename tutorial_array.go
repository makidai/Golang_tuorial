package main

import "fmt"

// ① var 変数名 [長さ]型
// ② var 変数名 [長さ]型 = [大きさ]型{初期値1, 初期値n} 
// ③ 変数名 := [...]型{初期値１, 初期値n}

//ex1:配列とは、同じ型を持つ値（要素）を並べたもの
//ex2:複数の宣言方法がある

//宣言方法(1)
var arr1 [2]string 
//宣言方法(2)
var arr2 [2]string = [2]string{"Golang", "Ruby"}
//宣言方法(3)
var arr3 = [...]string{"Golang", "Ruby"}

func main(){
    arr1[0] = "Golang"
    arr1[1] = "Ruby"
    fmt.Println(arr1, arr2, arr3) //=> [Golang Ruby] [Golang Ruby] [Golang Ruby]
}
