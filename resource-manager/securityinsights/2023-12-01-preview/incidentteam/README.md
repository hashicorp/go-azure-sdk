
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/incidentteam` Documentation

The `incidentteam` SDK allows for interaction with Azure Resource Manager `securityinsights` (API Version `2023-12-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/incidentteam"
```


### Client Initialization

```go
client := incidentteam.NewIncidentTeamClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IncidentTeamClient.IncidentsCreateTeam`

```go
ctx := context.TODO()
id := incidentteam.NewIncidentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "incidentIdentifier")

payload := incidentteam.TeamInformation{
	// ...
}


read, err := client.IncidentsCreateTeam(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
