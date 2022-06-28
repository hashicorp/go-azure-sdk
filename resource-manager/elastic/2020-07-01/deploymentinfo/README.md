
## `github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2020-07-01/deploymentinfo` Documentation

The `deploymentinfo` SDK allows for interaction with the Azure Resource Manager Service `elastic` (API Version `2020-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2020-07-01/deploymentinfo"
```


### Client Initialization

```go
client := deploymentinfo.NewDeploymentInfoClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeploymentInfoClient.List`

```go
ctx := context.TODO()
id := deploymentinfo.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorValue")

read, err := client.List(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
