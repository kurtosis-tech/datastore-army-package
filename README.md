Datastore Army Module
=====================
This repository contains [an executable Kurtosis module](https://docs.kurtosistech.com/modules.html) that starts [example datastore services](https://github.com/kurtosis-tech/example-microservices/tree/develop/datastore). It is principally for demo purposes, and is published to Dockerhub [here](https://hub.docker.com/repository/docker/kurtosistech/datastore-army-module).

To run it, use:

```
kurtosis module exec 'kurtosistech/datastore-army-module' --execute-params '{"numDatastores":2}' # Replace with the number of datastore services you want
```

Its args & result are JSON-serialized, corresponding to [these data structures](https://github.com/kurtosis-tech/datastore-army-module/blob/master/kurtosis-module/impl/api.go).
