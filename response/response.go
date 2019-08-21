package response

// a struct for the metadata contained in the JSON
type MetaJSON struct {
	Runtime   string // not the most efficent way to convey this...
	Errors    []string
	Message   string `json:"Message"` // custom message to send
	FileNames []string `json:"FileNames"`// list of filenames scanned
	TOS       string `json:"TOS"`
}

// a struct for the response json
type ResponseJSON struct {
	Meta   MetaJSON
	FDNS_A []string
	RDNS   []string
}
