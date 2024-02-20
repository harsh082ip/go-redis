package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func getVisits(w http.ResponseWriter, r *http.Request) {

	visits, err := client.Get(ctx, "visits").Result()
	if err != nil {
		fmt.Println(err)
	}

	visitsInt, err := strconv.Atoi(visits)
	if err != nil {
		fmt.Println(err)
	}

	visitsInt++

	client.Set(ctx, "visits", strconv.Itoa(visitsInt), 0)

	fmt.Fprintf(w, "No of visits is %v", visitsInt)

}
