
## `github.com/hashicorp/go-azure-sdk/resource-manager/mongocluster/2024-07-01/replicas` Documentation

The `replicas` SDK allows for interaction with the Azure Resource Manager Service `mongocluster` (API Version `2024-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mongocluster/2024-07-01/replicas"
```


### Client Initialization

```go
client := replicas.NewReplicasClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReplicasClient.ListByParent`

```go
ctx := context.TODO()
id := replicas.NewMongoClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "mongoClusterValue")

// alternatively `client.ListByParent(ctx, id)` can be used to do batched pagination
items, err := client.ListByParentComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
