
## `github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2024-02-01/slices` Documentation

The `slices` SDK allows for interaction with Azure Resource Manager `mobilenetwork` (API Version `2024-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2024-02-01/slices"
```


### Client Initialization

```go
client := slices.NewSlicesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SlicesClient.ListByMobileNetwork`

```go
ctx := context.TODO()
id := slices.NewMobileNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "mobileNetworkName")

// alternatively `client.ListByMobileNetwork(ctx, id)` can be used to do batched pagination
items, err := client.ListByMobileNetworkComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
