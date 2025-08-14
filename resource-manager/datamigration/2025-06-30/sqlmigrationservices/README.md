
## `github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2025-06-30/sqlmigrationservices` Documentation

The `sqlmigrationservices` SDK allows for interaction with Azure Resource Manager `datamigration` (API Version `2025-06-30`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2025-06-30/sqlmigrationservices"
```


### Client Initialization

```go
client := sqlmigrationservices.NewSqlMigrationServicesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SqlMigrationServicesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := sqlmigrationservices.NewSqlMigrationServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sqlMigrationServiceName")

payload := sqlmigrationservices.SqlMigrationService{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SqlMigrationServicesClient.Delete`

```go
ctx := context.TODO()
id := sqlmigrationservices.NewSqlMigrationServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sqlMigrationServiceName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SqlMigrationServicesClient.DeleteNode`

```go
ctx := context.TODO()
id := sqlmigrationservices.NewSqlMigrationServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sqlMigrationServiceName")

payload := sqlmigrationservices.DeleteNode{
	// ...
}


read, err := client.DeleteNode(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SqlMigrationServicesClient.Get`

```go
ctx := context.TODO()
id := sqlmigrationservices.NewSqlMigrationServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sqlMigrationServiceName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SqlMigrationServicesClient.ListAuthKeys`

```go
ctx := context.TODO()
id := sqlmigrationservices.NewSqlMigrationServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sqlMigrationServiceName")

read, err := client.ListAuthKeys(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SqlMigrationServicesClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SqlMigrationServicesClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SqlMigrationServicesClient.ListMigrations`

```go
ctx := context.TODO()
id := sqlmigrationservices.NewSqlMigrationServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sqlMigrationServiceName")

// alternatively `client.ListMigrations(ctx, id)` can be used to do batched pagination
items, err := client.ListMigrationsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SqlMigrationServicesClient.ListMonitoringData`

```go
ctx := context.TODO()
id := sqlmigrationservices.NewSqlMigrationServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sqlMigrationServiceName")

read, err := client.ListMonitoringData(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SqlMigrationServicesClient.RegenerateAuthKeys`

```go
ctx := context.TODO()
id := sqlmigrationservices.NewSqlMigrationServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sqlMigrationServiceName")

payload := sqlmigrationservices.RegenAuthKeys{
	// ...
}


read, err := client.RegenerateAuthKeys(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SqlMigrationServicesClient.Update`

```go
ctx := context.TODO()
id := sqlmigrationservices.NewSqlMigrationServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sqlMigrationServiceName")

payload := sqlmigrationservices.SqlMigrationServiceUpdate{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
