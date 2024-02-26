
## `github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-07-07/privatelinkresources` Documentation

The `privatelinkresources` SDK allows for interaction with the Azure Resource Manager Service `migrate` (API Version `2020-07-07`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-07-07/privatelinkresources"
```


### Client Initialization

```go
client := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateLinkResourcesClient.GetPrivateLinkResource`

```go
ctx := context.TODO()
id := privatelinkresources.NewPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "masterSiteValue", "privateLinkResourceValue")

read, err := client.GetPrivateLinkResource(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateLinkResourcesClient.GetPrivateLinkResources`

```go
ctx := context.TODO()
id := privatelinkresources.NewMasterSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "masterSiteValue")

// alternatively `client.GetPrivateLinkResources(ctx, id)` can be used to do batched pagination
items, err := client.GetPrivateLinkResourcesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
