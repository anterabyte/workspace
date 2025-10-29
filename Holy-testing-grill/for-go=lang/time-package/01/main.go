package main

import (
	"fmt"
	"time"
)


func main() {

	twoHour := 2 * time.Hour
	thirtyMinutes := 30 * time.Minute
	thirtySeconds := 30 * time.Second

	period := twoHour + thirtyMinutes + thirtySeconds

	fmt.Println("Total time period: ",period)
	

}
