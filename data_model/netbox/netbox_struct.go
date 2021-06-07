package netbox

type Netbox_Struct struct {
	Count    int           `json:"count"`
	Next     string        `json:"next"`
	Previous interface{}   `json:"previous"`
	Results  []interface{} `json:"results"`
}
