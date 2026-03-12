
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/staticsitebuildarmresources` Documentation

The `staticsitebuildarmresources` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/staticsitebuildarmresources"
```


### Client Initialization

```go
client := staticsitebuildarmresources.NewStaticSiteBuildARMResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `StaticSiteBuildARMResourcesClient.StaticSitesCreateOrUpdateStaticSiteBuildAppSettings`

```go
ctx := context.TODO()
id := staticsitebuildarmresources.NewBuildID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "buildName")

payload := staticsitebuildarmresources.StringDictionary{
	// ...
}


read, err := client.StaticSitesCreateOrUpdateStaticSiteBuildAppSettings(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StaticSiteBuildARMResourcesClient.StaticSitesCreateOrUpdateStaticSiteBuildFunctionAppSettings`

```go
ctx := context.TODO()
id := staticsitebuildarmresources.NewBuildID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "buildName")

payload := staticsitebuildarmresources.StringDictionary{
	// ...
}


read, err := client.StaticSitesCreateOrUpdateStaticSiteBuildFunctionAppSettings(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StaticSiteBuildARMResourcesClient.StaticSitesCreateZipDeploymentForStaticSiteBuild`

```go
ctx := context.TODO()
id := staticsitebuildarmresources.NewBuildID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "buildName")

payload := staticsitebuildarmresources.StaticSiteZipDeploymentARMResource{
	// ...
}


if err := client.StaticSitesCreateZipDeploymentForStaticSiteBuildThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `StaticSiteBuildARMResourcesClient.StaticSitesDeleteStaticSiteBuild`

```go
ctx := context.TODO()
id := staticsitebuildarmresources.NewBuildID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "buildName")

if err := client.StaticSitesDeleteStaticSiteBuildThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `StaticSiteBuildARMResourcesClient.StaticSitesGetBuildDatabaseConnectionsWithDetails`

```go
ctx := context.TODO()
id := staticsitebuildarmresources.NewBuildID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "buildName")

// alternatively `client.StaticSitesGetBuildDatabaseConnectionsWithDetails(ctx, id)` can be used to do batched pagination
items, err := client.StaticSitesGetBuildDatabaseConnectionsWithDetailsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `StaticSiteBuildARMResourcesClient.StaticSitesGetStaticSiteBuild`

```go
ctx := context.TODO()
id := staticsitebuildarmresources.NewBuildID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "buildName")

read, err := client.StaticSitesGetStaticSiteBuild(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StaticSiteBuildARMResourcesClient.StaticSitesGetStaticSiteBuilds`

```go
ctx := context.TODO()
id := staticsitebuildarmresources.NewStaticSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName")

// alternatively `client.StaticSitesGetStaticSiteBuilds(ctx, id)` can be used to do batched pagination
items, err := client.StaticSitesGetStaticSiteBuildsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `StaticSiteBuildARMResourcesClient.StaticSitesListStaticSiteBuildAppSettings`

```go
ctx := context.TODO()
id := staticsitebuildarmresources.NewBuildID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "buildName")

read, err := client.StaticSitesListStaticSiteBuildAppSettings(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StaticSiteBuildARMResourcesClient.StaticSitesListStaticSiteBuildFunctionAppSettings`

```go
ctx := context.TODO()
id := staticsitebuildarmresources.NewBuildID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "buildName")

read, err := client.StaticSitesListStaticSiteBuildFunctionAppSettings(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StaticSiteBuildARMResourcesClient.StaticSitesListStaticSiteBuildFunctions`

```go
ctx := context.TODO()
id := staticsitebuildarmresources.NewBuildID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "buildName")

// alternatively `client.StaticSitesListStaticSiteBuildFunctions(ctx, id)` can be used to do batched pagination
items, err := client.StaticSitesListStaticSiteBuildFunctionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
