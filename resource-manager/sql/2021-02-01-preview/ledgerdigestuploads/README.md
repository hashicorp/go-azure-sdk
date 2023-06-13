
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/ledgerdigestuploads` Documentation

The `ledgerdigestuploads` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2021-02-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/ledgerdigestuploads"
```


### Client Initialization

```go
client := ledgerdigestuploads.NewLedgerDigestUploadsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LedgerDigestUploadsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := ledgerdigestuploads.NewDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue")

payload := ledgerdigestuploads.LedgerDigestUploads{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `LedgerDigestUploadsClient.Disable`

```go
ctx := context.TODO()
id := ledgerdigestuploads.NewDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue")

if err := client.DisableThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `LedgerDigestUploadsClient.Get`

```go
ctx := context.TODO()
id := ledgerdigestuploads.NewDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LedgerDigestUploadsClient.ListByDatabase`

```go
ctx := context.TODO()
id := ledgerdigestuploads.NewDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue")

// alternatively `client.ListByDatabase(ctx, id)` can be used to do batched pagination
items, err := client.ListByDatabaseComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
