package main

import "fmt"

func main() {
    var lado int32
    fmt.Scanf("%d", &lado)

    area := lado * lado

    fmt.Println(area)
}
