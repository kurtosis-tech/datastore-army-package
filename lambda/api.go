package lambda

type Params struct {
	NumDatastores uint32 `json:"numDatastores"`
}

type Result struct {
	CreatedServiceIdsSet map[string]bool 	`json:"createdServiceIdsSet"`
}
