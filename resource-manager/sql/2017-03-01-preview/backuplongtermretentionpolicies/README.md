
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/backuplongtermretentionpolicies` Documentation

The `backuplongtermretentionpolicies` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2017-03-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/backuplongtermretentionpolicies"
```


### Client Initialization

```go
client := backuplongtermretentionpolicies.NewBackupLongTermRetentionPoliciesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BackupLongTermRetentionPoliciesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := backuplongtermretentionpolicies.NewDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue")

payload := backuplongtermretentionpolicies.BackupLongTermRetentionPolicy{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BackupLongTermRetentionPoliciesClient.Get`

```go
ctx := context.TODO()
id := backuplongtermretentionpolicies.NewDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BackupLongTermRetentionPoliciesClient.ListByDatabase`

```go
ctx := context.TODO()
id := backuplongtermretentionpolicies.NewDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue")

read, err := client.ListByDatabase(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
