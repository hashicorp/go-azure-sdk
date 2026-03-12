
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/staticsiteuserprovidedfunctionapparmresourceoperationgroup` Documentation

The `staticsiteuserprovidedfunctionapparmresourceoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/staticsiteuserprovidedfunctionapparmresourceoperationgroup"
```


### Client Initialization

```go
client := staticsiteuserprovidedfunctionapparmresourceoperationgroup.NewStaticSiteUserProvidedFunctionAppARMResourceOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `StaticSiteUserProvidedFunctionAppARMResourceOperationGroupClient.StaticSitesDetachUserProvidedFunctionAppFromStaticSite`

```go
ctx := context.TODO()
id := staticsiteuserprovidedfunctionapparmresourceoperationgroup.NewUserProvidedFunctionAppID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "userProvidedFunctionAppName")

read, err := client.StaticSitesDetachUserProvidedFunctionAppFromStaticSite(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StaticSiteUserProvidedFunctionAppARMResourceOperationGroupClient.StaticSitesGetUserProvidedFunctionAppForStaticSite`

```go
ctx := context.TODO()
id := staticsiteuserprovidedfunctionapparmresourceoperationgroup.NewUserProvidedFunctionAppID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "userProvidedFunctionAppName")

read, err := client.StaticSitesGetUserProvidedFunctionAppForStaticSite(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StaticSiteUserProvidedFunctionAppARMResourceOperationGroupClient.StaticSitesGetUserProvidedFunctionAppsForStaticSite`

```go
ctx := context.TODO()
id := staticsiteuserprovidedfunctionapparmresourceoperationgroup.NewStaticSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName")

// alternatively `client.StaticSitesGetUserProvidedFunctionAppsForStaticSite(ctx, id)` can be used to do batched pagination
items, err := client.StaticSitesGetUserProvidedFunctionAppsForStaticSiteComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `StaticSiteUserProvidedFunctionAppARMResourceOperationGroupClient.StaticSitesRegisterUserProvidedFunctionAppWithStaticSite`

```go
ctx := context.TODO()
id := staticsiteuserprovidedfunctionapparmresourceoperationgroup.NewUserProvidedFunctionAppID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "userProvidedFunctionAppName")

payload := staticsiteuserprovidedfunctionapparmresourceoperationgroup.StaticSiteUserProvidedFunctionAppARMResource{
	// ...
}


if err := client.StaticSitesRegisterUserProvidedFunctionAppWithStaticSiteThenPoll(ctx, id, payload, staticsiteuserprovidedfunctionapparmresourceoperationgroup.DefaultStaticSitesRegisterUserProvidedFunctionAppWithStaticSiteOperationOptions()); err != nil {
	// handle the error
}
```
