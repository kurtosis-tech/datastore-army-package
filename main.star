datastore_package = import_module('/src/datastore-army-package.star')
helpers = import_module('/src/helpers.star')

PACKAGE_NAME_FOR_LOGGING = "datastore-army-package"


def run(plan, args):
    plan.print("Deploying package " + PACKAGE_NAME_FOR_LOGGING + " with args: \n{0}".format(args))
    args = helpers.apply_default_to_input_args(plan, args)

    if args.num_datastores == 0:
        fail("'num_datastores' is zero in package parameter. Nothing will be deployed.")

    output = datastore_package.add_multiple_datastore_services(plan, args.num_datastores, args.parallel)

    plan.print("Package " + PACKAGE_NAME_FOR_LOGGING + " successfully deployed")
    return output
