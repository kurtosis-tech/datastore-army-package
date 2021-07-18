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
	datastorePortSpec = "1323/tcp"
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

	createdServiceIdsSet := map[string]bool{}
	for i := uint32(0); i < paramsJson.NumDatastores; i++ {
		serviceId, err := l.addDatastoreService()
		if err != nil {
			return nil, stacktrace.Propagate(err, "An error occurred adding a datastore service")
		}
		createdServiceIdsSet[string(serviceId)] = true
	}
	resultObj := Result{CreatedServiceIdsSet: createdServiceIdsSet}
	resultJsonBytes, err := json.Marshal(resultObj)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred serialzing the Lambda result object to JSON")
	}

	response := &kurtosis_lambda_rpc_api_bindings.ExecuteResponse{ResponseJson: string(resultJsonBytes)}
	return response, nil
}

func (l *LambdaService) addDatastoreService() (services.ServiceID, error) {
	nextDatastoreServiceId := services.ServiceID(fmt.Sprintf("datastore-%v", l.numDatstoresAdded))

	containerCreationConfig := services.NewContainerCreationConfigBuilder(
		datastoreImage,
	).WithUsedPorts(map[string]bool{
		datastorePortSpec: true,
	}).Build()

	containerRunConfigSupplier := func(
			ipAddr string,
			generatedFileFilepaths map[string]string,
			staticFileFilepaths map[services.StaticFileID]string) (*services.ContainerRunConfig, error) {
		result := services.NewContainerRunConfigBuilder().Build()
		return result, nil
	}

	if _, _, err := l.networkCtx.AddService(nextDatastoreServiceId, containerCreationConfig, containerRunConfigSupplier); err != nil {
		return "", stacktrace.Propagate(err, "An error occurred adding datastore service '%v'", nextDatastoreServiceId)
	}
	l.numDatstoresAdded = l.numDatstoresAdded + 1
	return nextDatastoreServiceId, nil
}

