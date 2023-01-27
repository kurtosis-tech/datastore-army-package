PARALLEL_ARG_NAME = 'parallel'
NUM_DATASTORES_ARG_NAME = 'num_datastores'
DEFAULT_NUM_DATASTORES = 2

def apply_default_to_input_args(plan, input_args):
    num_datastores = DEFAULT_NUM_DATASTORES
    if hasattr(input_args, NUM_DATASTORES_ARG_NAME):
        num_datastores = input_args.num_datastores
    else:
        plan.print("'{0}' not set in package args. Default value '{1}' will be applied.".format(NUM_DATASTORES_ARG_NAME, DEFAULT_NUM_DATASTORES))

    parallel = False
    if hasattr(input_args, PARALLEL_ARG_NAME):
        parallel = input_args.parallel
    else:
        plan.print("'{0}' not set in package args. Services will be added sequentially.".format(PARALLEL_ARG_NAME))

    return struct(num_datastores=num_datastores, parallel=parallel)
