package lambda

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kurtosis-tech/kurtosis-client/golang/lib/networks"
	"github.com/kurtosis-tech/kurtosis-client/golang/lib/services"
	"github.com/kurtosis-tech/kurtosis-lambda-api-lib/golang/kurtosis_lambda_rpc_api_bindings"
	"github.com/palantir/stacktrace"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	datastoreImage = "kurtosistech/example-microservices_datastore"
	datastorePortNumber uint32 = 1323
	datastoreProtocol = "tcp"
)

type LambdaService struct {
	kurtosis_lambda_rpc_api_bindings.UnimplementedLambdaServiceServer

	networkCtx *networks.NetworkContext

	numDatstoresAdded int
}

func NewLambdaService(networkCtx *networks.NetworkContext) *LambdaService {
	return &LambdaService{
		networkCtx:                       networkCtx,
		numDatstoresAdded:                0,
	}
}

func (l *LambdaService) IsAvailable(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (l *LambdaService) Execute(ctx context.Context, args *kurtosis_lambda_rpc_api_bindings.ExecuteArgs) (*kurtosis_lambda_rpc_api_bindings.ExecuteResponse, error) {
	paramsJsonStr := args.ParamsJson
	paramsJson := new(Params)
	if err := json.Unmarshal([]byte(paramsJsonStr), paramsJson); err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred unmarshalling the params JSON")
	}

	createdServiceIdPorts := map[string]uint32{}
	for i := uint32(0); i < paramsJson.NumDatastores; i++ {
		serviceId, err := l.addDatastoreService()
		if err != nil {
			return nil, stacktrace.Propagate(err, "An error occurred adding a datastore service")
		}
		createdServiceIdPorts[string(serviceId)] = datastorePortNumber
	}
	resultObj := Result{
		CreatedServiceIdPorts: createdServiceIdPorts,
	}
	resultJsonBytes, err := json.Marshal(resultObj)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred serialzing the Lambda result object to JSON")
	}

	response := &kurtosis_lambda_rpc_api_bindings.ExecuteResponse{ResponseJson: string(resultJsonBytes)}
	return response, nil
}
// ====================================================================================================
//                                       Private helper functions
// ====================================================================================================
func (l *LambdaService) addDatastoreService() (services.ServiceID, error) {
	nextDatastoreServiceId := services.ServiceID(fmt.Sprintf("datastore-%v", l.numDatstoresAdded))

	datastoreContainerConfigSupplier := getDatastoreContainerConfigSupplier()

	if _, _, err := l.networkCtx.AddService(nextDatastoreServiceId, datastoreContainerConfigSupplier); err != nil {
		return "", stacktrace.Propagate(err, "An error occurred adding datastore service '%v'", nextDatastoreServiceId)
	}
	l.numDatstoresAdded = l.numDatstoresAdded + 1
	return nextDatastoreServiceId, nil
}


func getDatastoreContainerConfigSupplier() func(ipAddr string, sharedDirectory *services.SharedDirectory) (*services.ContainerConfig, error) {
	containerConfigSupplier  := func(ipAddr string, sharedDirectory *services.SharedDirectory) (*services.ContainerConfig, error) {
		containerConfig := services.NewContainerConfigBuilder(
			datastoreImage,
		).WithUsedPorts(
			map[string]bool{fmt.Sprintf("%v/%v", datastorePortNumber, datastoreProtocol): true},
		).Build()
		return containerConfig, nil
	}
	return containerConfigSupplier
}
