
## `github.com/hashicorp/go-azure-sdk/resource-manager/resources/2015-11-01/providers` Documentation

The `providers` SDK allows for interaction with Azure Resource Manager `resources` (API Version `2015-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/resources/2015-11-01/providers"
```


### Client Initialization

```go
client := providers.NewProvidersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ProvidersClient.Get`

```go
ctx := context.TODO()
id := providers.NewProviderID("12345678-1234-9876-4563-123456789012", "resourceProviderNamespace")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProvidersClient.List`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id, providers.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, providers.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ProvidersClient.Register`

```go
ctx := context.TODO()
id := providers.NewProviderID("12345678-1234-9876-4563-123456789012", "resourceProviderNamespace")

read, err := client.Register(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProvidersClient.Unregister`

```go
ctx := context.TODO()
id := providers.NewProviderID("12345678-1234-9876-4563-123456789012", "resourceProviderNamespace")

read, err := client.Unregister(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
