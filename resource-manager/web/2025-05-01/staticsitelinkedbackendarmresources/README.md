
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/staticsitelinkedbackendarmresources` Documentation

The `staticsitelinkedbackendarmresources` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/staticsitelinkedbackendarmresources"
```


### Client Initialization

```go
client := staticsitelinkedbackendarmresources.NewStaticSiteLinkedBackendARMResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `StaticSiteLinkedBackendARMResourcesClient.StaticSitesGetLinkedBackend`

```go
ctx := context.TODO()
id := staticsitelinkedbackendarmresources.NewLinkedBackendID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "linkedBackendName")

read, err := client.StaticSitesGetLinkedBackend(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StaticSiteLinkedBackendARMResourcesClient.StaticSitesGetLinkedBackends`

```go
ctx := context.TODO()
id := staticsitelinkedbackendarmresources.NewStaticSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName")

// alternatively `client.StaticSitesGetLinkedBackends(ctx, id)` can be used to do batched pagination
items, err := client.StaticSitesGetLinkedBackendsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `StaticSiteLinkedBackendARMResourcesClient.StaticSitesLinkBackend`

```go
ctx := context.TODO()
id := staticsitelinkedbackendarmresources.NewLinkedBackendID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "linkedBackendName")

payload := staticsitelinkedbackendarmresources.StaticSiteLinkedBackendARMResource{
	// ...
}


if err := client.StaticSitesLinkBackendThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `StaticSiteLinkedBackendARMResourcesClient.StaticSitesUnlinkBackend`

```go
ctx := context.TODO()
id := staticsitelinkedbackendarmresources.NewLinkedBackendID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "linkedBackendName")

read, err := client.StaticSitesUnlinkBackend(ctx, id, staticsitelinkedbackendarmresources.DefaultStaticSitesUnlinkBackendOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StaticSiteLinkedBackendARMResourcesClient.StaticSitesValidateBackend`

```go
ctx := context.TODO()
id := staticsitelinkedbackendarmresources.NewLinkedBackendID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "linkedBackendName")

payload := staticsitelinkedbackendarmresources.StaticSiteLinkedBackendARMResource{
	// ...
}


if err := client.StaticSitesValidateBackendThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
