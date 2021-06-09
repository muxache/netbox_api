package netbox_api

import (
	"encoding/json"
	"fmt"

	"github.com/muxache/netbox_api/controller/get_netbox"
	"github.com/muxache/netbox_api/data_model/netbox"
)

//GetALLCircuits allows get all l2circuits from api 'https://netbox.ti.ru/api/circuits/circuits/'
func GetALLDevices(token string) {
	var (
		url        string = "https://netbox.ti.ru/api/dcim/devices/"
		nbDevices netbox.NetBox_Devices_GET
	)
	res := get_netbox.GetToNetBox(url, token)
	for _, r := range res.Results {
		rByte, _ := json.Marshal(r)
		json.Unmarshal(rByte, &nbDevices)
		fmt.Println(nbDevices.Name)
	}
}
