datastore_module = import_module('github.com/kurtosis-tech/datastore-army-module/src/datastore-army-module.star')
helpers = import_module('github.com/kurtosis-tech/datastore-army-module/src/helpers.star')

MODULE_NAME_FOR_LOGGING = "datastore_army_module"


def run(input_args):
    print("Deploying module " + MODULE_NAME_FOR_LOGGING + " with args:")
    print(input_args)
    helpers.apply_default_to_input_args(input_args)

    if input_args.num_datastores == 0:
        fail("'num_datastores' is zero in module parameter. Nothing will be deployed.")

    output = datastore_module.add_multiple_datastore_services(input_args.num_datastores)

    print("Module " + MODULE_NAME_FOR_LOGGING + " deployed successfully.")
    return output
