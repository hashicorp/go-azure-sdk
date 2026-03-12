
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/deployments` Documentation

The `deployments` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/deployments"
```


### Client Initialization

```go
client := deployments.NewDeploymentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeploymentsClient.WebAppsCreateDeployment`

```go
ctx := context.TODO()
id := deployments.NewDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "deploymentName")

payload := deployments.Deployment{
	// ...
}


read, err := client.WebAppsCreateDeployment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeploymentsClient.WebAppsDeleteDeployment`

```go
ctx := context.TODO()
id := deployments.NewDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "deploymentName")

read, err := client.WebAppsDeleteDeployment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeploymentsClient.WebAppsGetDeployment`

```go
ctx := context.TODO()
id := deployments.NewDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "deploymentName")

read, err := client.WebAppsGetDeployment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeploymentsClient.WebAppsListDeploymentLog`

```go
ctx := context.TODO()
id := deployments.NewDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "deploymentName")

read, err := client.WebAppsListDeploymentLog(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeploymentsClient.WebAppsListDeployments`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

// alternatively `client.WebAppsListDeployments(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListDeploymentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
