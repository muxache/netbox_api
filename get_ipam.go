package netbox_api

import (
	"encoding/json"

	"github.com/muxache/netbox_api/controller/get_netbox"
	model "github.com/muxache/netbox_api/data_model/netbox"
)

func GetIPAM(url, token string) []model.NetBox_IPAM_Get {
	var (
		nbIPAM []model.NetBox_IPAM_Get
	)
	res := get_netbox.GetFromNetBox(url, token)
	for _, r := range res.Results {
		var nb model.NetBox_IPAM_Get
		rByte, _ := json.Marshal(r)
		json.Unmarshal(rByte, &nb)
		nbIPAM = append(nbIPAM, nb)
	}
	return nbIPAM
}
