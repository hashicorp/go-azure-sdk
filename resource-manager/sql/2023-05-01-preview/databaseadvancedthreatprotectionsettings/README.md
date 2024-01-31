
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-05-01-preview/databaseadvancedthreatprotectionsettings` Documentation

The `databaseadvancedthreatprotectionsettings` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2023-05-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-05-01-preview/databaseadvancedthreatprotectionsettings"
```


### Client Initialization

```go
client := databaseadvancedthreatprotectionsettings.NewDatabaseAdvancedThreatProtectionSettingsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DatabaseAdvancedThreatProtectionSettingsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := databaseadvancedthreatprotectionsettings.NewSqlDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue")

payload := databaseadvancedthreatprotectionsettings.DatabaseAdvancedThreatProtection{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DatabaseAdvancedThreatProtectionSettingsClient.Get`

```go
ctx := context.TODO()
id := databaseadvancedthreatprotectionsettings.NewSqlDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DatabaseAdvancedThreatProtectionSettingsClient.ListByDatabase`

```go
ctx := context.TODO()
id := databaseadvancedthreatprotectionsettings.NewSqlDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue")

// alternatively `client.ListByDatabase(ctx, id)` can be used to do batched pagination
items, err := client.ListByDatabaseComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
