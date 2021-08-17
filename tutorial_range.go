package main

import "fmt"

//スライスとマップの作成 
var slice1 = []string{"Golang", "Ruby"}
var map1 = map[string]string{"Lang1":"Golang", "Lang2":"Ruby"}

func main(){
    //ex1:スライスやマップに使用すると反復毎に2つの変数を返す。
    //ex2:スライスの場合、1つ目の変数は `インデックス(index)`で、`2つ目は要素(value)`である。
    for index, value := range slice1{
        fmt.Println(index,value)
        //=> 0 Golang
        //=> 1 Ruby
    }

    //ex3:マップの場合、1つ目の変数は`キー(key)`で、２つ目の変数は`バリュー(value)`である。
    for key, value := range map1{
        fmt.Println(key, value)
        //=> Lang1 Golang
        //=> Lang2 Ruby
    }

    //ex4:インデックスや値は、 _ へ代入することで省略することが可能。

    for _,value := range map1{
        fmt.Println(value)
        //=> Golang
        //=> Ruby
    }

}