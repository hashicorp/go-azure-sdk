
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-11-01/restorepoints` Documentation

The `restorepoints` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2021-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-11-01/restorepoints"
```


### Client Initialization

```go
client := restorepoints.NewRestorePointsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RestorePointsClient.Create`

```go
ctx := context.TODO()
id := restorepoints.NewDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue")

payload := restorepoints.CreateDatabaseRestorePointDefinition{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `RestorePointsClient.Delete`

```go
ctx := context.TODO()
id := restorepoints.NewRestorePointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "restorePointValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RestorePointsClient.Get`

```go
ctx := context.TODO()
id := restorepoints.NewRestorePointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "restorePointValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RestorePointsClient.ListByDatabase`

```go
ctx := context.TODO()
id := restorepoints.NewDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue")

// alternatively `client.ListByDatabase(ctx, id)` can be used to do batched pagination
items, err := client.ListByDatabaseComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
