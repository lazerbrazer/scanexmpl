package scanning

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
)

type singlePortResp struct {
	Port        int    `json:"port"`
	Description string `json:"description"`
}

func Scan() {
	port := 80
	lastRange := port + 10
	server := "localhost"
	var address string
	//example: address=localhost:80

	var output []singlePortResp
	var portResp singlePortResp
	for i := port; i <= lastRange; i++ {
		address = server + ":" + strconv.Itoa(i)
		_, err := net.Dial("tcp", address)
		portResp.Port = i
		if err == nil {
			portResp.Description = "Closed"
		} else {
			portResp.Description = "Opened"
		}
		output = append(output, portResp)
	}
	//fmt.Printf("output:%+v\n", output)
	PrettyPrint(output)
}

func PrettyPrint(data interface{}) {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}
