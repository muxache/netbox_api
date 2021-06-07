package cadc

type CADC_Vacant_Target_Struct struct {
	ID                string
	VacantT           string
	OxdResults        []OxdResults
	LastTargetId      string
	LastTargetDesc    string
	LastTargetComment string
	LastTargetName    string
	LastUpdated       string
	LastUpdatedUser   string
	LastCreated       string
	LastCreatedUser   string
	VacantBit         bool
	VacantBitOxd      bool
	Error             string
	RRData            []RRData
}

type OxdResults struct {
	OxdInt      string
	OxdHostName string
}

type RRData struct {
	Origin string
	Age    string
	Site   string
	Rd     string
}
