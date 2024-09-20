
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-11-01/incidents` Documentation

The `incidents` SDK allows for interaction with Azure Resource Manager `securityinsights` (API Version `2023-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-11-01/incidents"
```


### Client Initialization

```go
client := incidents.NewIncidentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IncidentsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := incidents.NewIncidentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "incidentId")

payload := incidents.Incident{
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


### Example Usage: `IncidentsClient.Delete`

```go
ctx := context.TODO()
id := incidents.NewIncidentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "incidentId")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IncidentsClient.Get`

```go
ctx := context.TODO()
id := incidents.NewIncidentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "incidentId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IncidentsClient.List`

```go
ctx := context.TODO()
id := incidents.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

// alternatively `client.List(ctx, id, incidents.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, incidents.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
