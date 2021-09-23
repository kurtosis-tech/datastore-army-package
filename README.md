Datastore Army Lambda
=====================
This repository contains [a Kurtosis Lambda](https://github.com/kurtosis-tech/kurtosis-lambda-api-lib/pull/11) that starts [example datastore services](https://github.com/kurtosis-tech/example-microservices/tree/develop/datastore). It is principally for demo purposes, and is published to Dockerhub [here](https://hub.docker.com/r/kurtosistech/datastore-army-lambda).

To use it in the [Kurtosis sandbox](https://docs.kurtosistech.com/sandbox.html):

```javascript
loadLambdaResult = await networkCtx.loadLambda("datastore-army-lambda", "kurtosistech/datastore-army-lambda", "{}")
lambdaCtx = loadLambdaResult.value
executeResult = await lambdaCtx.execute("{\"numDatastores\":2}") // Replace with the number of datastore services you want
executeResultObj = JSON.parse(executeResult.value)
console.log(executeResultObj)
```

Its args & result are JSON-serialized, corresponding to [these data structures](https://github.com/kurtosis-tech/datastore-army-lambda/blob/master/lambda/api.go).
