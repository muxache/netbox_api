package netbox

import "time"

type NetBox_IPAM_Get struct {
	ID      int    `json:"id"`
	URL     string `json:"url"`
	Display string `json:"display"`
	Family  struct {
		Value int    `json:"value"`
		Label string `json:"label"`
	} `json:"family"`
	Prefix string      `json:"prefix"`
	Site   interface{} `json:"site"`
	Vrf    interface{} `json:"vrf"`
	Tenant interface{} `json:"tenant"`
	Vlan   interface{} `json:"vlan"`
	Status struct {
		Value string `json:"value"`
		Label string `json:"label"`
	} `json:"status"`
	Role         interface{}   `json:"role"`
	IsPool       bool          `json:"is_pool"`
	Description  string        `json:"description"`
	Tags         []interface{} `json:"tags"`
	CustomFields struct {
	} `json:"custom_fields"`
	Created     string    `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
	Children    int       `json:"children"`
	Depth       int       `json:"_depth"`
}
