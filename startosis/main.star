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
        container_image_name=DATASTORE_IMAGE,
        used_ports={
            DATASTORE_PORT_ID: struct(number=DATASTORE_PORT_NUMBER, protocol=DATASTORE_PORT_PROTOCOL)
        }
    )
    service_added = add_service(service_id=unique_service_id, service_config=service_config)
    print("Service successfully added: " + str(service_added))
    return DATASTORE_PORT_ID


def add_multiple_datastore_services(num_services):
    service_id_to_port_id = {}
    for index in range(num_services):
        service_id = SERVICE_ID_PREFIX + str(index)
        service_id_to_port_id[service_id] = add_datastore_service(service_id)
    return service_id_to_port_id


def convert_output(service_id_to_port_id_map):
    deployed_datastores = []
    for service_id in service_id_to_port_id_map:
        deployed_datastores.append(module_io.DatastoreDescription({
            "service_id": service_id,
            "port_id": service_id_to_port_id_map[service_id],
        }))
    module_output = module_io.ModuleOutput({
        "deployed_datastores": deployed_datastores
    })
    return module_output


def main(input_args):
    print("Deploying " + MODULE_NAME_FOR_LOGGING + ". Number of stores: " + str(input_args.num_datastores))
    service_id_to_port_id_map = add_multiple_datastore_services(input_args.num_datastores)
    module_output = convert_output(service_id_to_port_id_map)
    print("Module " + MODULE_NAME_FOR_LOGGING + " deployed successfully. Output is: " + str(module_output))
    return module_output
