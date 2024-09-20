
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/checkdataconnectorrequirements` Documentation

The `checkdataconnectorrequirements` SDK allows for interaction with Azure Resource Manager `securityinsights` (API Version `2022-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/checkdataconnectorrequirements"
```


### Client Initialization

```go
client := checkdataconnectorrequirements.NewCheckDataConnectorRequirementsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CheckDataConnectorRequirementsClient.DataConnectorsCheckRequirementsPost`

```go
ctx := context.TODO()
id := checkdataconnectorrequirements.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

payload := checkdataconnectorrequirements.DataConnectorsCheckRequirements{
	// ...
}


read, err := client.DataConnectorsCheckRequirementsPost(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
