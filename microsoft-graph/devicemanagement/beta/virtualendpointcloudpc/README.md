
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


### Example Usage: `VirtualEndpointCloudPCClient.ChangeVirtualEndpointCloudPCUserAccountType`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

payload := virtualendpointcloudpc.ChangeVirtualEndpointCloudPCUserAccountTypeRequest{
	// ...
}


read, err := client.ChangeVirtualEndpointCloudPCUserAccountType(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
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


### Example Usage: `VirtualEndpointCloudPCClient.CreateVirtualEndpointCloudPCSnapshot`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

read, err := client.CreateVirtualEndpointCloudPCSnapshot(ctx, id)
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
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

read, err := client.EndVirtualEndpointCloudPCGracePeriod(ctx, id)
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


### Example Usage: `VirtualEndpointCloudPCClient.ListVirtualEndpointCloudPCBulkResizes`

```go
ctx := context.TODO()

payload := virtualendpointcloudpc.ListVirtualEndpointCloudPCBulkResizesRequest{
	// ...
}


// alternatively `client.ListVirtualEndpointCloudPCBulkResizes(ctx, payload, virtualendpointcloudpc.DefaultListVirtualEndpointCloudPCBulkResizesOperationOptions())` can be used to do batched pagination
items, err := client.ListVirtualEndpointCloudPCBulkResizesComplete(ctx, payload, virtualendpointcloudpc.DefaultListVirtualEndpointCloudPCBulkResizesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
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


### Example Usage: `VirtualEndpointCloudPCClient.PowerOffVirtualEndpointCloudPC`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

read, err := client.PowerOffVirtualEndpointCloudPC(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.PowerOnVirtualEndpointCloudPC`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

read, err := client.PowerOnVirtualEndpointCloudPC(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.RebootVirtualEndpointCloudPC`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

read, err := client.RebootVirtualEndpointCloudPC(ctx, id)
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
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

payload := virtualendpointcloudpc.RenameVirtualEndpointCloudPCRequest{
	// ...
}


read, err := client.RenameVirtualEndpointCloudPC(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.ReprovisionVirtualEndpointCloudPC`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

payload := virtualendpointcloudpc.ReprovisionVirtualEndpointCloudPCRequest{
	// ...
}


read, err := client.ReprovisionVirtualEndpointCloudPC(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.ResizeVirtualEndpointCloudPC`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

payload := virtualendpointcloudpc.ResizeVirtualEndpointCloudPCRequest{
	// ...
}


read, err := client.ResizeVirtualEndpointCloudPC(ctx, id, payload)
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
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

payload := virtualendpointcloudpc.RestoreVirtualEndpointCloudPCRequest{
	// ...
}


read, err := client.RestoreVirtualEndpointCloudPC(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.RetryVirtualEndpointCloudPCPartnerAgentInstallation`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

read, err := client.RetryVirtualEndpointCloudPCPartnerAgentInstallation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.SetVirtualEndpointCloudPCReviewStatus`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

payload := virtualendpointcloudpc.SetVirtualEndpointCloudPCReviewStatusRequest{
	// ...
}


read, err := client.SetVirtualEndpointCloudPCReviewStatus(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.StartVirtualEndpointCloudPC`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

read, err := client.StartVirtualEndpointCloudPC(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointCloudPCClient.StopVirtualEndpointCloudPC`

```go
ctx := context.TODO()
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

read, err := client.StopVirtualEndpointCloudPC(ctx, id)
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
id := virtualendpointcloudpc.NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

read, err := client.TroubleshootVirtualEndpointCloudPC(ctx, id)
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


### Example Usage: `VirtualEndpointCloudPCClient.ValidateVirtualEndpointCloudPCsBulkResizes`

```go
ctx := context.TODO()

payload := virtualendpointcloudpc.ValidateVirtualEndpointCloudPCsBulkResizesRequest{
	// ...
}


// alternatively `client.ValidateVirtualEndpointCloudPCsBulkResizes(ctx, payload, virtualendpointcloudpc.DefaultValidateVirtualEndpointCloudPCsBulkResizesOperationOptions())` can be used to do batched pagination
items, err := client.ValidateVirtualEndpointCloudPCsBulkResizesComplete(ctx, payload, virtualendpointcloudpc.DefaultValidateVirtualEndpointCloudPCsBulkResizesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
