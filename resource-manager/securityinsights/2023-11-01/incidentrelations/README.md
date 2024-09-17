
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-11-01/incidentrelations` Documentation

The `incidentrelations` SDK allows for interaction with Azure Resource Manager `securityinsights` (API Version `2023-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-11-01/incidentrelations"
```


### Client Initialization

```go
client := incidentrelations.NewIncidentRelationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IncidentRelationsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := incidentrelations.NewRelationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "incidentIdValue", "relationValue")

payload := incidentrelations.Relation{
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


### Example Usage: `IncidentRelationsClient.Delete`

```go
ctx := context.TODO()
id := incidentrelations.NewRelationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "incidentIdValue", "relationValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IncidentRelationsClient.Get`

```go
ctx := context.TODO()
id := incidentrelations.NewRelationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "incidentIdValue", "relationValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IncidentRelationsClient.List`

```go
ctx := context.TODO()
id := incidentrelations.NewIncidentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "incidentIdValue")

// alternatively `client.List(ctx, id, incidentrelations.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, incidentrelations.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
