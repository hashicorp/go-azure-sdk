
## `github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2024-01-01/applicationgroup` Documentation

The `applicationgroup` SDK allows for interaction with the Azure Resource Manager Service `eventhub` (API Version `2024-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2024-01-01/applicationgroup"
```


### Client Initialization

```go
client := applicationgroup.NewApplicationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ApplicationGroupClient.CreateOrUpdateApplicationGroup`

```go
ctx := context.TODO()
id := applicationgroup.NewApplicationGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue", "applicationGroupValue")

payload := applicationgroup.ApplicationGroup{
	// ...
}


read, err := client.CreateOrUpdateApplicationGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationGroupClient.Delete`

```go
ctx := context.TODO()
id := applicationgroup.NewApplicationGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue", "applicationGroupValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationGroupClient.Get`

```go
ctx := context.TODO()
id := applicationgroup.NewApplicationGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue", "applicationGroupValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationGroupClient.ListByNamespace`

```go
ctx := context.TODO()
id := applicationgroup.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue")

// alternatively `client.ListByNamespace(ctx, id)` can be used to do batched pagination
items, err := client.ListByNamespaceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
