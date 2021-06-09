package netbox

import "time"

type NetBox_Devices_Get struct {
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
	Serial string `json:"serial"`
	Site   struct {
		ID   int    `json:"id"`
		URL  string `json:"url"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"site"`
	Status struct {
		Value string `json:"value"`
		Label string `json:"label"`
	} `json:"status"`
	Comments     string `json:"comments"`
	CustomFields struct {
	} `json:"custom_fields"`
	ConfigContext struct {
	} `json:"config_context"`
	Created     string    `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
}