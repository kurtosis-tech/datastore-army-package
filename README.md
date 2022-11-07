Datastore Army Module
=====================
This repository contains an executable Kurtosis module that starts [example datastore services](https://github.com/kurtosis-tech/example-microservices/tree/develop/datastore). It is principally for demo purposes.

To run it, use:

```
kurtosis startosis exec github.com/kurtosis-tech/datastore-army-module --args '{"numDatastores":2}' # Replace with the number of datastore services you want
```

Its args is JSON-serialized, corresponding to the `ModuleInput` datastructure defined in the `types.proto` file at the [root of the module](https://github.com/kurtosis-tech/datastore-army-module).
The `ModuleOutput` is the object returned by the module.
