
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-03-01/manualtrigger` Documentation

The `manualtrigger` SDK allows for interaction with the Azure Resource Manager Service `securityinsights` (API Version `2024-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-03-01/manualtrigger"
```


### Client Initialization

```go
client := manualtrigger.NewManualTriggerClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManualTriggerClient.EntitiesRunPlaybook`

```go
ctx := context.TODO()
id := manualtrigger.NewEntityID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "entityIdentifierValue")

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
id := manualtrigger.NewIncidentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "incidentIdentifierValue")

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
