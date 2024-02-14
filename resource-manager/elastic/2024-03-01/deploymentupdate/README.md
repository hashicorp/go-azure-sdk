
## `github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2024-03-01/deploymentupdate` Documentation

The `deploymentupdate` SDK allows for interaction with the Azure Resource Manager Service `elastic` (API Version `2024-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2024-03-01/deploymentupdate"
```


### Client Initialization

```go
client := deploymentupdate.NewDeploymentUpdateClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeploymentUpdateClient.MonitorUpgrade`

```go
ctx := context.TODO()
id := deploymentupdate.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorValue")

payload := deploymentupdate.ElasticMonitorUpgrade{
	// ...
}


if err := client.MonitorUpgradeThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
