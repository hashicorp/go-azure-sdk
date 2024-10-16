
## `github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2024-09-01-preview/schemas` Documentation

The `schemas` SDK allows for interaction with Azure Resource Manager `deviceregistry` (API Version `2024-09-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2024-09-01-preview/schemas"
```


### Client Initialization

```go
client := schemas.NewSchemasClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SchemasClient.CreateOrReplace`

```go
ctx := context.TODO()
id := schemas.NewSchemaID("12345678-1234-9876-4563-123456789012", "example-resource-group", "schemaRegistryName", "schemaName")

payload := schemas.Schema{
	// ...
}


read, err := client.CreateOrReplace(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SchemasClient.Delete`

```go
ctx := context.TODO()
id := schemas.NewSchemaID("12345678-1234-9876-4563-123456789012", "example-resource-group", "schemaRegistryName", "schemaName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SchemasClient.Get`

```go
ctx := context.TODO()
id := schemas.NewSchemaID("12345678-1234-9876-4563-123456789012", "example-resource-group", "schemaRegistryName", "schemaName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SchemasClient.ListBySchemaRegistry`

```go
ctx := context.TODO()
id := schemas.NewSchemaRegistryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "schemaRegistryName")

// alternatively `client.ListBySchemaRegistry(ctx, id)` can be used to do batched pagination
items, err := client.ListBySchemaRegistryComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
