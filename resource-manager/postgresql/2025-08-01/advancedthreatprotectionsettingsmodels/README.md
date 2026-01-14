
## `github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/advancedthreatprotectionsettingsmodels` Documentation

The `advancedthreatprotectionsettingsmodels` SDK allows for interaction with Azure Resource Manager `postgresql` (API Version `2025-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/advancedthreatprotectionsettingsmodels"
```


### Client Initialization

```go
client := advancedthreatprotectionsettingsmodels.NewAdvancedThreatProtectionSettingsModelsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AdvancedThreatProtectionSettingsModelsClient.AdvancedThreatProtectionSettingsGet`

```go
ctx := context.TODO()
id := advancedthreatprotectionsettingsmodels.NewFlexibleServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerName")

read, err := client.AdvancedThreatProtectionSettingsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdvancedThreatProtectionSettingsModelsClient.AdvancedThreatProtectionSettingsListByServer`

```go
ctx := context.TODO()
id := advancedthreatprotectionsettingsmodels.NewFlexibleServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerName")

// alternatively `client.AdvancedThreatProtectionSettingsListByServer(ctx, id)` can be used to do batched pagination
items, err := client.AdvancedThreatProtectionSettingsListByServerComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AdvancedThreatProtectionSettingsModelsClient.ServerThreatProtectionSettingsCreateOrUpdate`

```go
ctx := context.TODO()
id := advancedthreatprotectionsettingsmodels.NewFlexibleServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerName")

payload := advancedthreatprotectionsettingsmodels.AdvancedThreatProtectionSettingsModel{
	// ...
}


if err := client.ServerThreatProtectionSettingsCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
