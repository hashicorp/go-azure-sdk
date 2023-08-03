
## `github.com/hashicorp/go-azure-sdk/resource-manager/maps/2023-06-01/operations` Documentation

The `operations` SDK allows for interaction with the Azure Resource Manager Service `maps` (API Version `2023-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/maps/2023-06-01/operations"
```


### Client Initialization

```go
client := operations.NewOperationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OperationsClient.MapsListSubscriptionOperations`

```go
ctx := context.TODO()
id := operations.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.MapsListSubscriptionOperations(ctx, id)` can be used to do batched pagination
items, err := client.MapsListSubscriptionOperationsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
