package main

import (
	"fmt"
)

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

func GetElem(booklist *Sqlist, i int) (Book, error) {
	if i < 0 || i > MAXSIZE { // golang会给所有未初始化的变量赋予零值。
		return Book{}, fmt.Errorf("%v索引不符合要求", i)
	}

	return booklist.Book[i], nil
}

func main() {
	booklist := new(Sqlist)
	stauts := InitSqlist(booklist)
	fmt.Println(stauts)

	fmt.Println(GetElem(booklist, 2))

	booklist.Book[2] = Book{id: "123456", name: "sf", price: 12.5}

	if i, err := GetElem(booklist, 2); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(i)
	}
}
