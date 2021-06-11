package netbox_api

import (
	"encoding/json"

	"github.com/muxache/netbox_api/controller/get_netbox"
	"github.com/muxache/netbox_api/data_model/netbox"
)

//GetALLCircuits allows get all l2circuits from api 'https://netbox.ti.ru/api/circuits/circuits/'
func GetALLCircuits(url, token string) []netbox.Netbox_Circuits_Get {
	var (
		//url        string = "https://netbox.ti.ru/api/circuits/circuits/"

		resArr []netbox.Netbox_Circuits_Get
	)
	res := get_netbox.GetFromNetBox(url, token)
	for _, r := range res.Results {
		var nbCircuits netbox.Netbox_Circuits_Get
		rByte, _ := json.Marshal(r)
		json.Unmarshal(rByte, &nbCircuits)
		resArr = append(resArr, nbCircuits)
	}
	return resArr
}
