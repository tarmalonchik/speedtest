package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

func main() {
	raw, err := base64.StdEncoding.DecodeString("W3snbm9kZV9uYW1lJzogJ2ZpbjEnLCAnbGFiZWxfbmFtZSc6ICdmaW4xJywgJ3NlcnZlcl9jb3VudHJ5JzogJ0ZpbmxhbmQg8J+Hq/Cfh64nLCAnc2VydmVyX2NhcGFjaXR5JzogJzEwMCcsICdlbmFibGVfc2VuZGVyJzogJ3RydWUnLCAn │\n│ bG93X3ByaW9yaXR5JzogJ2ZhbHNlJywgJ2FscGhhXzNfY29kZSc6ICdmaW4nLCAnZW5hYmxlX2lwX3RyYWNrZXInOiAndHJ1ZScsICdpc19pbmZyYXN0cnVjdHVyZSc6IEZhbHNlLCAnaXBfYWRkcmVzcyc6ICc5NS4yMTcuNi4yNTMnLCAnc3BlZWR0ZXN0X2lzX3NlcnZlcic6IC │\n│ dmYWxzZScsICdwcm92aWRlcic6ICdoZXR6bmVyJ30sIHsnbm9kZV9uYW1lJzogJ3VzYTEnLCAnbGFiZWxfbmFtZSc6ICd1c2ExJywgJ3NlcnZlcl9jb3VudHJ5JzogJ1VTQSDwn4e68J+HuCcsICdzZXJ2ZXJfY2FwYWNpdHknOiAnMTAwJywgJ2VuYWJsZV9zZW5kZXInOiAndHJ1 │\n│ ZScsICdsb3dfcHJpb3JpdHknOiAnZmFsc2UnLCAnYWxwaGFfM19jb2RlJzogJ3VzYScsICdlbmFibGVfaXBfdHJhY2tlcic6ICd0cnVlJywgJ2lzX2luZnJhc3RydWN0dXJlJzogRmFsc2UsICdpcF9hZGRyZXNzJzogJzUuNzguMTE2LjgxJywgJ3NwZWVkdGVzdF9pc19zZXJ2ZX │\n│ InOiAnZmFsc2UnLCAncHJvdmlkZXInOiAnaGV0em5lcid9LCB7J25vZGVfbmFtZSc6ICcnLCAnbGFiZWxfbmFtZSc6ICcnLCAnc2VydmVyX2NvdW50cnknOiAnJywgJ3NlcnZlcl9jYXBhY2l0eSc6ICcnLCAnZW5hYmxlX3NlbmRlcic6ICcnLCAnbG93X3ByaW9yaXR5JzogJycs │\n│ ICdhbHBoYV8zX2NvZGUnOiAnJywgJ2VuYWJsZV9pcF90cmFja2VyJzogJycsICdpc19pbmZyYXN0cnVjdHVyZSc6IFRydWUsICdpcF9hZGRyZXNzJzogJzk1LjIxNi4yMTguMTc5JywgJ3NwZWVkdGVzdF9pc19zZXJ2ZXInOiAndHJ1ZScsICdwcm92aWRlcic6ICdoZXR6bmVyJ3 │\n│ 1d")
	fmt.Println(string(raw))
	fmt.Println(err)
	var t []T2

	err = json.Unmarshal(raw, &t)
	fmt.Println(err)
	fmt.Println(t)
}

type T2 struct {
	NodeName          string `json:"node_name"`
	LabelName         string `json:"label_name"`
	ServerCountry     string `json:"server_country"`
	ServerCapacity    string `json:"server_capacity"`
	EnableSender      string `json:"enable_sender"`
	LowPriority       string `json:"low_priority"`
	Alpha3Code        string `json:"alpha_3_code"`
	EnableIpTracker   string `json:"enable_ip_tracker"`
	IsInfrastructure  bool   `json:"is_infrastructure"`
	IpAddress         string `json:"ip_address"`
	SpeedtestIsServer string `json:"speedtest_is_server"`
	Provider          string `json:"provider"`
}
