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
	WORKERS  = 4
)

func main() {
	//load data to be requested
	names, err := csv.Read("names.csv", ',')
	if err != nil {
		fmt.Println(err)
	}

	//login on API
	start := time.Now()
	login, err := requests.Login(EMAIL, PASSWORD)
	if err != nil {
		return
	}

	//request API
	results, err := requests.AsyncAPIRequest(names, EMAIL, PASSWORD, WORKERS, login.Value)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("finished in ", time.Since(start))

	//download API response to files
	err = csv.Write("result", "result", results)
	if err != nil {
		fmt.Println(err)
	}
}
