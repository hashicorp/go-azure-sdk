
## `github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2024-08-01/migrations` Documentation

The `migrations` SDK allows for interaction with Azure Resource Manager `postgresql` (API Version `2024-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2024-08-01/migrations"
```


### Client Initialization

```go
client := migrations.NewMigrationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MigrationsClient.Create`

```go
ctx := context.TODO()
id := migrations.NewMigrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerName", "migrationName")

payload := migrations.MigrationResource{
	// ...
}


read, err := client.Create(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MigrationsClient.Delete`

```go
ctx := context.TODO()
id := migrations.NewMigrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerName", "migrationName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MigrationsClient.Get`

```go
ctx := context.TODO()
id := migrations.NewMigrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerName", "migrationName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MigrationsClient.ListByTargetServer`

```go
ctx := context.TODO()
id := migrations.NewFlexibleServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerName")

// alternatively `client.ListByTargetServer(ctx, id, migrations.DefaultListByTargetServerOperationOptions())` can be used to do batched pagination
items, err := client.ListByTargetServerComplete(ctx, id, migrations.DefaultListByTargetServerOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MigrationsClient.Update`

```go
ctx := context.TODO()
id := migrations.NewMigrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerName", "migrationName")

payload := migrations.MigrationResourceForPatch{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
