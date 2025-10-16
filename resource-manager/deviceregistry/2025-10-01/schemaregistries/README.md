
## `github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2025-10-01/schemaregistries` Documentation

The `schemaregistries` SDK allows for interaction with Azure Resource Manager `deviceregistry` (API Version `2025-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2025-10-01/schemaregistries"
```


### Client Initialization

```go
client := schemaregistries.NewSchemaRegistriesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SchemaRegistriesClient.CreateOrReplace`

```go
ctx := context.TODO()
id := schemaregistries.NewSchemaRegistryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "schemaRegistryName")

payload := schemaregistries.SchemaRegistry{
	// ...
}


if err := client.CreateOrReplaceThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SchemaRegistriesClient.Delete`

```go
ctx := context.TODO()
id := schemaregistries.NewSchemaRegistryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "schemaRegistryName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SchemaRegistriesClient.Get`

```go
ctx := context.TODO()
id := schemaregistries.NewSchemaRegistryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "schemaRegistryName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SchemaRegistriesClient.ListByResourceGroup`

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


### Example Usage: `SchemaRegistriesClient.ListBySubscription`

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


### Example Usage: `SchemaRegistriesClient.Update`

```go
ctx := context.TODO()
id := schemaregistries.NewSchemaRegistryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "schemaRegistryName")

payload := schemaregistries.SchemaRegistryUpdate{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
