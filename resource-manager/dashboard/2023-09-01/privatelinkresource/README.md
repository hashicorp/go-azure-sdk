
## `github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2023-09-01/privatelinkresource` Documentation

The `privatelinkresource` SDK allows for interaction with Azure Resource Manager `dashboard` (API Version `2023-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2023-09-01/privatelinkresource"
```


### Client Initialization

```go
client := privatelinkresource.NewPrivateLinkResourceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateLinkResourceClient.Get`

```go
ctx := context.TODO()
id := privatelinkresource.NewPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "grafanaName", "privateLinkResourceName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateLinkResourceClient.List`

```go
ctx := context.TODO()
id := privatelinkresource.NewGrafanaID("12345678-1234-9876-4563-123456789012", "example-resource-group", "grafanaName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
