#vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv PARAMETERS vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv#

NUM_DATASTORES = 3

#^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ END OF PARAMETERS ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^#


#vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv STATIC CONST vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv#

MODULE_NAME_FOR_LOGGING = "datastore_army_module"

DATASTORE_IMAGE = "kurtosistech/example-datastore-server"
DATASTORE_PORT_ID = "grpc"
DATASTORE_PORT_NUMBER = 1323
DATASTORE_PORT_PROTOCOL = "TCP"

SERVICE_ID_PREFIX = "datastore-"

#^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ END OF STATIC CONST ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^#


#vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv SCRIPT vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv#

def add_datastore_service(unique_service_id):
    print("Adding service " + unique_service_id)

    service_config = struct(
        container_image_name = DATASTORE_IMAGE,
        used_ports = {
            DATASTORE_PORT_ID: struct(number = DATASTORE_PORT_NUMBER, protocol = DATASTORE_PORT_PROTOCOL)
        }
    )
    add_service(service_id = unique_service_id, service_config = service_config)

def add_multiple_datastore_services(num_services):
    for index in range(num_services):
        service_id = SERVICE_ID_PREFIX + str(index)
        add_datastore_service(service_id)

print("Deploying module " + MODULE_NAME_FOR_LOGGING)

add_multiple_datastore_services(NUM_DATASTORES)

print("Module " + MODULE_NAME_FOR_LOGGING + " deployed successfully.")
