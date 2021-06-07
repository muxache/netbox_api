package netbox

import "time"

type Netbox_Circuits_Get struct {
	ID       int    `json:"id"`
	URL      string `json:"url"`
	Cid      string `json:"cid"`
	Provider struct {
		ID   int    `json:"id"`
		URL  string `json:"url"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"provider"`
	Type struct {
		ID   int    `json:"id"`
		URL  string `json:"url"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"type"`
	Status struct {
		Value string `json:"value"`
		Label string `json:"label"`
	} `json:"status"`
	Tenant       interface{} `json:"tenant"`
	InstallDate  string      `json:"install_date"`
	CommitRate   interface{} `json:"commit_rate"`
	Description  string      `json:"description"`
	TerminationA struct {
		ID   int    `json:"id"`
		URL  string `json:"url"`
		Site struct {
			ID   int    `json:"id"`
			URL  string `json:"url"`
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"site"`
		PortSpeed         int         `json:"port_speed"`
		UpstreamSpeed     interface{} `json:"upstream_speed"`
		XconnectID        string      `json:"xconnect_id"`
		ConnectedEndpoint struct {
			ID     int    `json:"id"`
			URL    string `json:"url"`
			Device struct {
				ID          int    `json:"id"`
				URL         string `json:"url"`
				Name        string `json:"name"`
				DisplayName string `json:"display_name"`
			} `json:"device"`
			Name  string `json:"name"`
			Cable int    `json:"cable"`
		} `json:"connected_endpoint"`
		ConnectedEndpointType      string `json:"connected_endpoint_type"`
		ConnectedEndpointReachable bool   `json:"connected_endpoint_reachable"`
	} `json:"termination_a"`
	TerminationZ interface{}   `json:"termination_z"`
	Comments     string        `json:"comments"`
	Tags         []interface{} `json:"tags"`
	CustomFields struct {
	} `json:"custom_fields"`
	Created     string    `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
}
