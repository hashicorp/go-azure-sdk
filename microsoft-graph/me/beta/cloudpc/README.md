
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/cloudpc` Documentation

The `cloudpc` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/cloudpc"
```


### Client Initialization

```go
client := cloudpc.NewCloudPCClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CloudPCClient.CreateCloudPC`

```go
ctx := context.TODO()

payload := cloudpc.CloudPC{
	// ...
}


read, err := client.CreateCloudPC(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.CreateCloudPCBulkResize`

```go
ctx := context.TODO()

payload := cloudpc.CreateCloudPCBulkResizeRequest{
	// ...
}


read, err := client.CreateCloudPCBulkResize(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.CreateCloudPCChangeUserAccountType`

```go
ctx := context.TODO()
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

payload := cloudpc.CreateCloudPCChangeUserAccountTypeRequest{
	// ...
}


read, err := client.CreateCloudPCChangeUserAccountType(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.CreateCloudPCCreateSnapshot`

```go
ctx := context.TODO()
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

read, err := client.CreateCloudPCCreateSnapshot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.CreateCloudPCEndGracePeriod`

```go
ctx := context.TODO()
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

read, err := client.CreateCloudPCEndGracePeriod(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.CreateCloudPCPowerOff`

```go
ctx := context.TODO()
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

read, err := client.CreateCloudPCPowerOff(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.CreateCloudPCPowerOn`

```go
ctx := context.TODO()
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

read, err := client.CreateCloudPCPowerOn(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.CreateCloudPCReboot`

```go
ctx := context.TODO()
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

read, err := client.CreateCloudPCReboot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.CreateCloudPCRename`

```go
ctx := context.TODO()
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

payload := cloudpc.CreateCloudPCRenameRequest{
	// ...
}


read, err := client.CreateCloudPCRename(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.CreateCloudPCReprovision`

```go
ctx := context.TODO()
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

payload := cloudpc.CreateCloudPCReprovisionRequest{
	// ...
}


read, err := client.CreateCloudPCReprovision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.CreateCloudPCResize`

```go
ctx := context.TODO()
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

payload := cloudpc.CreateCloudPCResizeRequest{
	// ...
}


read, err := client.CreateCloudPCResize(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.CreateCloudPCRetryPartnerAgentInstallation`

```go
ctx := context.TODO()
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

read, err := client.CreateCloudPCRetryPartnerAgentInstallation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.CreateCloudPCTroubleshoot`

```go
ctx := context.TODO()
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

read, err := client.CreateCloudPCTroubleshoot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.DeleteCloudPC`

```go
ctx := context.TODO()
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

read, err := client.DeleteCloudPC(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.GetCloudPC`

```go
ctx := context.TODO()
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

read, err := client.GetCloudPC(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.GetCloudPCCount`

```go
ctx := context.TODO()


read, err := client.GetCloudPCCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.ListCloudPCs`

```go
ctx := context.TODO()


// alternatively `client.ListCloudPCs(ctx)` can be used to do batched pagination
items, err := client.ListCloudPCsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CloudPCClient.RestoreMeCloudPC`

```go
ctx := context.TODO()
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

payload := cloudpc.RestoreMeCloudPCRequest{
	// ...
}


read, err := client.RestoreMeCloudPC(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.StartMeCloudPC`

```go
ctx := context.TODO()
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

read, err := client.StartMeCloudPC(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.StopMeCloudPC`

```go
ctx := context.TODO()
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

read, err := client.StopMeCloudPC(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.UpdateCloudPC`

```go
ctx := context.TODO()
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

payload := cloudpc.CloudPC{
	// ...
}


read, err := client.UpdateCloudPC(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.ValidateMeCloudPCsBulkResize`

```go
ctx := context.TODO()

payload := cloudpc.ValidateMeCloudPCsBulkResizeRequest{
	// ...
}


read, err := client.ValidateMeCloudPCsBulkResize(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
