
## `github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2020-12-01/availableskus` Documentation

The `availableskus` SDK allows for interaction with the Azure Resource Manager Service `databoxedge` (API Version `2020-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2020-12-01/availableskus"
```


### Client Initialization

```go
client := availableskus.NewAvailableSkusClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AvailableSkusClient.List`

```go
ctx := context.TODO()
id := availableskus.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
