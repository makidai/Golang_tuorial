package main 

import "fmt"

//ex1:`var`によって変数を宣言し、型の明示が必要である。
var lang string 
//※変数に初期値を与えないとゼロ値（Zero values)が設定されます。(数値型には0、bool型にはfalse、string型には""(空の文字列)が与えられる。

//ex2:初期値を渡した状態で変数を宣言すると型の明示を省略が可能。
var lang2 = "Golang"//初期値を渡す。

func main() {
    //ex3:関数内では｀:=`を利用することでより短いコードで変数の宣言を行うことが可能
    lang3 := "JS" 
    fmt.Println(lang, lang2, lang3) //=> Golang JS
}
