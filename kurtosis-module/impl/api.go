package impl

type ExecuteParams struct {
	NumDatastores uint32 `json:"numDatastores"`
}

type ExecuteResult struct {
	CreatedServiceIdPorts map[string]uint32 `json:"createdServiceIdPorts"`
}
