
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-02-01-preview/longtermretentionmanagedinstancebackups` Documentation

The `longtermretentionmanagedinstancebackups` SDK allows for interaction with Azure Resource Manager `sql` (API Version `2023-02-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-02-01-preview/longtermretentionmanagedinstancebackups"
```


### Client Initialization

```go
client := longtermretentionmanagedinstancebackups.NewLongTermRetentionManagedInstanceBackupsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LongTermRetentionManagedInstanceBackupsClient.Delete`

```go
ctx := context.TODO()
id := longtermretentionmanagedinstancebackups.NewLongTermRetentionManagedInstanceBackupID("12345678-1234-9876-4563-123456789012", "locationValue", "longTermRetentionManagedInstanceValue", "longTermRetentionDatabaseValue", "longTermRetentionManagedInstanceBackupValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `LongTermRetentionManagedInstanceBackupsClient.DeleteByResourceGroup`

```go
ctx := context.TODO()
id := longtermretentionmanagedinstancebackups.NewLongTermRetentionDatabaseLongTermRetentionManagedInstanceBackupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationValue", "longTermRetentionManagedInstanceValue", "longTermRetentionDatabaseValue", "longTermRetentionManagedInstanceBackupValue")

if err := client.DeleteByResourceGroupThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `LongTermRetentionManagedInstanceBackupsClient.Get`

```go
ctx := context.TODO()
id := longtermretentionmanagedinstancebackups.NewLongTermRetentionManagedInstanceBackupID("12345678-1234-9876-4563-123456789012", "locationValue", "longTermRetentionManagedInstanceValue", "longTermRetentionDatabaseValue", "longTermRetentionManagedInstanceBackupValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LongTermRetentionManagedInstanceBackupsClient.GetByResourceGroup`

```go
ctx := context.TODO()
id := longtermretentionmanagedinstancebackups.NewLongTermRetentionDatabaseLongTermRetentionManagedInstanceBackupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationValue", "longTermRetentionManagedInstanceValue", "longTermRetentionDatabaseValue", "longTermRetentionManagedInstanceBackupValue")

read, err := client.GetByResourceGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LongTermRetentionManagedInstanceBackupsClient.ListByDatabase`

```go
ctx := context.TODO()
id := longtermretentionmanagedinstancebackups.NewLongTermRetentionDatabaseID("12345678-1234-9876-4563-123456789012", "locationValue", "longTermRetentionManagedInstanceValue", "longTermRetentionDatabaseValue")

// alternatively `client.ListByDatabase(ctx, id, longtermretentionmanagedinstancebackups.DefaultListByDatabaseOperationOptions())` can be used to do batched pagination
items, err := client.ListByDatabaseComplete(ctx, id, longtermretentionmanagedinstancebackups.DefaultListByDatabaseOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `LongTermRetentionManagedInstanceBackupsClient.ListByInstance`

```go
ctx := context.TODO()
id := longtermretentionmanagedinstancebackups.NewLongTermRetentionManagedInstanceID("12345678-1234-9876-4563-123456789012", "locationValue", "longTermRetentionManagedInstanceValue")

// alternatively `client.ListByInstance(ctx, id, longtermretentionmanagedinstancebackups.DefaultListByInstanceOperationOptions())` can be used to do batched pagination
items, err := client.ListByInstanceComplete(ctx, id, longtermretentionmanagedinstancebackups.DefaultListByInstanceOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `LongTermRetentionManagedInstanceBackupsClient.ListByLocation`

```go
ctx := context.TODO()
id := longtermretentionmanagedinstancebackups.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

// alternatively `client.ListByLocation(ctx, id, longtermretentionmanagedinstancebackups.DefaultListByLocationOperationOptions())` can be used to do batched pagination
items, err := client.ListByLocationComplete(ctx, id, longtermretentionmanagedinstancebackups.DefaultListByLocationOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `LongTermRetentionManagedInstanceBackupsClient.ListByResourceGroupDatabase`

```go
ctx := context.TODO()
id := longtermretentionmanagedinstancebackups.NewLongTermRetentionManagedInstanceLongTermRetentionDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationValue", "longTermRetentionManagedInstanceValue", "longTermRetentionDatabaseValue")

// alternatively `client.ListByResourceGroupDatabase(ctx, id, longtermretentionmanagedinstancebackups.DefaultListByResourceGroupDatabaseOperationOptions())` can be used to do batched pagination
items, err := client.ListByResourceGroupDatabaseComplete(ctx, id, longtermretentionmanagedinstancebackups.DefaultListByResourceGroupDatabaseOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `LongTermRetentionManagedInstanceBackupsClient.ListByResourceGroupInstance`

```go
ctx := context.TODO()
id := longtermretentionmanagedinstancebackups.NewLocationLongTermRetentionManagedInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationValue", "longTermRetentionManagedInstanceValue")

// alternatively `client.ListByResourceGroupInstance(ctx, id, longtermretentionmanagedinstancebackups.DefaultListByResourceGroupInstanceOperationOptions())` can be used to do batched pagination
items, err := client.ListByResourceGroupInstanceComplete(ctx, id, longtermretentionmanagedinstancebackups.DefaultListByResourceGroupInstanceOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `LongTermRetentionManagedInstanceBackupsClient.ListByResourceGroupLocation`

```go
ctx := context.TODO()
id := longtermretentionmanagedinstancebackups.NewProviderLocationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationValue")

// alternatively `client.ListByResourceGroupLocation(ctx, id, longtermretentionmanagedinstancebackups.DefaultListByResourceGroupLocationOperationOptions())` can be used to do batched pagination
items, err := client.ListByResourceGroupLocationComplete(ctx, id, longtermretentionmanagedinstancebackups.DefaultListByResourceGroupLocationOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
