package netbox_api

import (
	"encoding/json"
	"fmt"

	"github.com/muxache/netbox_api/controller/get_netbox"
	"github.com/muxache/netbox_api/data_model/netbox"
)

func GetALLCircuits(token string) {
	var (
		url        string = "https://netbox.ti.ru/api/circuits/circuits/"
		nbCircuits netbox.Netbox_Circuits_Get
	)
	res := get_netbox.GetToNetBox(url, token)
	for _, r := range res.Results {
		rByte, _ := json.Marshal(r)
		json.Unmarshal(rByte, &nbCircuits)
		fmt.Println(nbCircuits.Cid)
	}
}
