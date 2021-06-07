package netbox

type Netbox_Circuits_PUT struct {
	ID       int    `json:"id"`
	URL      string `json:"url"`
	Cid      string `json:"cid"`
	Provider struct {
		ID int `json:"id"`
	} `json:"provider"`
	Type struct {
		ID int `json:"id"`
	} `json:"type"`

	Description string `json:"description"`

	Comments     string `json:"comments"`
	CustomFields struct {
	} `json:"custom_fields"`
}
