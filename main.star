load('github.com/kurtosis-tech/datastore-army-module/src/datastore-army-module.star', add_multiple_datastore_services="add_multiple_datastore_services")

MODULE_NAME_FOR_LOGGING = "datastore_army_module"


def main(input_args):
    print("Deploying module " + MODULE_NAME_FOR_LOGGING + " with args:")
    print(input_args)

    # if not proto.has(input_args, "num_datastores"):
    #     print("'num_datastores' is not set in module parameter. Nothing will be deployed.")

    if input_args.num_datastores == 0:
        print("'num_datastores' is zero in module parameter. Nothing will be deployed.")

    output = add_multiple_datastore_services(input_args.num_datastores)

    print("Module " + MODULE_NAME_FOR_LOGGING + " deployed successfully.")
    print(output) # TODO(gb): remove once we print it in the framework
    return output
