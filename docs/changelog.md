# TBD

# 0.2.1
### Changes
* Upgrade to module-api-lib 0.14.1

# 0.2.0
### Changes
* Upgrade to module-api-lib 0.12.2 which supports the latest Kurt Core

### Breaking Changes
* The returned object now contains a mapping of `datastore_service_id` -> `datastore_port_id`, so the user can retrieve or public or private ports as they please
    * Users should swap the old `createdServiceIdPorts` property -> `createdServiceIdsToPortIds`

# 0.1.5
### Changes
* Replaced `kurtosistech/example-microservices_datastore` with the newest `kurtosistech/example-datastore-server` datastore image which implements GRPC

# 0.1.4
### Fixes
* Correct README link

# 0.1.3
### Fixes
* Fixed bug that occurred with calling execute multiple times, where the IDs wouldn't be updated

# 0.1.2
### Fixes
* Add `kurtosistech` org name to image so that CI can publish it to Dockerhub

# 0.1.1
### Features
* Adding CI

# 0.1.0
* Initial tagged commit
