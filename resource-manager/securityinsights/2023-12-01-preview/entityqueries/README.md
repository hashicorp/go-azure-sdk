
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/entityqueries` Documentation

The `entityqueries` SDK allows for interaction with Azure Resource Manager `securityinsights` (API Version `2023-12-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/entityqueries"
```


### Client Initialization

```go
client := entityqueries.NewEntityQueriesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EntityQueriesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := entityqueries.NewEntityQueryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "entityQueryId")

payload := entityqueries.CustomEntityQuery{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntityQueriesClient.Delete`

```go
ctx := context.TODO()
id := entityqueries.NewEntityQueryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "entityQueryId")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntityQueriesClient.EntityQueryTemplatesGet`

```go
ctx := context.TODO()
id := entityqueries.NewEntityQueryTemplateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "entityQueryTemplateId")

read, err := client.EntityQueryTemplatesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntityQueriesClient.EntityQueryTemplatesList`

```go
ctx := context.TODO()
id := entityqueries.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

// alternatively `client.EntityQueryTemplatesList(ctx, id, entityqueries.DefaultEntityQueryTemplatesListOperationOptions())` can be used to do batched pagination
items, err := client.EntityQueryTemplatesListComplete(ctx, id, entityqueries.DefaultEntityQueryTemplatesListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EntityQueriesClient.Get`

```go
ctx := context.TODO()
id := entityqueries.NewEntityQueryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "entityQueryId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntityQueriesClient.List`

```go
ctx := context.TODO()
id := entityqueries.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

// alternatively `client.List(ctx, id, entityqueries.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, entityqueries.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
