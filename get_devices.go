package netbox_api

import (
	"encoding/json"

	"github.com/muxache/netbox_api/controller/get_netbox"
	model "github.com/muxache/netbox_api/data_model/netbox"
)

//GetALLCircuits allows get all devices from api 'https://netbox.ti.ru/api/dcim/devices/'
func GetDevices(url, token string) []model.NetBox_Devices_Get {
	var (
		nbDevices []model.NetBox_Devices_Get
	)
	res := get_netbox.GetFromNetBox(url, token)
	for _, r := range res.Results {
		var nb model.NetBox_Devices_Get
		rByte, _ := json.Marshal(r)
		json.Unmarshal(rByte, &nb)
		nbDevices = append(nbDevices, nb)
	}
	return nbDevices
}
