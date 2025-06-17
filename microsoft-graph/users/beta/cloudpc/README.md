
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/cloudpc` Documentation

The `cloudpc` SDK allows for interaction with Microsoft Graph `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/cloudpc"
```


### Client Initialization

```go
client := cloudpc.NewCloudPCClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CloudPCClient.ChangeCloudPCUserAccountType`

```go
ctx := context.TODO()
id := cloudpc.NewUserIdCloudPCID("userId", "cloudPCId")

payload := cloudpc.ChangeCloudPCUserAccountTypeRequest{
	// ...
}


read, err := client.ChangeCloudPCUserAccountType(ctx, id, payload, cloudpc.DefaultChangeCloudPCUserAccountTypeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.CreateCloudPC`

```go
ctx := context.TODO()
id := cloudpc.NewUserID("userId")

payload := cloudpc.CloudPC{
	// ...
}


read, err := client.CreateCloudPC(ctx, id, payload, cloudpc.DefaultCreateCloudPCOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.CreateCloudPCSnapshot`

```go
ctx := context.TODO()
id := cloudpc.NewUserIdCloudPCID("userId", "cloudPCId")

payload := cloudpc.CreateCloudPCSnapshotRequest{
	// ...
}


read, err := client.CreateCloudPCSnapshot(ctx, id, payload, cloudpc.DefaultCreateCloudPCSnapshotOperationOptions())
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
id := cloudpc.NewUserIdCloudPCID("userId", "cloudPCId")

read, err := client.DeleteCloudPC(ctx, id, cloudpc.DefaultDeleteCloudPCOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.EndCloudPCGracePeriod`

```go
ctx := context.TODO()
id := cloudpc.NewUserIdCloudPCID("userId", "cloudPCId")

read, err := client.EndCloudPCGracePeriod(ctx, id, cloudpc.DefaultEndCloudPCGracePeriodOperationOptions())
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
id := cloudpc.NewUserIdCloudPCID("userId", "cloudPCId")

read, err := client.GetCloudPC(ctx, id, cloudpc.DefaultGetCloudPCOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.GetCloudPCsCount`

```go
ctx := context.TODO()
id := cloudpc.NewUserID("userId")

read, err := client.GetCloudPCsCount(ctx, id, cloudpc.DefaultGetCloudPCsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.ListCloudPCS`

```go
ctx := context.TODO()
id := cloudpc.NewUserID("userId")

// alternatively `client.ListCloudPCS(ctx, id, cloudpc.DefaultListCloudPCSOperationOptions())` can be used to do batched pagination
items, err := client.ListCloudPCSComplete(ctx, id, cloudpc.DefaultListCloudPCSOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CloudPCClient.PowerOffCloudPC`

```go
ctx := context.TODO()
id := cloudpc.NewUserIdCloudPCID("userId", "cloudPCId")

read, err := client.PowerOffCloudPC(ctx, id, cloudpc.DefaultPowerOffCloudPCOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.PowerOnCloudPC`

```go
ctx := context.TODO()
id := cloudpc.NewUserIdCloudPCID("userId", "cloudPCId")

read, err := client.PowerOnCloudPC(ctx, id, cloudpc.DefaultPowerOnCloudPCOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.RebootCloudPC`

```go
ctx := context.TODO()
id := cloudpc.NewUserIdCloudPCID("userId", "cloudPCId")

read, err := client.RebootCloudPC(ctx, id, cloudpc.DefaultRebootCloudPCOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.RenameCloudPC`

```go
ctx := context.TODO()
id := cloudpc.NewUserIdCloudPCID("userId", "cloudPCId")

payload := cloudpc.RenameCloudPCRequest{
	// ...
}


read, err := client.RenameCloudPC(ctx, id, payload, cloudpc.DefaultRenameCloudPCOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.ReprovisionCloudPC`

```go
ctx := context.TODO()
id := cloudpc.NewUserIdCloudPCID("userId", "cloudPCId")

payload := cloudpc.ReprovisionCloudPCRequest{
	// ...
}


read, err := client.ReprovisionCloudPC(ctx, id, payload, cloudpc.DefaultReprovisionCloudPCOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.ResizeCloudPC`

```go
ctx := context.TODO()
id := cloudpc.NewUserIdCloudPCID("userId", "cloudPCId")

payload := cloudpc.ResizeCloudPCRequest{
	// ...
}


read, err := client.ResizeCloudPC(ctx, id, payload, cloudpc.DefaultResizeCloudPCOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.RestoreCloudPC`

```go
ctx := context.TODO()
id := cloudpc.NewUserIdCloudPCID("userId", "cloudPCId")

payload := cloudpc.RestoreCloudPCRequest{
	// ...
}


read, err := client.RestoreCloudPC(ctx, id, payload, cloudpc.DefaultRestoreCloudPCOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.RetryCloudPCPartnerAgentInstallation`

```go
ctx := context.TODO()
id := cloudpc.NewUserIdCloudPCID("userId", "cloudPCId")

read, err := client.RetryCloudPCPartnerAgentInstallation(ctx, id, cloudpc.DefaultRetryCloudPCPartnerAgentInstallationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.SetCloudPCReviewStatus`

```go
ctx := context.TODO()
id := cloudpc.NewUserIdCloudPCID("userId", "cloudPCId")

payload := cloudpc.SetCloudPCReviewStatusRequest{
	// ...
}


read, err := client.SetCloudPCReviewStatus(ctx, id, payload, cloudpc.DefaultSetCloudPCReviewStatusOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.StartCloudPC`

```go
ctx := context.TODO()
id := cloudpc.NewUserIdCloudPCID("userId", "cloudPCId")

read, err := client.StartCloudPC(ctx, id, cloudpc.DefaultStartCloudPCOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.StopCloudPC`

```go
ctx := context.TODO()
id := cloudpc.NewUserIdCloudPCID("userId", "cloudPCId")

read, err := client.StopCloudPC(ctx, id, cloudpc.DefaultStopCloudPCOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.TroubleshootCloudPC`

```go
ctx := context.TODO()
id := cloudpc.NewUserIdCloudPCID("userId", "cloudPCId")

read, err := client.TroubleshootCloudPC(ctx, id, cloudpc.DefaultTroubleshootCloudPCOperationOptions())
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
id := cloudpc.NewUserIdCloudPCID("userId", "cloudPCId")

payload := cloudpc.CloudPC{
	// ...
}


read, err := client.UpdateCloudPC(ctx, id, payload, cloudpc.DefaultUpdateCloudPCOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.ValidateCloudPCsBulkResizes`

```go
ctx := context.TODO()
id := cloudpc.NewUserID("userId")

payload := cloudpc.ValidateCloudPCsBulkResizesRequest{
	// ...
}


// alternatively `client.ValidateCloudPCsBulkResizes(ctx, id, payload, cloudpc.DefaultValidateCloudPCsBulkResizesOperationOptions())` can be used to do batched pagination
items, err := client.ValidateCloudPCsBulkResizesComplete(ctx, id, payload, cloudpc.DefaultValidateCloudPCsBulkResizesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
