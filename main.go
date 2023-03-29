package main

import (
	"fmt"
	"github.com/DanielFillol/API_Caller/csv"
	"github.com/DanielFillol/API_Caller/requests"
	"time"
)

const (
	EMAIL    = "root@root.com"
	PASSWORD = "JI1O2J2O1U09WURPIKHJFLSDKJWOIURJWOIUXJIOQWMKJX"
	WORKERS  = 10
)

func main() {
	//CSV read
	names, err := csv.Read("names.csv", ',')
	if err != nil {
		fmt.Println(err)
	}

	start := time.Now()
	results, err := requests.AsyncAPIRequest(names, EMAIL, PASSWORD, WORKERS)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("finished in ", time.Since(start))

	//CSV write
	err = csv.Write("result.csv", "result", results)
	if err != nil {
		fmt.Println(err)
	}
}
