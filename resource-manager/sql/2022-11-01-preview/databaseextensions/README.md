
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2022-11-01-preview/databaseextensions` Documentation

The `databaseextensions` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2022-11-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2022-11-01-preview/databaseextensions"
```


### Client Initialization

```go
client := databaseextensions.NewDatabaseExtensionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DatabaseExtensionsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := databaseextensions.NewExtensionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "extensionValue")

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
id := databaseextensions.NewExtensionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "extensionValue")

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
id := databaseextensions.NewDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue")

// alternatively `client.ListByDatabase(ctx, id)` can be used to do batched pagination
items, err := client.ListByDatabaseComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
