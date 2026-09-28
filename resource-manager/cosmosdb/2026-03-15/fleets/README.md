
## `github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2026-03-15/fleets` Documentation

The `fleets` SDK allows for interaction with Azure Resource Manager `cosmosdb` (API Version `2026-03-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2026-03-15/fleets"
```


### Client Initialization

```go
client := fleets.NewFleetsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `FleetsClient.FleetListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.FleetListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.FleetListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
