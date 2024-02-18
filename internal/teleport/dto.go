package teleport

type DatabaseConfigResult struct {
	Name     string
	Host     string
	Port     int
	User     string
	Database string
	Ca       string
	Cert     string
	Key      string
}

type DatabaseListResult []DatabaseListItem

type DatabaseListItem struct {
	Metadata DatabaseListResultMetadata
}

type DatabaseListResultMetadata struct {
	Name string
}
