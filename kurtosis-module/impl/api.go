package impl

type ExecuteParams struct {
	NumDatastores uint32 `json:"numDatastores"`
}

type ExecuteResult struct {
	CreatedServiceIdsToPortIds map[string]string `json:"createdServiceIdsToPortIds"`
}
