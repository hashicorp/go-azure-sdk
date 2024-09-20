
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-01-01/replicationstorageclassificationmappings` Documentation

The `replicationstorageclassificationmappings` SDK allows for interaction with Azure Resource Manager `recoveryservicessiterecovery` (API Version `2023-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-01-01/replicationstorageclassificationmappings"
```


### Client Initialization

```go
client := replicationstorageclassificationmappings.NewReplicationStorageClassificationMappingsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReplicationStorageClassificationMappingsClient.Create`

```go
ctx := context.TODO()
id := replicationstorageclassificationmappings.NewReplicationStorageClassificationMappingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "fabricName", "storageClassificationName", "storageClassificationMappingName")

payload := replicationstorageclassificationmappings.StorageClassificationMappingInput{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ReplicationStorageClassificationMappingsClient.Delete`

```go
ctx := context.TODO()
id := replicationstorageclassificationmappings.NewReplicationStorageClassificationMappingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "fabricName", "storageClassificationName", "storageClassificationMappingName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ReplicationStorageClassificationMappingsClient.Get`

```go
ctx := context.TODO()
id := replicationstorageclassificationmappings.NewReplicationStorageClassificationMappingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "fabricName", "storageClassificationName", "storageClassificationMappingName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReplicationStorageClassificationMappingsClient.List`

```go
ctx := context.TODO()
id := replicationstorageclassificationmappings.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ReplicationStorageClassificationMappingsClient.ListByReplicationStorageClassifications`

```go
ctx := context.TODO()
id := replicationstorageclassificationmappings.NewReplicationStorageClassificationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "fabricName", "storageClassificationName")

// alternatively `client.ListByReplicationStorageClassifications(ctx, id)` can be used to do batched pagination
items, err := client.ListByReplicationStorageClassificationsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
