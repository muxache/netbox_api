package netbox

import (
	"time"
)

type NetBox_Devices_GET struct {
	ID          int    `json:"id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	DeviceType  struct {
		ID           int    `json:"id"`
		URL          string `json:"url"`
		Manufacturer struct {
			ID   int    `json:"id"`
			URL  string `json:"url"`
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"manufacturer"`
		Model       string `json:"model"`
		Slug        string `json:"slug"`
		DisplayName string `json:"display_name"`
	} `json:"device_type"`
	DeviceRole struct {
		ID   int    `json:"id"`
		URL  string `json:"url"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"device_role"`
	Tenant   interface{} `json:"tenant"`
	Platform interface{} `json:"platform"`
	Serial   string      `json:"serial"`
	AssetTag interface{} `json:"asset_tag"`
	Site     struct {
		ID   int    `json:"id"`
		URL  string `json:"url"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"site"`
	Rack         interface{} `json:"rack"`
	Position     interface{} `json:"position"`
	Face         interface{} `json:"face"`
	ParentDevice interface{} `json:"parent_device"`
	Status       struct {
		Value string `json:"value"`
		Label string `json:"label"`
	} `json:"status"`
	PrimaryIP        interface{}   `json:"primary_ip"`
	PrimaryIP4       interface{}   `json:"primary_ip4"`
	PrimaryIP6       interface{}   `json:"primary_ip6"`
	Cluster          interface{}   `json:"cluster"`
	VirtualChassis   interface{}   `json:"virtual_chassis"`
	VcPosition       interface{}   `json:"vc_position"`
	VcPriority       interface{}   `json:"vc_priority"`
	Comments         string        `json:"comments"`
	LocalContextData interface{}   `json:"local_context_data"`
	Tags             []interface{} `json:"tags"`
	CustomFields     struct {
	} `json:"custom_fields"`
	ConfigContext struct {
	} `json:"config_context"`
	Created     string    `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
}