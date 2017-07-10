package main

import (
	"fmt"

	"github.com/ayumu83s/go-nasne/nasne"
)

func main() {
	apiClient, err := nasne.NewClient("192.168.11.3", nil)
	if err != nil {
		fmt.Println(err)
	}
	hddList, err := apiClient.Status.HDDListGet()
	if err != nil {
		panic(err)
	}
	fmt.Printf("HDDListGet:%+v\n", hddList)
	for _, hdd := range hddList.Hdd {
		hddInfo, err := apiClient.Status.HDDInfoGet(hdd.ID)
		if err != nil {
			panic(err)
		}
		fmt.Printf("HDDInfoGet:%+v\n", hddInfo)
	}
	titleList, err := apiClient.Recorded.TitleListGet(0, 0, 0, 1, 0, 1, 0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("TitleListGet:%+v\n", titleList)
}
