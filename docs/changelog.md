# TBD

# 0.2.7

# 0.2.6
### Features
* Added CircleCi workflow for running a scheduled pipeline every day to control successful module execution
* Added slack orb in the CircleCi config file to notify when the `check_module_execution` job fails

# 0.2.5
### Changes
* Migrate repo to use internal cli tool `kudet`, for updating release workflow
* Upgrade to module-api-lib 0.18.0
* Upgrade to core 1.57.0

# 0.2.4
### Changes
* Upgrade to module-api-lib 0.17.0
* Upgrade to core 1.55.2

# 0.2.3
### Changes
* Upgrade to module-api-lib 0.16.0 and core 1.54.1

# 0.2.2
### Changes
* Upgrade to module-api-lib 0.15.0

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
