package netbox_api

import (
	"encoding/json"
	"fmt"

	"github.com/muxache/netbox_api/controller/put_netbox"
	"github.com/muxache/netbox_api/data_model/netbox"
)

func PostCircuitToNetbox(target, description, comments, token string) {
	var nt netbox.Netbox_Circuits_PUT

	nt.Cid = target
	nt.Type.ID = 1
	nt.Provider.ID = 2
	nt.Description = description
	nt.Comments = comments

	postURL := "https://netbox.ti.ru/api/circuits/circuits/"

	res := put_netbox.MethodToNetbox(postURL, "POST", token, nt)
	resByte, _ := json.Marshal(res)
	var ntRes netbox.Netbox_Circuits_Get
	json.Unmarshal(resByte, &ntRes)

	fmt.Println(ntRes.URL)
}
