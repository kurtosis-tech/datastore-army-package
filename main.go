package main

import (
	"fmt"
	"github.com/kurtosis-tech/kurtosis-client/golang/kurtosis_core_rpc_api_bindings"
	"github.com/kurtosis-tech/kurtosis-client/golang/lib/networks"
	"github.com/kurtosis-tech/kurtosis-lambda-api-lib/golang/kurtosis_lambda_docker_api"
	"github.com/kurtosis-tech/kurtosis-lambda-api-lib/golang/kurtosis_lambda_rpc_api_bindings"
	"github.com/kurtosis-tech/kurtosis-lambda-api-lib/golang/kurtosis_lambda_rpc_api_consts"
	"github.com/kurtosis-tech/minimal-grpc-server/server"
	"github.com/mieubrisse/datastore-army-module/lambda"
	"github.com/palantir/stacktrace"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"os"
	"time"
)

const (
	ApiContainerSocketEnvVar = "API_CONTAINER_SOCKET"
	CustomParamsJsonEnvVar = "CUSTOM_PARAMS_JSON"

	successExitCode = 0
	errorExitCode = 1

	stopGracePeriod = 10 * time.Second
)

func main() {
	if err := runMain(); err != nil {
		logrus.Error("An error occurred running the main function:")
		fmt.Fprintln(logrus.StandardLogger().Out, err)
		os.Exit(errorExitCode)
	}
	os.Exit(successExitCode)
}

func runMain() error {
	apiContainerSocketStr, found := os.LookupEnv(ApiContainerSocketEnvVar)
	if !found {
		return stacktrace.NewError("No API container socket environment variable '%v' defined", ApiContainerSocketEnvVar)
	}

	// TODO SECURITY: Use HTTPS to verify we're hitting the correct API container
	conn, err := grpc.Dial(apiContainerSocketStr, grpc.WithInsecure())
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred dialling the API container at '%v'", apiContainerSocketStr)
	}

	apiClient := kurtosis_core_rpc_api_bindings.NewApiContainerServiceClient(conn)
	networkCtx := networks.NewNetworkContext(
		apiClient,
		kurtosis_lambda_docker_api.ExecutionVolumeMountpoint,
	)

	lambdaService := lambda.NewLambdaService(networkCtx)
	lambdaServiceRegistrationFunc := func(grpcServer *grpc.Server) {
		kurtosis_lambda_rpc_api_bindings.RegisterLambdaServiceServer(grpcServer, lambdaService)
	}

	rpcServer := server.NewMinimalGRPCServer(
		kurtosis_lambda_rpc_api_consts.ListenPort,
		kurtosis_lambda_rpc_api_consts.ListenProtocol,
		stopGracePeriod,
		[]func(*grpc.Server){lambdaServiceRegistrationFunc},
	)
	if err := rpcServer.Run(); err != nil {
		return stacktrace.Propagate(err, "An error occurred running the gRPC server serving the Lambda service")
	}
	return nil
}

