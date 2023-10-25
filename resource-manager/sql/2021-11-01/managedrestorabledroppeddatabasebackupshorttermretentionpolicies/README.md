
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-11-01/managedrestorabledroppeddatabasebackupshorttermretentionpolicies` Documentation

The `managedrestorabledroppeddatabasebackupshorttermretentionpolicies` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2021-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-11-01/managedrestorabledroppeddatabasebackupshorttermretentionpolicies"
```


### Client Initialization

```go
client := managedrestorabledroppeddatabasebackupshorttermretentionpolicies.NewManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := managedrestorabledroppeddatabasebackupshorttermretentionpolicies.NewManagedInstanceRestorableDroppedDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "restorableDroppedDatabaseIdValue")

payload := managedrestorabledroppeddatabasebackupshorttermretentionpolicies.ManagedBackupShortTermRetentionPolicy{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient.Get`

```go
ctx := context.TODO()
id := managedrestorabledroppeddatabasebackupshorttermretentionpolicies.NewManagedInstanceRestorableDroppedDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "restorableDroppedDatabaseIdValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient.ListByRestorableDroppedDatabase`

```go
ctx := context.TODO()
id := managedrestorabledroppeddatabasebackupshorttermretentionpolicies.NewManagedInstanceRestorableDroppedDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "restorableDroppedDatabaseIdValue")

// alternatively `client.ListByRestorableDroppedDatabase(ctx, id)` can be used to do batched pagination
items, err := client.ListByRestorableDroppedDatabaseComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient.Update`

```go
ctx := context.TODO()
id := managedrestorabledroppeddatabasebackupshorttermretentionpolicies.NewManagedInstanceRestorableDroppedDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "restorableDroppedDatabaseIdValue")

payload := managedrestorabledroppeddatabasebackupshorttermretentionpolicies.ManagedBackupShortTermRetentionPolicy{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
