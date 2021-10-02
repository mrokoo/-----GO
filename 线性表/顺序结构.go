package main

import "fmt"

const MAXSIZE = 10

type Book struct {
	id    string
	name  string
	price float64
}

type Sqlist struct {
	Book   *[MAXSIZE]Book // GO中将长度+类型单独作为一种类型存在。
	length int
}

func InitSqlist(booklist *Sqlist) string {
	booklist.Book = new([MAXSIZE]Book)

	if booklist.Book == nil {
		return "failed"
	}
	booklist.length = 0
	return "OK"
}

func main() {
	booklist := new(Sqlist)
	stauts := InitSqlist(booklist)
	fmt.Println(stauts)
}
