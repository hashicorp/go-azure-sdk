
## `github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/databoxedges` Documentation

The `databoxedges` SDK allows for interaction with Azure Resource Manager `databoxedge` (API Version `2023-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/databoxedges"
```


### Client Initialization

```go
client := databoxedges.NewDataboxedgesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DataboxedgesClient.AvailableSkusList`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.AvailableSkusList(ctx, id)` can be used to do batched pagination
items, err := client.AvailableSkusListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
