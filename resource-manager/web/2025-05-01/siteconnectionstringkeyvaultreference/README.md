
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/siteconnectionstringkeyvaultreference` Documentation

The `siteconnectionstringkeyvaultreference` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/siteconnectionstringkeyvaultreference"
```


### Client Initialization

```go
client := siteconnectionstringkeyvaultreference.NewSiteConnectionStringKeyVaultReferenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteConnectionStringKeyVaultReferenceClient.WebAppsGetSiteConnectionStringKeyVaultReference`

```go
ctx := context.TODO()
id := siteconnectionstringkeyvaultreference.NewConnectionStringID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "connectionStringKey")

read, err := client.WebAppsGetSiteConnectionStringKeyVaultReference(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteConnectionStringKeyVaultReferenceClient.WebAppsGetSiteConnectionStringKeyVaultReferences`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

// alternatively `client.WebAppsGetSiteConnectionStringKeyVaultReferences(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsGetSiteConnectionStringKeyVaultReferencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
