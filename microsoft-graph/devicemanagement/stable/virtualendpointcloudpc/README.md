
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/virtualendpointcloudpc` Documentation

The `virtualendpointcloudpc` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/virtualendpointcloudpc"
```


### Client Initialization

```go
client := virtualendpointcloudpc.NewVirtualEndpointCloudPCClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VirtualEndpointCloudPCClient.CreateVirtualEndpointCloudPC`

```go
ctx := context.TODO()

payload := virtualendpointcloudpc.CloudPC{
	// ...
}


read, err := client.CreateVirtualEndpointCloudPC(ctx, payload, virtualendpointcloudpc.DefaultCreateVirtualEndpointCloudPCOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.DeleteVirtualEndpointCloudPC`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCId")

read, err := client.DeleteVirtualEndpointCloudPC(ctx, id, virtualendpointcloudpc.DefaultDeleteVirtualEndpointCloudPCOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.EndVirtualEndpointCloudPCGracePeriod`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCId")

read, err := client.EndVirtualEndpointCloudPCGracePeriod(ctx, id, virtualendpointcloudpc.DefaultEndVirtualEndpointCloudPCGracePeriodOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.GetVirtualEndpointCloudPC`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCId")

read, err := client.GetVirtualEndpointCloudPC(ctx, id, virtualendpointcloudpc.DefaultGetVirtualEndpointCloudPCOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.GetVirtualEndpointCloudPCsCount`

```go
ctx := context.TODO()


read, err := client.GetVirtualEndpointCloudPCsCount(ctx, virtualendpointcloudpc.DefaultGetVirtualEndpointCloudPCsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.ListVirtualEndpointCloudPCS`

```go
ctx := context.TODO()


// alternatively `client.ListVirtualEndpointCloudPCS(ctx, virtualendpointcloudpc.DefaultListVirtualEndpointCloudPCSOperationOptions())` can be used to do batched pagination
items, err := client.ListVirtualEndpointCloudPCSComplete(ctx, virtualendpointcloudpc.DefaultListVirtualEndpointCloudPCSOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VirtualEndpointCloudPCClient.RebootVirtualEndpointCloudPC`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCId")

read, err := client.RebootVirtualEndpointCloudPC(ctx, id, virtualendpointcloudpc.DefaultRebootVirtualEndpointCloudPCOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.RenameVirtualEndpointCloudPC`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCId")

payload := virtualendpointcloudpc.RenameVirtualEndpointCloudPCRequest{
	// ...
}


read, err := client.RenameVirtualEndpointCloudPC(ctx, id, payload, virtualendpointcloudpc.DefaultRenameVirtualEndpointCloudPCOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.RestoreVirtualEndpointCloudPC`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCId")

payload := virtualendpointcloudpc.RestoreVirtualEndpointCloudPCRequest{
	// ...
}


read, err := client.RestoreVirtualEndpointCloudPC(ctx, id, payload, virtualendpointcloudpc.DefaultRestoreVirtualEndpointCloudPCOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.TroubleshootVirtualEndpointCloudPC`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCId")

read, err := client.TroubleshootVirtualEndpointCloudPC(ctx, id, virtualendpointcloudpc.DefaultTroubleshootVirtualEndpointCloudPCOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.UpdateVirtualEndpointCloudPC`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCId")

payload := virtualendpointcloudpc.CloudPC{
	// ...
}


read, err := client.UpdateVirtualEndpointCloudPC(ctx, id, payload, virtualendpointcloudpc.DefaultUpdateVirtualEndpointCloudPCOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
