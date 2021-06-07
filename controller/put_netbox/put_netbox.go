package put_netbox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func MethodToNetbox(url, method, token string, cc interface{}) interface{} {

	jsonStr, _ := json.Marshal(cc)
	body := bytes.NewReader([]byte(jsonStr))

	req, err := http.NewRequest(method, url, body)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error when sending request to the server")
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&cc)
	return cc
}
