
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/staticsites` Documentation

The `staticsites` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/staticsites"
```


### Client Initialization

```go
client := staticsites.NewStaticSitesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `StaticSitesClient.ApproveOrRejectPrivateEndpointConnection`

```go
ctx := context.TODO()
id := staticsites.NewStaticSitePrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "privateEndpointConnectionName")

payload := staticsites.RemotePrivateEndpointConnectionARMResource{
	// ...
}


if err := client.ApproveOrRejectPrivateEndpointConnectionThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `StaticSitesClient.DeletePrivateEndpointConnection`

```go
ctx := context.TODO()
id := staticsites.NewStaticSitePrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "privateEndpointConnectionName")

if err := client.DeletePrivateEndpointConnectionThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `StaticSitesClient.GetPrivateEndpointConnection`

```go
ctx := context.TODO()
id := staticsites.NewStaticSitePrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "privateEndpointConnectionName")

read, err := client.GetPrivateEndpointConnection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StaticSitesClient.GetPrivateEndpointConnectionList`

```go
ctx := context.TODO()
id := staticsites.NewStaticSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName")

// alternatively `client.GetPrivateEndpointConnectionList(ctx, id)` can be used to do batched pagination
items, err := client.GetPrivateEndpointConnectionListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
