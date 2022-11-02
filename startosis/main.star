module_io = import_types("github.com/kurtosis-tech/datastore-army-module/types.proto")

MODULE_NAME_FOR_LOGGING = "datastore_army_module"

DATASTORE_IMAGE = "kurtosistech/example-datastore-server"
DATASTORE_PORT_ID = "grpc"
DATASTORE_PORT_NUMBER = 1323
DATASTORE_PORT_PROTOCOL = "TCP"

SERVICE_ID_PREFIX = "datastore-"


def add_datastore_service(unique_service_id):
    print("Adding service " + unique_service_id)

    service_config = struct(
        container_image_name = DATASTORE_IMAGE,
        used_ports = {
            DATASTORE_PORT_ID: struct(number = DATASTORE_PORT_NUMBER, protocol = DATASTORE_PORT_PROTOCOL)
        }
    )
    add_service(service_id = unique_service_id, service_config = service_config)
    return DATASTORE_PORT_ID


def convert_output(service_id_to_port_id_map):
    deployed_datastores = []
    for service_id in service_id_to_port_id_map:
        deployed_datastores.append(module_io.ServiceIdPortId(
            service_id=service_id,
            port_id=service_id_to_port_id_map[service_id]
        ))
    return module_io.ModuleOutput(
        deployed_datastores=deployed_datastores
    )


def main(input_args):
    print("Deploying module " + MODULE_NAME_FOR_LOGGING + " with args:")
    print(input_args)

    service_id_to_port_id = {}
    for index in range(input_args.num_datastores):
        service_id = SERVICE_ID_PREFIX + str(index)
        service_id_to_port_id[service_id] = add_datastore_service(service_id)

    print("Module " + MODULE_NAME_FOR_LOGGING + " deployed successfully.")
    output = convert_output(service_id_to_port_id)
    print(output) # TODO(gb): remove once we print it in the framework
    return output
