package main

import (
	"fmt"
	"time"
)

func main() {
	presentDate := time.Now()

	fmt.Println("Present date is: ", presentDate)

	fmt.Println("Present date is: ", presentDate.Format("2006-01-02 15:04:05 Monday"))

}
