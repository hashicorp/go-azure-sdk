
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2024-01-01/replicationstorageclassifications` Documentation

The `replicationstorageclassifications` SDK allows for interaction with Azure Resource Manager `recoveryservicessiterecovery` (API Version `2024-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2024-01-01/replicationstorageclassifications"
```


### Client Initialization

```go
client := replicationstorageclassifications.NewReplicationStorageClassificationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReplicationStorageClassificationsClient.Get`

```go
ctx := context.TODO()
id := replicationstorageclassifications.NewReplicationStorageClassificationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "replicationFabricName", "replicationStorageClassificationName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReplicationStorageClassificationsClient.List`

```go
ctx := context.TODO()
id := replicationstorageclassifications.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ReplicationStorageClassificationsClient.ListByReplicationFabrics`

```go
ctx := context.TODO()
id := replicationstorageclassifications.NewReplicationFabricID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "replicationFabricName")

// alternatively `client.ListByReplicationFabrics(ctx, id)` can be used to do batched pagination
items, err := client.ListByReplicationFabricsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
