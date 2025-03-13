
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/incidentalerts` Documentation

The `incidentalerts` SDK allows for interaction with Azure Resource Manager `securityinsights` (API Version `2023-12-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/incidentalerts"
```


### Client Initialization

```go
client := incidentalerts.NewIncidentAlertsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IncidentAlertsClient.IncidentsListAlerts`

```go
ctx := context.TODO()
id := incidentalerts.NewIncidentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "incidentIdentifier")

read, err := client.IncidentsListAlerts(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
