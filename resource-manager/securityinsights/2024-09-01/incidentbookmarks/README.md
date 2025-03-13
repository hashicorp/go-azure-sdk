
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-09-01/incidentbookmarks` Documentation

The `incidentbookmarks` SDK allows for interaction with Azure Resource Manager `securityinsights` (API Version `2024-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-09-01/incidentbookmarks"
```


### Client Initialization

```go
client := incidentbookmarks.NewIncidentBookmarksClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IncidentBookmarksClient.IncidentsListBookmarks`

```go
ctx := context.TODO()
id := incidentbookmarks.NewIncidentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "incidentIdentifier")

read, err := client.IncidentsListBookmarks(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
