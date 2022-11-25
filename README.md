Datastore Army Package
=====================
This repository contains an executable Kurtosis package that starts [example datastore services](https://github.com/kurtosis-tech/example-microservices/tree/develop/datastore). It is principally for demo purposes.

To run it, use:

```
kurtosis run github.com/kurtosis-tech/datastore-army-module --args '{"num_datastores":2}'
```

Its args is JSON-serialized and should have the following structure:
```
{
	"num_datastores": 2 # Replace with the number of datastore services wanted
}
```

The output object returned by the package will be a mapping of `service_id` -> `port_id`.
For example:
```
{
	"datastore_1": "grpc",
	"datastore_2": "grpc",
}
```
