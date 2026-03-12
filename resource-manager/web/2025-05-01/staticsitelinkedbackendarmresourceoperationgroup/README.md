
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/staticsitelinkedbackendarmresourceoperationgroup` Documentation

The `staticsitelinkedbackendarmresourceoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/staticsitelinkedbackendarmresourceoperationgroup"
```


### Client Initialization

```go
client := staticsitelinkedbackendarmresourceoperationgroup.NewStaticSiteLinkedBackendARMResourceOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `StaticSiteLinkedBackendARMResourceOperationGroupClient.StaticSitesGetLinkedBackendForBuild`

```go
ctx := context.TODO()
id := staticsitelinkedbackendarmresourceoperationgroup.NewBuildLinkedBackendID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "buildName", "linkedBackendName")

read, err := client.StaticSitesGetLinkedBackendForBuild(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StaticSiteLinkedBackendARMResourceOperationGroupClient.StaticSitesGetLinkedBackendsForBuild`

```go
ctx := context.TODO()
id := staticsitelinkedbackendarmresourceoperationgroup.NewBuildID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "buildName")

// alternatively `client.StaticSitesGetLinkedBackendsForBuild(ctx, id)` can be used to do batched pagination
items, err := client.StaticSitesGetLinkedBackendsForBuildComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `StaticSiteLinkedBackendARMResourceOperationGroupClient.StaticSitesLinkBackendToBuild`

```go
ctx := context.TODO()
id := staticsitelinkedbackendarmresourceoperationgroup.NewBuildLinkedBackendID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "buildName", "linkedBackendName")

payload := staticsitelinkedbackendarmresourceoperationgroup.StaticSiteLinkedBackendARMResource{
	// ...
}


if err := client.StaticSitesLinkBackendToBuildThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `StaticSiteLinkedBackendARMResourceOperationGroupClient.StaticSitesUnlinkBackendFromBuild`

```go
ctx := context.TODO()
id := staticsitelinkedbackendarmresourceoperationgroup.NewBuildLinkedBackendID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "buildName", "linkedBackendName")

read, err := client.StaticSitesUnlinkBackendFromBuild(ctx, id, staticsitelinkedbackendarmresourceoperationgroup.DefaultStaticSitesUnlinkBackendFromBuildOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StaticSiteLinkedBackendARMResourceOperationGroupClient.StaticSitesValidateBackendForBuild`

```go
ctx := context.TODO()
id := staticsitelinkedbackendarmresourceoperationgroup.NewBuildLinkedBackendID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "buildName", "linkedBackendName")

payload := staticsitelinkedbackendarmresourceoperationgroup.StaticSiteLinkedBackendARMResource{
	// ...
}


if err := client.StaticSitesValidateBackendForBuildThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
