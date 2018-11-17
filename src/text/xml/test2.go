package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type server2 struct{
	ServerName string `xml:"serverName"`
	ServerIP string `xml:"serverIP"`
}
type Servers struct{
	XMLName xml.Name `xml:"servers"`
	Version string `xml:"version,attr"`
	Svs []server2 `xml:"server"`
}

func main()  {
	v := &Servers{Version:"1"}
	v.Svs = append(v.Svs,server2{"Shanghai_VPN","127.0.0.1"})
	v.Svs = append(v.Svs,server2{"Beijing_VPN","127.0.0.2"})
	//output,err := xml.Marshal(v)
	output,err := xml.MarshalIndent(v,""," ")
	if err != nil{
		fmt.Println(err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}



