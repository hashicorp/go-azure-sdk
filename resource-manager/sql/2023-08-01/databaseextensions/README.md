
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01/databaseextensions` Documentation

The `databaseextensions` SDK allows for interaction with Azure Resource Manager `sql` (API Version `2023-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01/databaseextensions"
```


### Client Initialization

```go
client := databaseextensions.NewDatabaseExtensionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DatabaseExtensionsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := databaseextensions.NewExtensionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseName", "extensionName")

payload := databaseextensions.DatabaseExtensions{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DatabaseExtensionsClient.Get`

```go
ctx := context.TODO()
id := databaseextensions.NewExtensionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseName", "extensionName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DatabaseExtensionsClient.ListByDatabase`

```go
ctx := context.TODO()
id := commonids.NewSqlDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseName")

// alternatively `client.ListByDatabase(ctx, id)` can be used to do batched pagination
items, err := client.ListByDatabaseComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
