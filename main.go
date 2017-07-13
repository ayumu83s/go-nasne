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
	hddList, err := apiClient.Status.HDDListGet(nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("HDDListGet:%+v\n", hddList)
	for _, hdd := range hddList.Hdd {
		hddInfo, err := apiClient.Status.HDDInfoGet(nil, hdd.ID)
		if err != nil {
			panic(err)
		}
		fmt.Printf("HDDInfoGet:%+v\n", hddInfo)
	}
	args := &nasne.RecordedTitleListArgs{
		RequestedCount: 1,
	}
	titleList, err := apiClient.Recorded.TitleListGet(nil, args)
	if err != nil {
		panic(err)
	}
	fmt.Printf("TitleListGet:%+v\n", titleList)

	recNgList, err := apiClient.Status.RecNgListGet(nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ReqNgNum:%d\n", recNgList.Number)
}
