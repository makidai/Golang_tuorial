package main

import "fmt"

func main(){
    //参照用配列
    var arr[2]string = [2]string{"Ruby","Golang"}

    //ex1:配列とは異なり長さ指定の必要なし
    var slice1 []string //スライス(1)
    var slice2 []string = []string{"Ruby", "Golang"} //スライス(2)

    //ex2:配列から要素を取り出し参照する形での宣言が可能
    var slice3 = arr[0:2] //スライス(3)

    //ex2:make()を利用した宣言が可能
    var slice4 = make([]string,2,2) //スライス(4)

    //ex3:配列とは異なり要素の追加が可能
    //append は新しいスライスを返すことに注意
    slice5 := [] string{"JavaScript"}
    newSlice := append(slice5, "Ruby") //sliceに"Ruby"を追加

    fmt.Println(slice1,slice2,slice3,slice4, slice5, newSlice) //=>[] [Ruby Golang] [Ruby Golang] [ ] [JavaScript] [JavaScript Ruby]

	slice := []string{"Golang", "Ruby"}

    //ex4:長さ(length)と容量(capacity)の両方を持っている。
    //長さ(length) は、それに含まれる要素の数
    //容量(capacity) は、スライスの最初の要素から数えて、元となる配列の要素数
    fmt.Println(len(slice)) //=> 2
    fmt.Println(cap(slice)) //=> 2

    //ex5:型が一致している場合、他のスライスに代入することが可能。
    var slice6 []string //sliceと同型slice2を作成
    slice6 = slice //sliceをslice2に代入
    fmt.Println(slice6) //=>[Golang Ruby]

    //ex6:スライスのゼロ値は nil
    var slice7 []int
    fmt.Println(slice7, len(slice), cap(slice)) //=> [] 0 0

    if slice7 == nil {
        fmt.Println("nil!") //=> nil! 
        //sliceの値がnilの場合にnil!を表示する。
    }
}