
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-04-01/replicationevents` Documentation

The `replicationevents` SDK allows for interaction with the Azure Resource Manager Service `recoveryservicessiterecovery` (API Version `2023-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-04-01/replicationevents"
```


### Client Initialization

```go
client := replicationevents.NewReplicationEventsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReplicationEventsClient.Get`

```go
ctx := context.TODO()
id := replicationevents.NewReplicationEventID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationEventValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReplicationEventsClient.List`

```go
ctx := context.TODO()
id := replicationevents.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue")

// alternatively `client.List(ctx, id, replicationevents.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, replicationevents.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
