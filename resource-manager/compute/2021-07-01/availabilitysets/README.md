
## `github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/availabilitysets` Documentation

The `availabilitysets` SDK allows for interaction with the Azure Resource Manager Service `compute` (API Version `2021-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/availabilitysets"
```


### Client Initialization

```go
client := availabilitysets.NewAvailabilitySetsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AvailabilitySetsClient.ResourceSkusList`

```go
ctx := context.TODO()
id := availabilitysets.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ResourceSkusList(ctx, id, availabilitysets.DefaultResourceSkusListOperationOptions())` can be used to do batched pagination
items, err := client.ResourceSkusListComplete(ctx, id, availabilitysets.DefaultResourceSkusListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
