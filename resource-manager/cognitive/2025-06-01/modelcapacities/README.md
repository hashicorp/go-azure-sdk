
## `github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-06-01/modelcapacities` Documentation

The `modelcapacities` SDK allows for interaction with Azure Resource Manager `cognitive` (API Version `2025-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-06-01/modelcapacities"
```


### Client Initialization

```go
client := modelcapacities.NewModelCapacitiesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ModelCapacitiesClient.List`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id, modelcapacities.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, modelcapacities.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ModelCapacitiesClient.LocationBasedModelCapacitiesList`

```go
ctx := context.TODO()
id := modelcapacities.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

// alternatively `client.LocationBasedModelCapacitiesList(ctx, id, modelcapacities.DefaultLocationBasedModelCapacitiesListOperationOptions())` can be used to do batched pagination
items, err := client.LocationBasedModelCapacitiesListComplete(ctx, id, modelcapacities.DefaultLocationBasedModelCapacitiesListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
