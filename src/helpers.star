module_io = import_types("github.com/kurtosis-tech/datastore-army-module/types.proto")


def convert_output(service_id_to_port_id_map):
    deployed_datastores = []
    for service_id in service_id_to_port_id_map:
        deployed_datastores.append(module_io.ServiceIdPortId(
            service_id=service_id,
            port_id=service_id_to_port_id_map[service_id]
        ))
    return module_io.ModuleOutput(
        created_service_ids_to_port_ids=deployed_datastores
    )
