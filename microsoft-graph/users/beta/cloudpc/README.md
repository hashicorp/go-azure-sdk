
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/cloudpc` Documentation

The `cloudpc` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/cloudpc"
```


### Client Initialization

```go
client := cloudpc.NewCloudPCClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CloudPCClient.CreateCloudPC`

```go
ctx := context.TODO()
id := cloudpc.NewUserID("userIdValue")

payload := cloudpc.CloudPC{
	// ...
}


read, err := client.CreateCloudPC(ctx, id, payload)
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
id := cloudpc.NewUserID("userIdValue")

payload := cloudpc.CreateCloudPCBulkResizeRequest{
	// ...
}


read, err := client.CreateCloudPCBulkResize(ctx, id, payload)
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
id := cloudpc.NewUserIdCloudPCID("userIdValue", "cloudPCIdValue")

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
id := cloudpc.NewUserIdCloudPCID("userIdValue", "cloudPCIdValue")

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
id := cloudpc.NewUserIdCloudPCID("userIdValue", "cloudPCIdValue")

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
id := cloudpc.NewUserIdCloudPCID("userIdValue", "cloudPCIdValue")

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
id := cloudpc.NewUserIdCloudPCID("userIdValue", "cloudPCIdValue")

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
id := cloudpc.NewUserIdCloudPCID("userIdValue", "cloudPCIdValue")

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
id := cloudpc.NewUserIdCloudPCID("userIdValue", "cloudPCIdValue")

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
id := cloudpc.NewUserIdCloudPCID("userIdValue", "cloudPCIdValue")

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
id := cloudpc.NewUserIdCloudPCID("userIdValue", "cloudPCIdValue")

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
id := cloudpc.NewUserIdCloudPCID("userIdValue", "cloudPCIdValue")

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
id := cloudpc.NewUserIdCloudPCID("userIdValue", "cloudPCIdValue")

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
id := cloudpc.NewUserIdCloudPCID("userIdValue", "cloudPCIdValue")

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
id := cloudpc.NewUserIdCloudPCID("userIdValue", "cloudPCIdValue")

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
id := cloudpc.NewUserID("userIdValue")

read, err := client.GetCloudPCCount(ctx, id)
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
id := cloudpc.NewUserID("userIdValue")

// alternatively `client.ListCloudPCs(ctx, id)` can be used to do batched pagination
items, err := client.ListCloudPCsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CloudPCClient.RestoreUserCloudPC`

```go
ctx := context.TODO()
id := cloudpc.NewUserIdCloudPCID("userIdValue", "cloudPCIdValue")

payload := cloudpc.RestoreUserCloudPCRequest{
	// ...
}


read, err := client.RestoreUserCloudPC(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.StartUserCloudPC`

```go
ctx := context.TODO()
id := cloudpc.NewUserIdCloudPCID("userIdValue", "cloudPCIdValue")

read, err := client.StartUserCloudPC(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.StopUserCloudPC`

```go
ctx := context.TODO()
id := cloudpc.NewUserIdCloudPCID("userIdValue", "cloudPCIdValue")

read, err := client.StopUserCloudPC(ctx, id)
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
id := cloudpc.NewUserIdCloudPCID("userIdValue", "cloudPCIdValue")

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


### Example Usage: `CloudPCClient.ValidateUserCloudPCsBulkResize`

```go
ctx := context.TODO()
id := cloudpc.NewUserID("userIdValue")

payload := cloudpc.ValidateUserCloudPCsBulkResizeRequest{
	// ...
}


read, err := client.ValidateUserCloudPCsBulkResize(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
