package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Имя: Илья")
	fmt.Println("Дата:", time.Now().Format("02.01.2006"))
}
