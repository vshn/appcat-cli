# appcat-cli

`appcat-cli` generates yaml files for custom kubernetes resources from command line arguments.

**Warning:** This tool is currently under heavy development. Stuff may change & break between releases!


## Goal & Purpose
`appcat-cli` is developed as a converter tool for [k8ify](https://github.com/vshn/k8ify).
It's purpose is to generate yaml files for custom kubernetes resources to be processed by `k8ify` into kubernetes manifests.

### Non-Goals

Out of scope(for now) of this project are:
- Checking scope and availability of custom resources on target cluster

## Mode of Operation

### Command Line Arguments

`appcatcli` expects input in the form of:<br>
`appcatcli ServiceName --kind ServiceKind [Options]`<br>

#### Defaults:
Global:
| Key  | Value  |
| ------ | -------|
|`spec.writeConnectionSecretToRef.Name`| ServiceName+refSlug |

Exoscale:
|ServiceKind |Key  | Value  |
| ------ | -------|-------|
|ExoscalePostgreSQL|`spec.parameters.service.majorVersion`| 14|
|ExoscalePostgreSQL|`spec.parameters.service.zone`| ch-dk-2|
|ExoscalePostgreSQL|`spec.parameters.size.plan`| hobbyist-2|
|ExoscalePostgreSQL|`spec.parameters.backup.timeOfDay`| 12:00:00|
|ExoscalePostgreSQL|`spec.parameters.maintenance.dayOfWeek`| sunday|
|ExoscalePostgreSQL|`spec.parameters.maintenance.timeOfDaye`| 00:00:00|
|ExoscaleRedis|`spec.parameters.maintenance.dayOfWeek`|sunday|
|ExoscaleRedis|`spec.parameters.maintenance.timeOfDay`|00:00:00|
|ExoscaleRedis|`spec.parameters.service.zone`|ch-dk-2|
|ExoscaleKafka|`spec.parameters.service.version`|3.4.0|
|ExoscaleKafka|`spec.parameters.service.zone`|ch-dk-2|
|ExoscaleKafka|`spec.parameters.size.plan`|startup-2|
|ExoscaleKafka|`spec.parameters.maintenance.dayOfWeek`|sunday|
|ExoscaleKafka|`spec.parameters.maintenance.timeOfDay`|00:00:00|
|ExoscaleMySQL|`spec.parameters.service.majorVersion`|8|
|ExoscaleMySQL|`spec.parameters.service.zone`|ch-dk-2|
|ExoscaleMySQL|`spec.parameters.size.plan`|hobbyist-2|
|ExoscaleMySQL|`spec.parameters.backup.timeOfDay`|12:00:00|
|ExoscaleMySQL|`spec.parameters.maintenance.dayOfWeek`|sunday|
|ExoscaleMySQL|`spec.parameters.maintenance.timeOfDay`|00:00:00|
|ExoscaleOpenSearch|`spec.parameters.service.majorVersion`|2|
|ExoscaleOpenSearch|`spec.parameters.service.zone`|ch-dk-2|
|ExoscaleOpenSearch|`spec.parameters.size.plan`|hobbyist-2|
|ExoscaleOpenSearch|`spec.parameters.backup.timeOfDay`|12:00:00|
|ExoscaleOpenSearch|`spec.parameters.maintenance.dayOfWeek`|sunday|
|ExoscaleOpenSearch|`spec.parameters.maintenance.timeOfDay`|00:00:00|

VSHN:
|ServiceKind |Key  | Value  |
| ------ | -------|-------|
|VSHNPostgreSQL|`spec.parameters.service.majorVersion`|14|
|VSHNPostgreSQL|`spec.parameters.size.cpu`|600m|
|VSHNPostgreSQL|`spec.parameters.size.disk`|80Gi|
|VSHNPostgreSQL|`spec.parameters.size.memory`|3500Mi|
|VSHNPostgreSQL|`spec.parameters.size.requests.cpu`|300m|
|VSHNPostgreSQL|`spec.parameters.size.requests.memory`|1000Mi|
|VSHNPostgreSQL|`spec.parameters.backup.schedule`||30 23 * * *|
|VSHNPostgreSQL|`spec.parameters.backup.retention`|12|
|VSHNPostgreSQL|`spec.parameters.scheduling.nodeSelector`|{"appuio.io/node-class": "plus"}|
|VSHNRedis|`spec.parameters.tls.authClients`|true|
|VSHNRedis|`spec.parameters.tls.enabled`|true|
|VSHNRedis|`spec.parameters.service.version`|7.0|
|VSHNRedis|`spec.parameters.service.redisSettings`||activedefrag yes|
|VSHNRedis|`spec.parameters.size.disk`|80Gi|
|VSHNRedis|`spec.parameters.size.cpuLimits`|1000m|
|VSHNRedis|`spec.parameters.size.cpuRequests`|500m|
|VSHNRedis|`spec.parameters.size.memoryRequests`|500Mi|
|VSHNRedis|`spec.parameters.size.memoryLimits`|1Gi|
|VSHNKeycloak|`spec.parameters.service.postgreSQLParameters.encryption.enabled`|true|

## Testing

For testing we use unit test as well as golden tests both can be run via:
```shell
go test .
```

## License

This project is licensed under the [BSD 3-Clause License](LICENSE)
