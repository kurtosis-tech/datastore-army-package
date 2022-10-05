#vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv PARAMETERS vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv#

num_datastores = 3

#^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ END OF PARAMETERS ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^#


#vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv STATIC CONST vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv#

module_name_for_logging = "datastore_army_module"

datastore_image = "kurtosistech/example-datastore-server"
datastore_port_id = "grpc"
datastore_port_number = 1323

service_id_prefix = "datastore-"

#^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ END OF STATIC CONST ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^#


#vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv SCRIPT vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv#

def add_datastore_service(unique_service_id):
    print("Adding service " + unique_service_id)

    service_config = struct(
        container_image_name = datastore_image,
        used_ports = {
            datastore_port_id: struct(number = datastore_port_number, protocol = "TCP")
        }
    )
    add_service(service_id = unique_service_id, service_config = service_config)

def add_multiple_datastore_services(num_datastores):
    for i in range(num_datastores):
        service_id = service_id_prefix + str(i)
        add_datastore_service(service_id)

print("Deploying module " + module_name_for_logging)

add_multiple_datastore_services(num_datastores)

print("Module " + module_name_for_logging + " deployed successfully.")
