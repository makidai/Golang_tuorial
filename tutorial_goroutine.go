package main 

import( 
	"fmt"
	"log"
	"runtime"
)
func main(){
    fmt.Println("Hello World")
    //ex1:関数(またはメソッド)の呼び出しの前にgoを付けると、異なるgoroutineで関数を実行することが可能。
    go hello()
    go goodMorning()
    //ex2:runtime.NumGoroutine()を使用する事で現在起動しているgoroutine(ゴルーチン)の数を知ることが可能。
    log.Println(runtime.NumGoroutine()) 
    //=> Hello World
    //=> 2021/08/18 23:00:00 3
    //3がgoroutineの数。
    //ここではmain, hello, goodMorningの３つを指す。
}

func hello(){
    fmt.Println("Hello")
}

func goodMorning(){
    fmt.Println("Good Morning")
}