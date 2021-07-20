package lambda

type Params struct {
	NumDatastores uint32 `json:"numDatastores"`
}

type Result struct {
	CreatedServiceIdPorts map[string]uint32 `json:"createdServiceIdPorts"`
}
