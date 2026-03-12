
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/staticsitearmresources` Documentation

The `staticsitearmresources` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/staticsitearmresources"
```


### Client Initialization

```go
client := staticsitearmresources.NewStaticSiteARMResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `StaticSiteARMResourcesClient.StaticSitesCreateOrUpdateStaticSite`

```go
ctx := context.TODO()
id := staticsitearmresources.NewStaticSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName")

payload := staticsitearmresources.StaticSiteARMResource{
	// ...
}


if err := client.StaticSitesCreateOrUpdateStaticSiteThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `StaticSiteARMResourcesClient.StaticSitesCreateOrUpdateStaticSiteAppSettings`

```go
ctx := context.TODO()
id := staticsitearmresources.NewStaticSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName")

payload := staticsitearmresources.StringDictionary{
	// ...
}


read, err := client.StaticSitesCreateOrUpdateStaticSiteAppSettings(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StaticSiteARMResourcesClient.StaticSitesCreateOrUpdateStaticSiteFunctionAppSettings`

```go
ctx := context.TODO()
id := staticsitearmresources.NewStaticSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName")

payload := staticsitearmresources.StringDictionary{
	// ...
}


read, err := client.StaticSitesCreateOrUpdateStaticSiteFunctionAppSettings(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StaticSiteARMResourcesClient.StaticSitesCreateUserRolesInvitationLink`

```go
ctx := context.TODO()
id := staticsitearmresources.NewStaticSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName")

payload := staticsitearmresources.StaticSiteUserInvitationRequestResource{
	// ...
}


read, err := client.StaticSitesCreateUserRolesInvitationLink(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StaticSiteARMResourcesClient.StaticSitesCreateZipDeploymentForStaticSite`

```go
ctx := context.TODO()
id := staticsitearmresources.NewStaticSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName")

payload := staticsitearmresources.StaticSiteZipDeploymentARMResource{
	// ...
}


if err := client.StaticSitesCreateZipDeploymentForStaticSiteThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `StaticSiteARMResourcesClient.StaticSitesDeleteStaticSite`

```go
ctx := context.TODO()
id := staticsitearmresources.NewStaticSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName")

if err := client.StaticSitesDeleteStaticSiteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `StaticSiteARMResourcesClient.StaticSitesDeleteStaticSiteUser`

```go
ctx := context.TODO()
id := staticsitearmresources.NewUserID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "authProviderName", "userName")

read, err := client.StaticSitesDeleteStaticSiteUser(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StaticSiteARMResourcesClient.StaticSitesDetachStaticSite`

```go
ctx := context.TODO()
id := staticsitearmresources.NewStaticSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName")

if err := client.StaticSitesDetachStaticSiteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `StaticSiteARMResourcesClient.StaticSitesGetDatabaseConnectionsWithDetails`

```go
ctx := context.TODO()
id := staticsitearmresources.NewStaticSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName")

// alternatively `client.StaticSitesGetDatabaseConnectionsWithDetails(ctx, id)` can be used to do batched pagination
items, err := client.StaticSitesGetDatabaseConnectionsWithDetailsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `StaticSiteARMResourcesClient.StaticSitesGetPrivateLinkResources`

```go
ctx := context.TODO()
id := staticsitearmresources.NewStaticSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName")

read, err := client.StaticSitesGetPrivateLinkResources(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StaticSiteARMResourcesClient.StaticSitesGetStaticSite`

```go
ctx := context.TODO()
id := staticsitearmresources.NewStaticSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName")

read, err := client.StaticSitesGetStaticSite(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StaticSiteARMResourcesClient.StaticSitesGetStaticSitesByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.StaticSitesGetStaticSitesByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.StaticSitesGetStaticSitesByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `StaticSiteARMResourcesClient.StaticSitesList`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.StaticSitesList(ctx, id)` can be used to do batched pagination
items, err := client.StaticSitesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `StaticSiteARMResourcesClient.StaticSitesListStaticSiteAppSettings`

```go
ctx := context.TODO()
id := staticsitearmresources.NewStaticSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName")

read, err := client.StaticSitesListStaticSiteAppSettings(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StaticSiteARMResourcesClient.StaticSitesListStaticSiteConfiguredRoles`

```go
ctx := context.TODO()
id := staticsitearmresources.NewStaticSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName")

read, err := client.StaticSitesListStaticSiteConfiguredRoles(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StaticSiteARMResourcesClient.StaticSitesListStaticSiteFunctionAppSettings`

```go
ctx := context.TODO()
id := staticsitearmresources.NewStaticSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName")

read, err := client.StaticSitesListStaticSiteFunctionAppSettings(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StaticSiteARMResourcesClient.StaticSitesListStaticSiteFunctions`

```go
ctx := context.TODO()
id := staticsitearmresources.NewStaticSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName")

// alternatively `client.StaticSitesListStaticSiteFunctions(ctx, id)` can be used to do batched pagination
items, err := client.StaticSitesListStaticSiteFunctionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `StaticSiteARMResourcesClient.StaticSitesListStaticSiteSecrets`

```go
ctx := context.TODO()
id := staticsitearmresources.NewStaticSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName")

read, err := client.StaticSitesListStaticSiteSecrets(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StaticSiteARMResourcesClient.StaticSitesListStaticSiteUsers`

```go
ctx := context.TODO()
id := staticsitearmresources.NewAuthProviderID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "authProviderName")

// alternatively `client.StaticSitesListStaticSiteUsers(ctx, id)` can be used to do batched pagination
items, err := client.StaticSitesListStaticSiteUsersComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `StaticSiteARMResourcesClient.StaticSitesResetStaticSiteApiKey`

```go
ctx := context.TODO()
id := staticsitearmresources.NewStaticSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName")

payload := staticsitearmresources.StaticSiteResetPropertiesARMResource{
	// ...
}


read, err := client.StaticSitesResetStaticSiteApiKey(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StaticSiteARMResourcesClient.StaticSitesUpdateStaticSite`

```go
ctx := context.TODO()
id := staticsitearmresources.NewStaticSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName")

payload := staticsitearmresources.StaticSitePatchResource{
	// ...
}


read, err := client.StaticSitesUpdateStaticSite(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StaticSiteARMResourcesClient.StaticSitesUpdateStaticSiteUser`

```go
ctx := context.TODO()
id := staticsitearmresources.NewUserID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "authProviderName", "userName")

payload := staticsitearmresources.StaticSiteUserARMResource{
	// ...
}


read, err := client.StaticSitesUpdateStaticSiteUser(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
