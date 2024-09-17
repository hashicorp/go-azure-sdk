
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2022-10-01/replicationalertsettings` Documentation

The `replicationalertsettings` SDK allows for interaction with Azure Resource Manager `recoveryservicessiterecovery` (API Version `2022-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2022-10-01/replicationalertsettings"
```


### Client Initialization

```go
client := replicationalertsettings.NewReplicationAlertSettingsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReplicationAlertSettingsClient.Create`

```go
ctx := context.TODO()
id := replicationalertsettings.NewReplicationAlertSettingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationAlertSettingValue")

payload := replicationalertsettings.ConfigureAlertRequest{
	// ...
}


read, err := client.Create(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReplicationAlertSettingsClient.Get`

```go
ctx := context.TODO()
id := replicationalertsettings.NewReplicationAlertSettingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationAlertSettingValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReplicationAlertSettingsClient.List`

```go
ctx := context.TODO()
id := replicationalertsettings.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
