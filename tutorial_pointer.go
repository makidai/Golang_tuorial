package main

import "fmt"

func main(){
    var lang string 
    lang = "Go" 
    //ex1:`&`を用いることで変数のアドレスの取得が可能
    fmt.Println(&lang) //=> 0xc000096220

    //ex2:`*`を使用する事でポインタを値にとったポインタ変数の宣言が可能
    var langP *string 
    langP = &lang
    fmt.Println(langP)//=> 0xc000096220

    //ex3:ポインタ型変数名の前に`*`をつけることで変数の中身へのアクセスが可能
    fmt.Println(*langP) //=> Go
}
