
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/instancemsdeploystatusoperationgroup` Documentation

The `instancemsdeploystatusoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/instancemsdeploystatusoperationgroup"
```


### Client Initialization

```go
client := instancemsdeploystatusoperationgroup.NewInstanceMSDeployStatusOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InstanceMSDeployStatusOperationGroupClient.WebAppsCreateInstanceMSDeployOperationSlot`

```go
ctx := context.TODO()
id := instancemsdeploystatusoperationgroup.NewSlotInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "instanceId")

payload := instancemsdeploystatusoperationgroup.MSDeploy{
	// ...
}


if err := client.WebAppsCreateInstanceMSDeployOperationSlotThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `InstanceMSDeployStatusOperationGroupClient.WebAppsGetInstanceMSDeployLogSlot`

```go
ctx := context.TODO()
id := instancemsdeploystatusoperationgroup.NewSlotInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "instanceId")

read, err := client.WebAppsGetInstanceMSDeployLogSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InstanceMSDeployStatusOperationGroupClient.WebAppsGetInstanceMsDeployStatusSlot`

```go
ctx := context.TODO()
id := instancemsdeploystatusoperationgroup.NewSlotInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "instanceId")

read, err := client.WebAppsGetInstanceMsDeployStatusSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
