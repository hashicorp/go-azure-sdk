
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointcloudpc` Documentation

The `virtualendpointcloudpc` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointcloudpc"
```


### Client Initialization

```go
client := virtualendpointcloudpc.NewVirtualEndpointCloudPCClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VirtualEndpointCloudPCClient.CreateVirtualEndpointCloudPC`

```go
ctx := context.TODO()

payload := virtualendpointcloudpc.CloudPC{
	// ...
}


read, err := client.CreateVirtualEndpointCloudPC(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.CreateVirtualEndpointCloudPCBulkResize`

```go
ctx := context.TODO()

payload := virtualendpointcloudpc.CreateVirtualEndpointCloudPCBulkResizeRequest{
	// ...
}


read, err := client.CreateVirtualEndpointCloudPCBulkResize(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.CreateVirtualEndpointCloudPCChangeUserAccountType`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

payload := virtualendpointcloudpc.CreateVirtualEndpointCloudPCChangeUserAccountTypeRequest{
	// ...
}


read, err := client.CreateVirtualEndpointCloudPCChangeUserAccountType(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.CreateVirtualEndpointCloudPCCreateSnapshot`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

read, err := client.CreateVirtualEndpointCloudPCCreateSnapshot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.CreateVirtualEndpointCloudPCEndGracePeriod`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

read, err := client.CreateVirtualEndpointCloudPCEndGracePeriod(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.CreateVirtualEndpointCloudPCPowerOff`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

read, err := client.CreateVirtualEndpointCloudPCPowerOff(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.CreateVirtualEndpointCloudPCPowerOn`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

read, err := client.CreateVirtualEndpointCloudPCPowerOn(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.CreateVirtualEndpointCloudPCReboot`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

read, err := client.CreateVirtualEndpointCloudPCReboot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.CreateVirtualEndpointCloudPCRename`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

payload := virtualendpointcloudpc.CreateVirtualEndpointCloudPCRenameRequest{
	// ...
}


read, err := client.CreateVirtualEndpointCloudPCRename(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.CreateVirtualEndpointCloudPCReprovision`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

payload := virtualendpointcloudpc.CreateVirtualEndpointCloudPCReprovisionRequest{
	// ...
}


read, err := client.CreateVirtualEndpointCloudPCReprovision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.CreateVirtualEndpointCloudPCResize`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

payload := virtualendpointcloudpc.CreateVirtualEndpointCloudPCResizeRequest{
	// ...
}


read, err := client.CreateVirtualEndpointCloudPCResize(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.CreateVirtualEndpointCloudPCRetryPartnerAgentInstallation`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

read, err := client.CreateVirtualEndpointCloudPCRetryPartnerAgentInstallation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.CreateVirtualEndpointCloudPCTroubleshoot`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

read, err := client.CreateVirtualEndpointCloudPCTroubleshoot(ctx, id)
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
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

read, err := client.DeleteVirtualEndpointCloudPC(ctx, id)
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
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

read, err := client.GetVirtualEndpointCloudPC(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.GetVirtualEndpointCloudPCCount`

```go
ctx := context.TODO()


read, err := client.GetVirtualEndpointCloudPCCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.ListVirtualEndpointCloudPCs`

```go
ctx := context.TODO()


// alternatively `client.ListVirtualEndpointCloudPCs(ctx)` can be used to do batched pagination
items, err := client.ListVirtualEndpointCloudPCsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VirtualEndpointCloudPCClient.RestoreDeviceManagementVirtualEndpointCloudPC`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

payload := virtualendpointcloudpc.RestoreDeviceManagementVirtualEndpointCloudPCRequest{
	// ...
}


read, err := client.RestoreDeviceManagementVirtualEndpointCloudPC(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.StartDeviceManagementVirtualEndpointCloudPC`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

read, err := client.StartDeviceManagementVirtualEndpointCloudPC(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.StopDeviceManagementVirtualEndpointCloudPC`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

read, err := client.StopDeviceManagementVirtualEndpointCloudPC(ctx, id)
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
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

payload := virtualendpointcloudpc.CloudPC{
	// ...
}


read, err := client.UpdateVirtualEndpointCloudPC(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.ValidateDeviceManagementVirtualEndpointCloudPCsBulkResize`

```go
ctx := context.TODO()

payload := virtualendpointcloudpc.ValidateDeviceManagementVirtualEndpointCloudPCsBulkResizeRequest{
	// ...
}


read, err := client.ValidateDeviceManagementVirtualEndpointCloudPCsBulkResize(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
