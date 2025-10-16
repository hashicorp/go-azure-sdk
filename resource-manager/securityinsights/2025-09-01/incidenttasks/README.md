
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2025-09-01/incidenttasks` Documentation

The `incidenttasks` SDK allows for interaction with Azure Resource Manager `securityinsights` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2025-09-01/incidenttasks"
```


### Client Initialization

```go
client := incidenttasks.NewIncidentTasksClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IncidentTasksClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := incidenttasks.NewTaskID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "incidentId", "incidentTaskId")

payload := incidenttasks.IncidentTask{
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


### Example Usage: `IncidentTasksClient.Delete`

```go
ctx := context.TODO()
id := incidenttasks.NewTaskID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "incidentId", "incidentTaskId")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IncidentTasksClient.Get`

```go
ctx := context.TODO()
id := incidenttasks.NewTaskID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "incidentId", "incidentTaskId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IncidentTasksClient.List`

```go
ctx := context.TODO()
id := incidenttasks.NewIncidentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "incidentIdentifier")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
