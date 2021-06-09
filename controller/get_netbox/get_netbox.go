package get_netbox

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"os"

	"github.com/muxache/netbox_api/data_model/netbox"
)

//GetFromNetBox returns netbox struct with any contained data.
func GetFromNetBox(url, token string) netbox.Netbox_Struct {
	var (
		limit  int
		newUrl string
		nb     netbox.Netbox_Struct
	)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", token)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error when sending request to the server")
		os.Exit(1)
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&nb)
	if len(nb.Next) != 0 {
		limit, _ = strconv.Atoi(URLParse(nb.Next)["limit"][0])
		newUrl = nb.Next
		for i := limit; i <= nb.Count; i += limit {

			reqnext, _ := http.NewRequest("GET", newUrl, nil)
			reqnext.Header.Add("accept", "application/json")
			reqnext.Header.Add("Authorization", token)
			respnext, err1 := client.Do(reqnext)
			if err1 != nil {
				fmt.Println("Error in "+i+" when sending request to the server")
				os.Exit(1)
			}
			defer respnext.Body.Close()
			var pn netbox.Netbox_Struct
			json.NewDecoder(respnext.Body).Decode(&pn)
			newUrl = pn.Next
			nb.Results = append(nb.Results, pn.Results...)
		}
	}
	return nb
}

func URLParse(urlField string) url.Values {
	u, _ := url.Parse(urlField)
	m, _ := url.ParseQuery(u.RawQuery)
	return m
}

func URLSet(urlField, newLimit, offset string) string {
	u, _ := url.Parse(urlField)
	q := u.Query()
	q.Set("limit", newLimit)
	q.Set("offset", offset)
	u.RawQuery = q.Encode()
	return u.String()
}
