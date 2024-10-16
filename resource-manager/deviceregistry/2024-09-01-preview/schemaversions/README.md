
## `github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2024-09-01-preview/schemaversions` Documentation

The `schemaversions` SDK allows for interaction with Azure Resource Manager `deviceregistry` (API Version `2024-09-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2024-09-01-preview/schemaversions"
```


### Client Initialization

```go
client := schemaversions.NewSchemaVersionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SchemaVersionsClient.CreateOrReplace`

```go
ctx := context.TODO()
id := schemaversions.NewSchemaVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "schemaRegistryName", "schemaName", "schemaVersionName")

payload := schemaversions.SchemaVersion{
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


### Example Usage: `SchemaVersionsClient.Delete`

```go
ctx := context.TODO()
id := schemaversions.NewSchemaVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "schemaRegistryName", "schemaName", "schemaVersionName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SchemaVersionsClient.Get`

```go
ctx := context.TODO()
id := schemaversions.NewSchemaVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "schemaRegistryName", "schemaName", "schemaVersionName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SchemaVersionsClient.ListBySchema`

```go
ctx := context.TODO()
id := schemaversions.NewSchemaID("12345678-1234-9876-4563-123456789012", "example-resource-group", "schemaRegistryName", "schemaName")

// alternatively `client.ListBySchema(ctx, id)` can be used to do batched pagination
items, err := client.ListBySchemaComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
