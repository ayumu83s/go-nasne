package main

import (
	"fmt"

	"github.com/ayumu83s/go-nasne-api-client/nasne"
)

func main() {
	apiClient, err := nasne.NewClient("192.168.11.3", nil)
	if err != nil {
		fmt.Println(err)
	}
	res, err := apiClient.Status.HDDListGet()
	if err != nil {
		panic(err)
	}
	fmt.Printf("HDDListGet:%+v\n", res)
	for _, hdd := range res.Hdd {
		res, err := apiClient.Status.HDDInfoGet(hdd.ID)
		if err != nil {
			panic(err)
		}
		fmt.Printf("HDDInfoGet:%+v\n", res)
	}
}
