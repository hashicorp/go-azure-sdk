
## `github.com/hashicorp/go-azure-sdk/data-plane/batch/2022-01-01.15.0/computenodes` Documentation

The `computenodes` SDK allows for interaction with <unknown source data type> `batch` (API Version `2022-01-01.15.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/batch/2022-01-01.15.0/computenodes"
```


### Client Initialization

```go
client := computenodes.NewComputeNodesClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `ComputeNodesClient.ComputeNodeAddUser`

```go
ctx := context.TODO()
id := computenodes.NewNodeID("poolId", "nodeId")

payload := computenodes.ComputeNodeUser{
	// ...
}


read, err := client.ComputeNodeAddUser(ctx, id, payload, computenodes.DefaultComputeNodeAddUserOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComputeNodesClient.ComputeNodeDeleteUser`

```go
ctx := context.TODO()
id := computenodes.NewUserID("poolId", "nodeId", "userName")

read, err := client.ComputeNodeDeleteUser(ctx, id, computenodes.DefaultComputeNodeDeleteUserOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComputeNodesClient.ComputeNodeDisableScheduling`

```go
ctx := context.TODO()
id := computenodes.NewNodeID("poolId", "nodeId")

payload := computenodes.NodeDisableSchedulingParameter{
	// ...
}


read, err := client.ComputeNodeDisableScheduling(ctx, id, payload, computenodes.DefaultComputeNodeDisableSchedulingOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComputeNodesClient.ComputeNodeEnableScheduling`

```go
ctx := context.TODO()
id := computenodes.NewNodeID("poolId", "nodeId")

read, err := client.ComputeNodeEnableScheduling(ctx, id, computenodes.DefaultComputeNodeEnableSchedulingOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComputeNodesClient.ComputeNodeExtensionGet`

```go
ctx := context.TODO()
id := computenodes.NewExtensionID("poolId", "nodeId", "extensionName")

read, err := client.ComputeNodeExtensionGet(ctx, id, computenodes.DefaultComputeNodeExtensionGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComputeNodesClient.ComputeNodeExtensionList`

```go
ctx := context.TODO()
id := computenodes.NewNodeID("poolId", "nodeId")

// alternatively `client.ComputeNodeExtensionList(ctx, id, computenodes.DefaultComputeNodeExtensionListOperationOptions())` can be used to do batched pagination
items, err := client.ComputeNodeExtensionListComplete(ctx, id, computenodes.DefaultComputeNodeExtensionListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ComputeNodesClient.ComputeNodeGet`

```go
ctx := context.TODO()
id := computenodes.NewNodeID("poolId", "nodeId")

read, err := client.ComputeNodeGet(ctx, id, computenodes.DefaultComputeNodeGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComputeNodesClient.ComputeNodeGetRemoteDesktop`

```go
ctx := context.TODO()
id := computenodes.NewNodeID("poolId", "nodeId")

read, err := client.ComputeNodeGetRemoteDesktop(ctx, id, computenodes.DefaultComputeNodeGetRemoteDesktopOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComputeNodesClient.ComputeNodeGetRemoteLoginSettings`

```go
ctx := context.TODO()
id := computenodes.NewNodeID("poolId", "nodeId")

read, err := client.ComputeNodeGetRemoteLoginSettings(ctx, id, computenodes.DefaultComputeNodeGetRemoteLoginSettingsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComputeNodesClient.ComputeNodeList`

```go
ctx := context.TODO()
id := computenodes.NewPoolID("poolId")

// alternatively `client.ComputeNodeList(ctx, id, computenodes.DefaultComputeNodeListOperationOptions())` can be used to do batched pagination
items, err := client.ComputeNodeListComplete(ctx, id, computenodes.DefaultComputeNodeListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ComputeNodesClient.ComputeNodeReboot`

```go
ctx := context.TODO()
id := computenodes.NewNodeID("poolId", "nodeId")

payload := computenodes.NodeRebootParameter{
	// ...
}


read, err := client.ComputeNodeReboot(ctx, id, payload, computenodes.DefaultComputeNodeRebootOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComputeNodesClient.ComputeNodeReimage`

```go
ctx := context.TODO()
id := computenodes.NewNodeID("poolId", "nodeId")

payload := computenodes.NodeReimageParameter{
	// ...
}


read, err := client.ComputeNodeReimage(ctx, id, payload, computenodes.DefaultComputeNodeReimageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComputeNodesClient.ComputeNodeUpdateUser`

```go
ctx := context.TODO()
id := computenodes.NewUserID("poolId", "nodeId", "userName")

payload := computenodes.NodeUpdateUserParameter{
	// ...
}


read, err := client.ComputeNodeUpdateUser(ctx, id, payload, computenodes.DefaultComputeNodeUpdateUserOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComputeNodesClient.ComputeNodeUploadBatchServiceLogs`

```go
ctx := context.TODO()
id := computenodes.NewNodeID("poolId", "nodeId")

payload := computenodes.UploadBatchServiceLogsConfiguration{
	// ...
}


read, err := client.ComputeNodeUploadBatchServiceLogs(ctx, id, payload, computenodes.DefaultComputeNodeUploadBatchServiceLogsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComputeNodesClient.PoolRemoveNodes`

```go
ctx := context.TODO()
id := computenodes.NewPoolID("poolId")

payload := computenodes.NodeRemoveParameter{
	// ...
}


read, err := client.PoolRemoveNodes(ctx, id, payload, computenodes.DefaultPoolRemoveNodesOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
