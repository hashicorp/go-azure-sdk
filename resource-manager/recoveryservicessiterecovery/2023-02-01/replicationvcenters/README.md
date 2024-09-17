
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-02-01/replicationvcenters` Documentation

The `replicationvcenters` SDK allows for interaction with Azure Resource Manager `recoveryservicessiterecovery` (API Version `2023-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-02-01/replicationvcenters"
```


### Client Initialization

```go
client := replicationvcenters.NewReplicationvCentersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReplicationvCentersClient.Create`

```go
ctx := context.TODO()
id := replicationvcenters.NewReplicationvCenterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationFabricValue", "replicationvCenterValue")

payload := replicationvcenters.AddVCenterRequest{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ReplicationvCentersClient.Delete`

```go
ctx := context.TODO()
id := replicationvcenters.NewReplicationvCenterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationFabricValue", "replicationvCenterValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ReplicationvCentersClient.Get`

```go
ctx := context.TODO()
id := replicationvcenters.NewReplicationvCenterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationFabricValue", "replicationvCenterValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReplicationvCentersClient.List`

```go
ctx := context.TODO()
id := replicationvcenters.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ReplicationvCentersClient.ListByReplicationFabrics`

```go
ctx := context.TODO()
id := replicationvcenters.NewReplicationFabricID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationFabricValue")

// alternatively `client.ListByReplicationFabrics(ctx, id)` can be used to do batched pagination
items, err := client.ListByReplicationFabricsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ReplicationvCentersClient.Update`

```go
ctx := context.TODO()
id := replicationvcenters.NewReplicationvCenterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationFabricValue", "replicationvCenterValue")

payload := replicationvcenters.UpdateVCenterRequest{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
