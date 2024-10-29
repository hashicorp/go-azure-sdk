
## `github.com/hashicorp/go-azure-sdk/resource-manager/standbypool/2024-03-01/standbycontainergrouppoolruntimeviews` Documentation

The `standbycontainergrouppoolruntimeviews` SDK allows for interaction with Azure Resource Manager `standbypool` (API Version `2024-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/standbypool/2024-03-01/standbycontainergrouppoolruntimeviews"
```


### Client Initialization

```go
client := standbycontainergrouppoolruntimeviews.NewStandbyContainerGroupPoolRuntimeViewsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `StandbyContainerGroupPoolRuntimeViewsClient.Get`

```go
ctx := context.TODO()
id := standbycontainergrouppoolruntimeviews.NewRuntimeViewID("12345678-1234-9876-4563-123456789012", "example-resource-group", "standbyContainerGroupPoolName", "runtimeViewName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StandbyContainerGroupPoolRuntimeViewsClient.ListByStandbyPool`

```go
ctx := context.TODO()
id := standbycontainergrouppoolruntimeviews.NewStandbyContainerGroupPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "standbyContainerGroupPoolName")

// alternatively `client.ListByStandbyPool(ctx, id)` can be used to do batched pagination
items, err := client.ListByStandbyPoolComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
