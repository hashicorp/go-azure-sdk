
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2025-09-01/manualtrigger` Documentation

The `manualtrigger` SDK allows for interaction with Azure Resource Manager `securityinsights` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2025-09-01/manualtrigger"
```


### Client Initialization

```go
client := manualtrigger.NewManualTriggerClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManualTriggerClient.EntitiesRunPlaybook`

```go
ctx := context.TODO()
id := manualtrigger.NewEntityID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "entityIdentifier")

payload := manualtrigger.EntityManualTriggerRequestBody{
	// ...
}


read, err := client.EntitiesRunPlaybook(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManualTriggerClient.IncidentsRunPlaybook`

```go
ctx := context.TODO()
id := manualtrigger.NewIncidentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "incidentIdentifier")

payload := manualtrigger.ManualTriggerRequestBody{
	// ...
}


read, err := client.IncidentsRunPlaybook(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
