package main

import "fmt"

func main(){
    //ex1複数の宣言方法が存在する

    //①組み込み関数make()を利用して宣言
    //make(map[キーの型]値の型, キャパシティの初期値)
    //make(map[キーの型]値の型)

    map1 := make(map[string]string)
    map1["Name"] = "Mike"
    map1["Gender"] = "Male"

    //②初期値を指定して宣言
    //var 変数名 map[key]value = map[key]value{key1: value1, key2: value2, ..., keyN: valueN}
    map2 := map[string]int{"Age":25,"UserId":2}

    //ex2初期値を指定しない場合、変数はnil(nil マップ)に初期化される
    var map3 map[string]string

    fmt.Println(map1, map2, map3) //=>map[Gender:Male Name:Mike] map[Age:25 UserId:2] map[]

	//ex3:要素の挿入や削除が行える

    //連想配列の作成
    var mapEx = map[int]string{1:"Go",2:"Ruby"}
    fmt.Println(mapEx) //=> map[2:Ruby 1:Go]

    //要素の挿入や更新
    mapEx[3] = "Javascript"
    fmt.Println(mapEx) //=> map[1:Go 2:Ruby 3:Javascript]

    //要素の取得
    fmt.Println(mapEx[2]) //=> Ruby

    //要素の削除
    delete(mapEx,3)
    fmt.Println(mapEx) //=> map[2:Ruby 1:Go]
}