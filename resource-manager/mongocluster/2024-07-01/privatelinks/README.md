
## `github.com/hashicorp/go-azure-sdk/resource-manager/mongocluster/2024-07-01/privatelinks` Documentation

The `privatelinks` SDK allows for interaction with the Azure Resource Manager Service `mongocluster` (API Version `2024-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mongocluster/2024-07-01/privatelinks"
```


### Client Initialization

```go
client := privatelinks.NewPrivateLinksClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateLinksClient.ListByMongoCluster`

```go
ctx := context.TODO()
id := privatelinks.NewMongoClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "mongoClusterValue")

// alternatively `client.ListByMongoCluster(ctx, id)` can be used to do batched pagination
items, err := client.ListByMongoClusterComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
