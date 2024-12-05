
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/entities` Documentation

The `entities` SDK allows for interaction with Azure Resource Manager `securityinsights` (API Version `2023-12-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/entities"
```


### Client Initialization

```go
client := entities.NewEntitiesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EntitiesClient.Expand`

```go
ctx := context.TODO()
id := entities.NewEntityID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "entityIdentifier")

payload := entities.EntityExpandParameters{
	// ...
}


read, err := client.Expand(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitiesClient.Get`

```go
ctx := context.TODO()
id := entities.NewEntityID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "entityIdentifier")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitiesClient.GetInsights`

```go
ctx := context.TODO()
id := entities.NewEntityID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "entityIdentifier")

payload := entities.EntityGetInsightsParameters{
	// ...
}


read, err := client.GetInsights(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitiesClient.GetTimelinelist`

```go
ctx := context.TODO()
id := entities.NewEntityID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "entityIdentifier")

payload := entities.EntityTimelineParameters{
	// ...
}


read, err := client.GetTimelinelist(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitiesClient.List`

```go
ctx := context.TODO()
id := entities.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EntitiesClient.Queries`

```go
ctx := context.TODO()
id := entities.NewEntityID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "entityIdentifier")

read, err := client.Queries(ctx, id, entities.DefaultQueriesOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
