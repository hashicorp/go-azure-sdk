
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2023-12-01/provider` Documentation

The `provider` SDK allows for interaction with Azure Resource Manager `web` (API Version `2023-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2023-12-01/provider"
```


### Client Initialization

```go
client := provider.NewProviderClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ProviderClient.GetAvailableStacks`

```go
ctx := context.TODO()


// alternatively `client.GetAvailableStacks(ctx, provider.DefaultGetAvailableStacksOperationOptions())` can be used to do batched pagination
items, err := client.GetAvailableStacksComplete(ctx, provider.DefaultGetAvailableStacksOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ProviderClient.GetAvailableStacksOnPrem`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.GetAvailableStacksOnPrem(ctx, id, provider.DefaultGetAvailableStacksOnPremOperationOptions())` can be used to do batched pagination
items, err := client.GetAvailableStacksOnPremComplete(ctx, id, provider.DefaultGetAvailableStacksOnPremOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ProviderClient.GetFunctionAppStacks`

```go
ctx := context.TODO()


// alternatively `client.GetFunctionAppStacks(ctx, provider.DefaultGetFunctionAppStacksOperationOptions())` can be used to do batched pagination
items, err := client.GetFunctionAppStacksComplete(ctx, provider.DefaultGetFunctionAppStacksOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ProviderClient.GetFunctionAppStacksForLocation`

```go
ctx := context.TODO()
id := provider.NewLocationID("locationName")

// alternatively `client.GetFunctionAppStacksForLocation(ctx, id, provider.DefaultGetFunctionAppStacksForLocationOperationOptions())` can be used to do batched pagination
items, err := client.GetFunctionAppStacksForLocationComplete(ctx, id, provider.DefaultGetFunctionAppStacksForLocationOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ProviderClient.GetWebAppStacks`

```go
ctx := context.TODO()


// alternatively `client.GetWebAppStacks(ctx, provider.DefaultGetWebAppStacksOperationOptions())` can be used to do batched pagination
items, err := client.GetWebAppStacksComplete(ctx, provider.DefaultGetWebAppStacksOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ProviderClient.GetWebAppStacksForLocation`

```go
ctx := context.TODO()
id := provider.NewLocationID("locationName")

// alternatively `client.GetWebAppStacksForLocation(ctx, id, provider.DefaultGetWebAppStacksForLocationOperationOptions())` can be used to do batched pagination
items, err := client.GetWebAppStacksForLocationComplete(ctx, id, provider.DefaultGetWebAppStacksForLocationOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
