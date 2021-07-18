package lambda

import (
	"fmt"
	"github.com/kurtosis-tech/kurtosis-client/golang/kurtosis_core_rpc_api_bindings"
	"github.com/kurtosis-tech/kurtosis-client/golang/lib/networks"
	"github.com/kurtosis-tech/kurtosis-client/golang/lib/services"
	"github.com/kurtosis-tech/kurtosis-lambda-api-lib/golang/kurtosis_lambda_docker_api"
	"github.com/kurtosis-tech/kurtosis-lambda-api-lib/golang/kurtosis_lambda_rpc_api_bindings"
	"github.com/palantir/stacktrace"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"os"
)

const (
	ApiContainerSocketEnvVar = "API_CONTAINER_SOCKET"
	CustomParamsJsonEnvVar = "CUSTOM_PARAMS_JSON"

	successExitCode = 0
	errorExitCode = 1
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

	conn, err := grpc.Dial(apiContainerSocketStr)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred dialling the API container at '%v'", apiContainerSocketStr)
	}

	apiClient := kurtosis_core_rpc_api_bindings.NewApiContainerServiceClient(conn)
	networkCtx := networks.NewNetworkContext(
		apiClient,
		map[services.FilesArtifactID]string{},
		kurtosis_lambda_docker_api.ExecutionVolumeMountpoint,
	)

	fmt.Println("Hello world")
}