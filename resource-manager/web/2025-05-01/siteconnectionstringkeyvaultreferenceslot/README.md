
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/siteconnectionstringkeyvaultreferenceslot` Documentation

The `siteconnectionstringkeyvaultreferenceslot` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/siteconnectionstringkeyvaultreferenceslot"
```


### Client Initialization

```go
client := siteconnectionstringkeyvaultreferenceslot.NewSiteConnectionStringKeyVaultReferenceSlotClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteConnectionStringKeyVaultReferenceSlotClient.WebAppsGetSiteConnectionStringKeyVaultReferenceSlot`

```go
ctx := context.TODO()
id := siteconnectionstringkeyvaultreferenceslot.NewConfigReferenceConnectionStringID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "connectionStringKey")

read, err := client.WebAppsGetSiteConnectionStringKeyVaultReferenceSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteConnectionStringKeyVaultReferenceSlotClient.WebAppsGetSiteConnectionStringKeyVaultReferencesSlot`

```go
ctx := context.TODO()
id := siteconnectionstringkeyvaultreferenceslot.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

// alternatively `client.WebAppsGetSiteConnectionStringKeyVaultReferencesSlot(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsGetSiteConnectionStringKeyVaultReferencesSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
