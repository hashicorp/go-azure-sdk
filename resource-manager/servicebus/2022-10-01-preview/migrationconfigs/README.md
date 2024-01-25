
## `github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2022-10-01-preview/migrationconfigs` Documentation

The `migrationconfigs` SDK allows for interaction with the Azure Resource Manager Service `servicebus` (API Version `2022-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2022-10-01-preview/migrationconfigs"
```


### Client Initialization

```go
client := migrationconfigs.NewMigrationConfigsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MigrationConfigsClient.CompleteMigration`

```go
ctx := context.TODO()
id := migrationconfigs.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue")

read, err := client.CompleteMigration(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MigrationConfigsClient.CreateAndStartMigration`

```go
ctx := context.TODO()
id := migrationconfigs.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue")

payload := migrationconfigs.MigrationConfigProperties{
	// ...
}


if err := client.CreateAndStartMigrationThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `MigrationConfigsClient.Delete`

```go
ctx := context.TODO()
id := migrationconfigs.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MigrationConfigsClient.Get`

```go
ctx := context.TODO()
id := migrationconfigs.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MigrationConfigsClient.List`

```go
ctx := context.TODO()
id := migrationconfigs.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MigrationConfigsClient.Revert`

```go
ctx := context.TODO()
id := migrationconfigs.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue")

read, err := client.Revert(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
