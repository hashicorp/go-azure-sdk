
## `github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-06-01-preview/advancedthreatprotectionsettings` Documentation

The `advancedthreatprotectionsettings` SDK allows for interaction with the Azure Resource Manager Service `postgresql` (API Version `2023-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-06-01-preview/advancedthreatprotectionsettings"
```


### Client Initialization

```go
client := advancedthreatprotectionsettings.NewAdvancedThreatProtectionSettingsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AdvancedThreatProtectionSettingsClient.ServerThreatProtectionSettingsCreateOrUpdate`

```go
ctx := context.TODO()
id := advancedthreatprotectionsettings.NewFlexibleServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerValue")

payload := advancedthreatprotectionsettings.ServerThreatProtectionSettingsModel{
	// ...
}


if err := client.ServerThreatProtectionSettingsCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AdvancedThreatProtectionSettingsClient.ServerThreatProtectionSettingsGet`

```go
ctx := context.TODO()
id := advancedthreatprotectionsettings.NewFlexibleServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerValue")

read, err := client.ServerThreatProtectionSettingsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdvancedThreatProtectionSettingsClient.ServerThreatProtectionSettingsListByServer`

```go
ctx := context.TODO()
id := advancedthreatprotectionsettings.NewFlexibleServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerValue")

// alternatively `client.ServerThreatProtectionSettingsListByServer(ctx, id)` can be used to do batched pagination
items, err := client.ServerThreatProtectionSettingsListByServerComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
