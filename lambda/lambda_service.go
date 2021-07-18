package lambda

import (
	"context"
	"encoding/json"
	"github.com/kurtosis-tech/kurtosis-lambda-api-lib/golang/kurtosis_lambda_rpc_api_bindings"
	"github.com/palantir/stacktrace"
	"google.golang.org/protobuf/types/known/emptypb"
)

type LambdaService struct {
	kurtosis_lambda_rpc_api_bindings.UnimplementedLambdaServiceServer
}

func (l LambdaService) IsAvailable(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (l LambdaService) Execute(ctx context.Context, args *kurtosis_lambda_rpc_api_bindings.ExecuteArgs) (*kurtosis_lambda_rpc_api_bindings.ExecuteResponse, error) {
	paramsJsonStr := args.ParamsJson
	paramsJson := new(Params)
	if err := json.Unmarshal([]byte(paramsJsonStr), paramsJson); err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred unmarshalling the params JSON")
	}

}