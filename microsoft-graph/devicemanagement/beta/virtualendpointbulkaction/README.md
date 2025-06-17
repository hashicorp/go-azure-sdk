
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointbulkaction` Documentation

The `virtualendpointbulkaction` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointbulkaction"
```


### Client Initialization

```go
client := virtualendpointbulkaction.NewVirtualEndpointBulkActionClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VirtualEndpointBulkActionClient.CreateVirtualEndpointBulkAction`

```go
ctx := context.TODO()

payload := virtualendpointbulkaction.CloudPCBulkAction{
	// ...
}


read, err := client.CreateVirtualEndpointBulkAction(ctx, payload, virtualendpointbulkaction.DefaultCreateVirtualEndpointBulkActionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointBulkActionClient.DeleteVirtualEndpointBulkAction`

```go
ctx := context.TODO()
id := virtualendpointbulkaction.NewDeviceManagementVirtualEndpointBulkActionID("cloudPCBulkActionId")

read, err := client.DeleteVirtualEndpointBulkAction(ctx, id, virtualendpointbulkaction.DefaultDeleteVirtualEndpointBulkActionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointBulkActionClient.GetVirtualEndpointBulkAction`

```go
ctx := context.TODO()
id := virtualendpointbulkaction.NewDeviceManagementVirtualEndpointBulkActionID("cloudPCBulkActionId")

read, err := client.GetVirtualEndpointBulkAction(ctx, id, virtualendpointbulkaction.DefaultGetVirtualEndpointBulkActionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointBulkActionClient.GetVirtualEndpointBulkActionsCount`

```go
ctx := context.TODO()


read, err := client.GetVirtualEndpointBulkActionsCount(ctx, virtualendpointbulkaction.DefaultGetVirtualEndpointBulkActionsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointBulkActionClient.ListVirtualEndpointBulkActions`

```go
ctx := context.TODO()


// alternatively `client.ListVirtualEndpointBulkActions(ctx, virtualendpointbulkaction.DefaultListVirtualEndpointBulkActionsOperationOptions())` can be used to do batched pagination
items, err := client.ListVirtualEndpointBulkActionsComplete(ctx, virtualendpointbulkaction.DefaultListVirtualEndpointBulkActionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VirtualEndpointBulkActionClient.RetryVirtualEndpointBulkAction`

```go
ctx := context.TODO()
id := virtualendpointbulkaction.NewDeviceManagementVirtualEndpointBulkActionID("cloudPCBulkActionId")

payload := virtualendpointbulkaction.RetryVirtualEndpointBulkActionRequest{
	// ...
}


read, err := client.RetryVirtualEndpointBulkAction(ctx, id, payload, virtualendpointbulkaction.DefaultRetryVirtualEndpointBulkActionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointBulkActionClient.UpdateVirtualEndpointBulkAction`

```go
ctx := context.TODO()
id := virtualendpointbulkaction.NewDeviceManagementVirtualEndpointBulkActionID("cloudPCBulkActionId")

payload := virtualendpointbulkaction.CloudPCBulkAction{
	// ...
}


read, err := client.UpdateVirtualEndpointBulkAction(ctx, id, payload, virtualendpointbulkaction.DefaultUpdateVirtualEndpointBulkActionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
