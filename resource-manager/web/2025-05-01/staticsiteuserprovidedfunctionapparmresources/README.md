
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/staticsiteuserprovidedfunctionapparmresources` Documentation

The `staticsiteuserprovidedfunctionapparmresources` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/staticsiteuserprovidedfunctionapparmresources"
```


### Client Initialization

```go
client := staticsiteuserprovidedfunctionapparmresources.NewStaticSiteUserProvidedFunctionAppARMResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `StaticSiteUserProvidedFunctionAppARMResourcesClient.StaticSitesDetachUserProvidedFunctionAppFromStaticSiteBuild`

```go
ctx := context.TODO()
id := staticsiteuserprovidedfunctionapparmresources.NewBuildUserProvidedFunctionAppID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "buildName", "userProvidedFunctionAppName")

read, err := client.StaticSitesDetachUserProvidedFunctionAppFromStaticSiteBuild(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StaticSiteUserProvidedFunctionAppARMResourcesClient.StaticSitesGetUserProvidedFunctionAppForStaticSiteBuild`

```go
ctx := context.TODO()
id := staticsiteuserprovidedfunctionapparmresources.NewBuildUserProvidedFunctionAppID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "buildName", "userProvidedFunctionAppName")

read, err := client.StaticSitesGetUserProvidedFunctionAppForStaticSiteBuild(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StaticSiteUserProvidedFunctionAppARMResourcesClient.StaticSitesGetUserProvidedFunctionAppsForStaticSiteBuild`

```go
ctx := context.TODO()
id := staticsiteuserprovidedfunctionapparmresources.NewBuildID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "buildName")

// alternatively `client.StaticSitesGetUserProvidedFunctionAppsForStaticSiteBuild(ctx, id)` can be used to do batched pagination
items, err := client.StaticSitesGetUserProvidedFunctionAppsForStaticSiteBuildComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `StaticSiteUserProvidedFunctionAppARMResourcesClient.StaticSitesRegisterUserProvidedFunctionAppWithStaticSiteBuild`

```go
ctx := context.TODO()
id := staticsiteuserprovidedfunctionapparmresources.NewBuildUserProvidedFunctionAppID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "buildName", "userProvidedFunctionAppName")

payload := staticsiteuserprovidedfunctionapparmresources.StaticSiteUserProvidedFunctionAppARMResource{
	// ...
}


if err := client.StaticSitesRegisterUserProvidedFunctionAppWithStaticSiteBuildThenPoll(ctx, id, payload, staticsiteuserprovidedfunctionapparmresources.DefaultStaticSitesRegisterUserProvidedFunctionAppWithStaticSiteBuildOperationOptions()); err != nil {
	// handle the error
}
```
