package main

// 表示本程式所屬的package 名稱，使用 main 表示為可以單獨執行之程式

import "fmt"

// 引用套件，類似Ｃ的#include <stdio.h>

func main() {
	// 宣告main function，在Golang之中main function可以不用任何參數。
	fmt.Printf("Hello world!!\n")
	// 印出Hello world
	// 在Golang裡，變數命名的大小寫會影響變數的可視範圍(Scope)
	// 以大寫為開頭的變數、function、類別表示著可以在local package之外被看見
}
