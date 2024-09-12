
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


### Example Usage: `CloudPCClient.ChangeCloudPCUserAccountType`

```go
ctx := context.TODO()
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

payload := cloudpc.ChangeCloudPCUserAccountTypeRequest{
	// ...
}


read, err := client.ChangeCloudPCUserAccountType(ctx, id, payload)
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


### Example Usage: `CloudPCClient.CreateCloudPCSnapshot`

```go
ctx := context.TODO()
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

read, err := client.CreateCloudPCSnapshot(ctx, id)
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
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

read, err := client.EndCloudPCGracePeriod(ctx, id)
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


read, err := client.GetCloudPCsCount(ctx, cloudpc.DefaultGetCloudPCsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCClient.ListCloudPCBulkResizes`

```go
ctx := context.TODO()

payload := cloudpc.ListCloudPCBulkResizesRequest{
	// ...
}


// alternatively `client.ListCloudPCBulkResizes(ctx, payload, cloudpc.DefaultListCloudPCBulkResizesOperationOptions())` can be used to do batched pagination
items, err := client.ListCloudPCBulkResizesComplete(ctx, payload, cloudpc.DefaultListCloudPCBulkResizesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CloudPCClient.ListCloudPCS`

```go
ctx := context.TODO()


// alternatively `client.ListCloudPCS(ctx, cloudpc.DefaultListCloudPCSOperationOptions())` can be used to do batched pagination
items, err := client.ListCloudPCSComplete(ctx, cloudpc.DefaultListCloudPCSOperationOptions())
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
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

read, err := client.PowerOffCloudPC(ctx, id)
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
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

read, err := client.PowerOnCloudPC(ctx, id)
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
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

read, err := client.RebootCloudPC(ctx, id)
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
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

payload := cloudpc.RenameCloudPCRequest{
	// ...
}


read, err := client.RenameCloudPC(ctx, id, payload)
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
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

payload := cloudpc.ReprovisionCloudPCRequest{
	// ...
}


read, err := client.ReprovisionCloudPC(ctx, id, payload)
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
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

payload := cloudpc.ResizeCloudPCRequest{
	// ...
}


read, err := client.ResizeCloudPC(ctx, id, payload)
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
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

payload := cloudpc.RestoreCloudPCRequest{
	// ...
}


read, err := client.RestoreCloudPC(ctx, id, payload)
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
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

read, err := client.RetryCloudPCPartnerAgentInstallation(ctx, id)
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
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

payload := cloudpc.SetCloudPCReviewStatusRequest{
	// ...
}


read, err := client.SetCloudPCReviewStatus(ctx, id, payload)
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
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

read, err := client.StartCloudPC(ctx, id)
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
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

read, err := client.StopCloudPC(ctx, id)
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
id := cloudpc.NewMeCloudPCID("cloudPCIdValue")

read, err := client.TroubleshootCloudPC(ctx, id)
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


### Example Usage: `CloudPCClient.ValidateCloudPCsBulkResizes`

```go
ctx := context.TODO()

payload := cloudpc.ValidateCloudPCsBulkResizesRequest{
	// ...
}


// alternatively `client.ValidateCloudPCsBulkResizes(ctx, payload, cloudpc.DefaultValidateCloudPCsBulkResizesOperationOptions())` can be used to do batched pagination
items, err := client.ValidateCloudPCsBulkResizesComplete(ctx, payload, cloudpc.DefaultValidateCloudPCsBulkResizesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
