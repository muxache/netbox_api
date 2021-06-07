package netbox_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/muxache/netbox_api/data_model/cadc"
)

func GetVacantVPLSTarget() (cadc.CADC_Vacant_Target_Struct, error) {
	defer recover()
	var (
		url string
		a   cadc.CADC_Vacant_Target_Struct
	)
	client := http.Client{}

	url = "http://noc.ti.ru:8887/json/vacant/vpls"

	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error when sending request to the server")
		panic(err)
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&a)

	if a.VacantT == "" {
		return a, errors.New("no vacant target")
	}
	return a, nil

}
