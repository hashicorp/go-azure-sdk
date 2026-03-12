
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/msdeploystatusoperationgroup` Documentation

The `msdeploystatusoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/msdeploystatusoperationgroup"
```


### Client Initialization

```go
client := msdeploystatusoperationgroup.NewMSDeployStatusOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MSDeployStatusOperationGroupClient.WebAppsCreateInstanceMSDeployOperation`

```go
ctx := context.TODO()
id := msdeploystatusoperationgroup.NewInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "instanceId")

payload := msdeploystatusoperationgroup.MSDeploy{
	// ...
}


if err := client.WebAppsCreateInstanceMSDeployOperationThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `MSDeployStatusOperationGroupClient.WebAppsGetInstanceMSDeployLog`

```go
ctx := context.TODO()
id := msdeploystatusoperationgroup.NewInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "instanceId")

read, err := client.WebAppsGetInstanceMSDeployLog(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MSDeployStatusOperationGroupClient.WebAppsGetInstanceMsDeployStatus`

```go
ctx := context.TODO()
id := msdeploystatusoperationgroup.NewInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "instanceId")

read, err := client.WebAppsGetInstanceMsDeployStatus(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
