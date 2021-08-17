package main

import "fmt"

func main(){
    //ex1:組み込み関数`make()`を使用する事で生成可能。
    ch := make(chan string) 

    //ex2:<-を用いる事で値の送受信が可能。
    //ex2+:作成したチャネルchに値を送信。
    go func(){ ch <- "str"}()

    //ex2+:チャネルchから値を受信
    msg := <- ch
    fmt.Println(msg) //=>str

	//ex3:ゴルーチンが予期しない状態とならないことを保証する。

    ch2 := make(chan bool) //bool型のチャネルchを作成

    // ゴルーチンとして以下の関数を起動。
    // 完了時にchannelの型であるboolの値を送信する事で、チャネルへ通知。
    go func(){
        fmt.Println("Hello")
        ch2 <- true // 通知を送信。値は何でも良い(boolの型であれば)
    }()

    <- ch2 //=>Hello  
    // channelの型であるboolの値を受け取るまでの完了待ち。送られてきた値は破棄
}
