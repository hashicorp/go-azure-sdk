
## `github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-07-01/privatelinkresource` Documentation

The `privatelinkresource` SDK allows for interaction with the Azure Resource Manager Service `purview` (API Version `2021-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-07-01/privatelinkresource"
```


### Client Initialization

```go
client := privatelinkresource.NewPrivateLinkResourceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateLinkResourceClient.GetByGroupId`

```go
ctx := context.TODO()
id := privatelinkresource.NewPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "groupIdValue")

read, err := client.GetByGroupId(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateLinkResourceClient.ListByAccount`

```go
ctx := context.TODO()
id := privatelinkresource.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

// alternatively `client.ListByAccount(ctx, id)` can be used to do batched pagination
items, err := client.ListByAccountComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
