package main

import (
	"fmt"
	"io/ioutil"
	"github.com/whyakari/httpbinip/ip"
)

func main() {
    ip, err := httpbin.GetIp()

    if err != nil {
        fmt.Println("Error getting IP:", err)
    }

    err = ioutil.WriteFile("list.txt", []byte(ip), 0644) 
    
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("IP saved to list.txt")

}
