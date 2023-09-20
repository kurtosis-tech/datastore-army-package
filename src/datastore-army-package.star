helpers = import_module('./helpers.star')

DATASTORE_IMAGE = "kurtosistech/example-datastore-server"
DATASTORE_PORT_ID = "grpc"
DATASTORE_PORT_NUMBER = 1323
DATASTORE_TRANSPORT_PROTOCOL = "TCP"
SERVICE_ID_PREFIX = "datastore-"


def add_multiple_datastore_services(plan, num_datastores, parallel):
    if parallel:
        return add_multiple_datastore_services_parallel(plan, num_datastores)
    else:
        return add_multiple_datastore_services_sequential(plan, num_datastores)

def add_multiple_datastore_services_sequential(plan, num_datastores):
    plan.print("Adding {0} datastore services one by one".format(num_datastores))
    service_id_to_service_obj = {}
    for index in range(num_datastores):
        service_id = get_service_name(index)
        service_id_to_service_obj[service_id] = plan.add_service(service_id, config = get_service_config())
    return service_id_to_service_obj

def add_multiple_datastore_services_parallel(plan, num_datastores):
    all_service_configs = {}
    for index in range(num_datastores):
        service_id = get_service_name(index)
        all_service_configs[service_id] = get_service_config()

    plan.print("Adding {0} datastore services all at once".format(len(all_service_configs)))
    service_id_to_service_obj = plan.add_services(all_service_configs)
    return service_id_to_service_obj

def get_service_name(service_idx):
    return SERVICE_ID_PREFIX + str(service_idx)

def get_service_config():
    return ServiceConfig(
        image = DATASTORE_IMAGE,
        ports = {
            DATASTORE_PORT_ID: PortSpec(number = DATASTORE_PORT_NUMBER, transport_protocol = DATASTORE_TRANSPORT_PROTOCOL)
        }
    )
