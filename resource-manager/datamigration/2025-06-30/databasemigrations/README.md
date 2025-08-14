
## `github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2025-06-30/databasemigrations` Documentation

The `databasemigrations` SDK allows for interaction with Azure Resource Manager `datamigration` (API Version `2025-06-30`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2025-06-30/databasemigrations"
```


### Client Initialization

```go
client := databasemigrations.NewDatabaseMigrationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DatabaseMigrationsClient.MongoToCosmosDbRUMongoCreate`

```go
ctx := context.TODO()
id := databasemigrations.NewDatabaseAccountProviders2DatabaseMigrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "databaseAccountName", "databaseMigrationName")

payload := databasemigrations.DatabaseMigrationCosmosDbMongo{
	// ...
}


if err := client.MongoToCosmosDbRUMongoCreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DatabaseMigrationsClient.MongoToCosmosDbRUMongoDelete`

```go
ctx := context.TODO()
id := databasemigrations.NewDatabaseAccountProviders2DatabaseMigrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "databaseAccountName", "databaseMigrationName")

if err := client.MongoToCosmosDbRUMongoDeleteThenPoll(ctx, id, databasemigrations.DefaultMongoToCosmosDbRUMongoDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `DatabaseMigrationsClient.MongoToCosmosDbRUMongoGet`

```go
ctx := context.TODO()
id := databasemigrations.NewDatabaseAccountProviders2DatabaseMigrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "databaseAccountName", "databaseMigrationName")

read, err := client.MongoToCosmosDbRUMongoGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DatabaseMigrationsClient.MongoToCosmosDbRUMongoGetForScope`

```go
ctx := context.TODO()
id := databasemigrations.NewDatabaseAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "databaseAccountName")

// alternatively `client.MongoToCosmosDbRUMongoGetForScope(ctx, id)` can be used to do batched pagination
items, err := client.MongoToCosmosDbRUMongoGetForScopeComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DatabaseMigrationsClient.MongoToCosmosDbvCoreMongoCreate`

```go
ctx := context.TODO()
id := databasemigrations.NewMongoClusterProviders2DatabaseMigrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "mongoClusterName", "databaseMigrationName")

payload := databasemigrations.DatabaseMigrationCosmosDbMongo{
	// ...
}


if err := client.MongoToCosmosDbvCoreMongoCreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DatabaseMigrationsClient.MongoToCosmosDbvCoreMongoDelete`

```go
ctx := context.TODO()
id := databasemigrations.NewMongoClusterProviders2DatabaseMigrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "mongoClusterName", "databaseMigrationName")

if err := client.MongoToCosmosDbvCoreMongoDeleteThenPoll(ctx, id, databasemigrations.DefaultMongoToCosmosDbvCoreMongoDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `DatabaseMigrationsClient.MongoToCosmosDbvCoreMongoGet`

```go
ctx := context.TODO()
id := databasemigrations.NewMongoClusterProviders2DatabaseMigrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "mongoClusterName", "databaseMigrationName")

read, err := client.MongoToCosmosDbvCoreMongoGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DatabaseMigrationsClient.MongoToCosmosDbvCoreMongoGetForScope`

```go
ctx := context.TODO()
id := databasemigrations.NewMongoClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "mongoClusterName")

// alternatively `client.MongoToCosmosDbvCoreMongoGetForScope(ctx, id)` can be used to do batched pagination
items, err := client.MongoToCosmosDbvCoreMongoGetForScopeComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DatabaseMigrationsClient.SqlDbCreateOrUpdate`

```go
ctx := context.TODO()
id := databasemigrations.NewDatabaseMigrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseMigrationName")

payload := databasemigrations.DatabaseMigrationSqlDb{
	// ...
}


if err := client.SqlDbCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DatabaseMigrationsClient.SqlDbDelete`

```go
ctx := context.TODO()
id := databasemigrations.NewDatabaseMigrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseMigrationName")

if err := client.SqlDbDeleteThenPoll(ctx, id, databasemigrations.DefaultSqlDbDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `DatabaseMigrationsClient.SqlDbGet`

```go
ctx := context.TODO()
id := databasemigrations.NewDatabaseMigrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseMigrationName")

read, err := client.SqlDbGet(ctx, id, databasemigrations.DefaultSqlDbGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DatabaseMigrationsClient.SqlDbcancel`

```go
ctx := context.TODO()
id := databasemigrations.NewDatabaseMigrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseMigrationName")

payload := databasemigrations.MigrationOperationInput{
	// ...
}


if err := client.SqlDbcancelThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DatabaseMigrationsClient.SqlDbretry`

```go
ctx := context.TODO()
id := databasemigrations.NewDatabaseMigrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseMigrationName")

payload := databasemigrations.MigrationOperationInput{
	// ...
}


if err := client.SqlDbretryThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DatabaseMigrationsClient.SqlMiCreateOrUpdate`

```go
ctx := context.TODO()
id := databasemigrations.NewProviders2DatabaseMigrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName", "databaseMigrationName")

payload := databasemigrations.DatabaseMigrationSqlMi{
	// ...
}


if err := client.SqlMiCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DatabaseMigrationsClient.SqlMiDelete`

```go
ctx := context.TODO()
id := databasemigrations.NewProviders2DatabaseMigrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName", "databaseMigrationName")

if err := client.SqlMiDeleteThenPoll(ctx, id, databasemigrations.DefaultSqlMiDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `DatabaseMigrationsClient.SqlMiGet`

```go
ctx := context.TODO()
id := databasemigrations.NewProviders2DatabaseMigrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName", "databaseMigrationName")

read, err := client.SqlMiGet(ctx, id, databasemigrations.DefaultSqlMiGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DatabaseMigrationsClient.SqlMicancel`

```go
ctx := context.TODO()
id := databasemigrations.NewProviders2DatabaseMigrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName", "databaseMigrationName")

payload := databasemigrations.MigrationOperationInput{
	// ...
}


if err := client.SqlMicancelThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DatabaseMigrationsClient.SqlMicutover`

```go
ctx := context.TODO()
id := databasemigrations.NewProviders2DatabaseMigrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName", "databaseMigrationName")

payload := databasemigrations.MigrationOperationInput{
	// ...
}


if err := client.SqlMicutoverThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DatabaseMigrationsClient.SqlVMCreateOrUpdate`

```go
ctx := context.TODO()
id := databasemigrations.NewSqlVirtualMachineProviders2DatabaseMigrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sqlVirtualMachineName", "databaseMigrationName")

payload := databasemigrations.DatabaseMigrationSqlVM{
	// ...
}


if err := client.SqlVMCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DatabaseMigrationsClient.SqlVMGet`

```go
ctx := context.TODO()
id := databasemigrations.NewSqlVirtualMachineProviders2DatabaseMigrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sqlVirtualMachineName", "databaseMigrationName")

read, err := client.SqlVMGet(ctx, id, databasemigrations.DefaultSqlVMGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DatabaseMigrationsClient.SqlVMcancel`

```go
ctx := context.TODO()
id := databasemigrations.NewSqlVirtualMachineProviders2DatabaseMigrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sqlVirtualMachineName", "databaseMigrationName")

payload := databasemigrations.MigrationOperationInput{
	// ...
}


if err := client.SqlVMcancelThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DatabaseMigrationsClient.SqlVMcutover`

```go
ctx := context.TODO()
id := databasemigrations.NewSqlVirtualMachineProviders2DatabaseMigrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sqlVirtualMachineName", "databaseMigrationName")

payload := databasemigrations.MigrationOperationInput{
	// ...
}


if err := client.SqlVMcutoverThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
