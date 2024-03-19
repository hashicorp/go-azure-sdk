
## `github.com/hashicorp/go-azure-sdk/resource-manager/eventgrid/2023-12-15-preview/clientgroups` Documentation

The `clientgroups` SDK allows for interaction with the Azure Resource Manager Service `eventgrid` (API Version `2023-12-15-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/eventgrid/2023-12-15-preview/clientgroups"
```


### Client Initialization

```go
client := clientgroups.NewClientGroupsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ClientGroupsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := clientgroups.NewClientGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue", "clientGroupValue")

payload := clientgroups.ClientGroup{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ClientGroupsClient.Delete`

```go
ctx := context.TODO()
id := clientgroups.NewClientGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue", "clientGroupValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ClientGroupsClient.Get`

```go
ctx := context.TODO()
id := clientgroups.NewClientGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue", "clientGroupValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ClientGroupsClient.ListByNamespace`

```go
ctx := context.TODO()
id := clientgroups.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue")

// alternatively `client.ListByNamespace(ctx, id, clientgroups.DefaultListByNamespaceOperationOptions())` can be used to do batched pagination
items, err := client.ListByNamespaceComplete(ctx, id, clientgroups.DefaultListByNamespaceOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
