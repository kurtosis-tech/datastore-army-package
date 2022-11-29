NUM_DATASTORES_ARG_NAME = 'num_datastores'
DEFAULT_NUM_DATASTORES = 2

def apply_default_to_input_args(input_args):
    num_datastores = DEFAULT_NUM_DATASTORES
    if hasattr(input_args, NUM_DATASTORES_ARG_NAME):
        num_datastores = input_args.num_datastores
    else:
        print("'{0}' not set in package args. Default value '{1}' will be applied.".format(NUM_DATASTORES_ARG_NAME, DEFAULT_NUM_DATASTORES))
    return struct(num_datastores=num_datastores)
