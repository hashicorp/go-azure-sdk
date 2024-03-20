
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-03-01/incidententities` Documentation

The `incidententities` SDK allows for interaction with the Azure Resource Manager Service `securityinsights` (API Version `2024-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-03-01/incidententities"
```


### Client Initialization

```go
client := incidententities.NewIncidentEntitiesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IncidentEntitiesClient.IncidentsListEntities`

```go
ctx := context.TODO()
id := incidententities.NewIncidentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "incidentIdentifierValue")

read, err := client.IncidentsListEntities(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
