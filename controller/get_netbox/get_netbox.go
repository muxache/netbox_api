package get_netbox

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	model "github.com/muxache/netbox_api/data_model/netbox"
)

type HttpResponse struct {
	URL      string
	Response *http.Response
	Error    error
	Data     model.Netbox_Struct
}

//GetFromNetBox returns netbox struct with any contained data.
func GetFromNetBox(url, token string) model.Netbox_Struct {
	var (
		urlList []string
	)

	res := get(url, token)
	if res.Count == 0 {
		return model.Netbox_Struct{}
	}
	if res.Next == "" {
		return res
	}

	for i := 0; i < (res.Count / 64); i++ {

		urlList = append(urlList, urlSet(res.Next, "64", strconv.Itoa(64*i)))
	}
	return asyncHttpGets(urlList, token)

}

func get(url, token string) model.Netbox_Struct {
	var nb model.Netbox_Struct
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", token)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error when sending request to the server")
		os.Exit(1)
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&nb)
	return nb
}

//asyncHttpGets makes assync get requests
func asyncHttpGets(urls []string, token string) model.Netbox_Struct {
	var nball model.Netbox_Struct
	var urlsss []string
	ch := make(chan *HttpResponse)
	client := http.Client{}
	for _, url := range urls {
		go func(url string) {
			var nb model.Netbox_Struct
			req, _ := http.NewRequest("GET", url, nil)
			req.Header.Add("accept", "application/json")
			req.Header.Add("Authorization", token)
			resp, err := client.Do(req)
			json.NewDecoder(resp.Body).Decode(&nb)
			ch <- &HttpResponse{url, resp, err, nb}
			if err != nil && resp != nil && resp.StatusCode == http.StatusOK {
				resp.Body.Close()
			}
		}(url)
	}

	for {
		select {
		case r := <-ch:
			nball.Results = append(nball.Results, r.Data.Results...)
			urlsss = append(urlsss, r.URL)
			if len(urlsss) == len(urls) {
				return nball
			}
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func urlSet(urlField, newLimit, offset string) string {
	u, _ := url.Parse(urlField)
	q := u.Query()
	q.Set("limit", newLimit)
	q.Set("offset", offset)
	u.RawQuery = q.Encode()
	return u.String()
}
