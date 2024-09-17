
## `github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/privatelinkhubprivatelinkresources` Documentation

The `privatelinkhubprivatelinkresources` SDK allows for interaction with Azure Resource Manager `synapse` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/privatelinkhubprivatelinkresources"
```


### Client Initialization

```go
client := privatelinkhubprivatelinkresources.NewPrivateLinkHubPrivateLinkResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateLinkHubPrivateLinkResourcesClient.Get`

```go
ctx := context.TODO()
id := privatelinkhubprivatelinkresources.NewPrivateLinkHubPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateLinkHubValue", "privateLinkResourceValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateLinkHubPrivateLinkResourcesClient.List`

```go
ctx := context.TODO()
id := privatelinkhubprivatelinkresources.NewPrivateLinkHubID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateLinkHubValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
