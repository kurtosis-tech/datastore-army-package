package impl

import (
	"encoding/json"
	"fmt"
	"github.com/kurtosis-tech/kurtosis-client/golang/lib/networks"
	"github.com/kurtosis-tech/kurtosis-client/golang/lib/services"
	"github.com/palantir/stacktrace"
)

const (
	datastoreImage = "kurtosistech/example-microservices_datastore"
	datastorePortNumber uint32 = 1323
	datastoreProtocol = "tcp"
)

type DatastoreArmyKurtosisModule struct {
	numDatstoresAdded int
}

func NewDatastoreArmyKurtosisModule() *DatastoreArmyKurtosisModule {
	return &DatastoreArmyKurtosisModule{}
}

func (module *DatastoreArmyKurtosisModule) Execute(networkCtx *networks.NetworkContext, serializedParams string) (serializedResult string, resultError error) {
	params := new(ExecuteParams)
	if err := json.Unmarshal([]byte(serializedParams), params); err != nil {
		return "", stacktrace.Propagate(err, "An error occurred unmarshalling the params JSON")
	}

	createdServiceIdPorts := map[string]uint32{}
	for i := uint32(0); i < params.NumDatastores; i++ {
		serviceId, err := module.addDatastoreService(networkCtx)
		if err != nil {
			return "", stacktrace.Propagate(err, "An error occurred adding a datastore service")
		}
		createdServiceIdPorts[string(serviceId)] = datastorePortNumber
	}
	resultObj := ExecuteResult{
		CreatedServiceIdPorts: createdServiceIdPorts,
	}
	resultJsonBytes, err := json.Marshal(resultObj)
	if err != nil {
		return "", stacktrace.Propagate(err, "An error occurred serialzing the Lambda result object to JSON")
	}
	return string(resultJsonBytes), nil
}

// ====================================================================================================
//                                       Private helper functions
// ====================================================================================================
func (module *DatastoreArmyKurtosisModule) addDatastoreService(networkCtx *networks.NetworkContext) (services.ServiceID, error) {
	nextDatastoreServiceId := services.ServiceID(fmt.Sprintf("datastore-%v", module.numDatstoresAdded))

	datastoreContainerConfigSupplier := getDatastoreContainerConfigSupplier()

	if _, _, err := networkCtx.AddService(nextDatastoreServiceId, datastoreContainerConfigSupplier); err != nil {
		return "", stacktrace.Propagate(err, "An error occurred adding datastore service '%v'", nextDatastoreServiceId)
	}
	module.numDatstoresAdded = module.numDatstoresAdded + 1
	return nextDatastoreServiceId, nil
}


func getDatastoreContainerConfigSupplier() func(ipAddr string, sharedDirectory *services.SharedPath) (*services.ContainerConfig, error) {
	containerConfigSupplier  := func(ipAddr string, sharedDirectory *services.SharedPath) (*services.ContainerConfig, error) {
		containerConfig := services.NewContainerConfigBuilder(
			datastoreImage,
		).WithUsedPorts(
			map[string]bool{fmt.Sprintf("%v/%v", datastorePortNumber, datastoreProtocol): true},
		).Build()
		return containerConfig, nil
	}
	return containerConfigSupplier
}
