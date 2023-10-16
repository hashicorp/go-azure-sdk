
## `github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/services` Documentation

The `services` SDK allows for interaction with the Azure Resource Manager Service `mobilenetwork` (API Version `2023-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/services"
```


### Client Initialization

```go
client := services.NewServicesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ServicesClient.ListByMobileNetwork`

```go
ctx := context.TODO()
id := services.NewMobileNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "mobileNetworkValue")

// alternatively `client.ListByMobileNetwork(ctx, id)` can be used to do batched pagination
items, err := client.ListByMobileNetworkComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
