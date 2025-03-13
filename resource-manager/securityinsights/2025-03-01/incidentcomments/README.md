
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2025-03-01/incidentcomments` Documentation

The `incidentcomments` SDK allows for interaction with Azure Resource Manager `securityinsights` (API Version `2025-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2025-03-01/incidentcomments"
```


### Client Initialization

```go
client := incidentcomments.NewIncidentCommentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IncidentCommentsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := incidentcomments.NewCommentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "incidentId", "incidentCommentId")

payload := incidentcomments.IncidentComment{
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


### Example Usage: `IncidentCommentsClient.Delete`

```go
ctx := context.TODO()
id := incidentcomments.NewCommentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "incidentId", "incidentCommentId")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IncidentCommentsClient.Get`

```go
ctx := context.TODO()
id := incidentcomments.NewCommentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "incidentId", "incidentCommentId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IncidentCommentsClient.List`

```go
ctx := context.TODO()
id := incidentcomments.NewIncidentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "incidentIdentifier")

// alternatively `client.List(ctx, id, incidentcomments.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, incidentcomments.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
