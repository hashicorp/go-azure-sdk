
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/msdeploystatuses` Documentation

The `msdeploystatuses` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/msdeploystatuses"
```


### Client Initialization

```go
client := msdeploystatuses.NewMSDeployStatusesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MSDeployStatusesClient.WebAppsCreateMSDeployOperation`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

payload := msdeploystatuses.MSDeploy{
	// ...
}


if err := client.WebAppsCreateMSDeployOperationThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `MSDeployStatusesClient.WebAppsGetMSDeployLog`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsGetMSDeployLog(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MSDeployStatusesClient.WebAppsGetMSDeployStatus`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsGetMSDeployStatus(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
