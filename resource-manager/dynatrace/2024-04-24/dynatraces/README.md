
## `github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2024-04-24/dynatraces` Documentation

The `dynatraces` SDK allows for interaction with Azure Resource Manager `dynatrace` (API Version `2024-04-24`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2024-04-24/dynatraces"
```


### Client Initialization

```go
client := dynatraces.NewDynatracesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DynatracesClient.CreationSupportedGet`

```go
ctx := context.TODO()
id := dynatraces.NewSubscriptionStatusID("12345678-1234-9876-4563-123456789012", "dynatraceEnvironmentId")

// alternatively `client.CreationSupportedGet(ctx, id)` can be used to do batched pagination
items, err := client.CreationSupportedGetComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DynatracesClient.CreationSupportedList`

```go
ctx := context.TODO()
id := dynatraces.NewSubscriptionStatusID("12345678-1234-9876-4563-123456789012", "dynatraceEnvironmentId")

// alternatively `client.CreationSupportedList(ctx, id)` can be used to do batched pagination
items, err := client.CreationSupportedListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DynatracesClient.MonitorsGetAllConnectedResourcesCount`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

payload := dynatraces.MarketplaceSubscriptionIdRequest{
	// ...
}


read, err := client.MonitorsGetAllConnectedResourcesCount(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DynatracesClient.MonitorsGetMarketplaceSaaSResourceDetails`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

payload := dynatraces.MarketplaceSaaSResourceDetailsRequest{
	// ...
}


read, err := client.MonitorsGetMarketplaceSaaSResourceDetails(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
