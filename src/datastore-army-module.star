load('github.com/kurtosis-tech/datastore-army-module/src/helpers.star', convert_output="convert_output")


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


def add_multiple_datastore_services(num_datastores):
    service_id_to_port_id = {}
    for index in range(num_datastores):
        service_id = SERVICE_ID_PREFIX + str(index)
        service_id_to_port_id[service_id] = add_datastore_service(service_id)

    return convert_output(service_id_to_port_id)
