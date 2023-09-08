
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecloudpc` Documentation

The `mecloudpc` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecloudpc"
```


### Client Initialization

```go
client := mecloudpc.NewMeCloudPCClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCloudPCClient.CreateMeCloudPC`

```go
ctx := context.TODO()

payload := mecloudpc.CloudPC{
	// ...
}


read, err := client.CreateMeCloudPC(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCloudPCClient.CreateMeCloudPCBulkResize`

```go
ctx := context.TODO()

payload := mecloudpc.CreateMeCloudPCBulkResizeRequest{
	// ...
}


read, err := client.CreateMeCloudPCBulkResize(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCloudPCClient.CreateMeCloudPCByIdChangeUserAccountType`

```go
ctx := context.TODO()
id := mecloudpc.NewMeCloudPCID("cloudPCIdValue")

payload := mecloudpc.CreateMeCloudPCByIdChangeUserAccountTypeRequest{
	// ...
}


read, err := client.CreateMeCloudPCByIdChangeUserAccountType(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCloudPCClient.CreateMeCloudPCByIdEndGracePeriod`

```go
ctx := context.TODO()
id := mecloudpc.NewMeCloudPCID("cloudPCIdValue")

read, err := client.CreateMeCloudPCByIdEndGracePeriod(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCloudPCClient.CreateMeCloudPCByIdPowerOff`

```go
ctx := context.TODO()
id := mecloudpc.NewMeCloudPCID("cloudPCIdValue")

read, err := client.CreateMeCloudPCByIdPowerOff(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCloudPCClient.CreateMeCloudPCByIdPowerOn`

```go
ctx := context.TODO()
id := mecloudpc.NewMeCloudPCID("cloudPCIdValue")

read, err := client.CreateMeCloudPCByIdPowerOn(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCloudPCClient.CreateMeCloudPCByIdReboot`

```go
ctx := context.TODO()
id := mecloudpc.NewMeCloudPCID("cloudPCIdValue")

read, err := client.CreateMeCloudPCByIdReboot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCloudPCClient.CreateMeCloudPCByIdRename`

```go
ctx := context.TODO()
id := mecloudpc.NewMeCloudPCID("cloudPCIdValue")

payload := mecloudpc.CreateMeCloudPCByIdRenameRequest{
	// ...
}


read, err := client.CreateMeCloudPCByIdRename(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCloudPCClient.CreateMeCloudPCByIdReprovision`

```go
ctx := context.TODO()
id := mecloudpc.NewMeCloudPCID("cloudPCIdValue")

payload := mecloudpc.CreateMeCloudPCByIdReprovisionRequest{
	// ...
}


read, err := client.CreateMeCloudPCByIdReprovision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCloudPCClient.CreateMeCloudPCByIdRetryPartnerAgentInstallation`

```go
ctx := context.TODO()
id := mecloudpc.NewMeCloudPCID("cloudPCIdValue")

read, err := client.CreateMeCloudPCByIdRetryPartnerAgentInstallation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCloudPCClient.CreateMeCloudPCByIdTroubleshoot`

```go
ctx := context.TODO()
id := mecloudpc.NewMeCloudPCID("cloudPCIdValue")

read, err := client.CreateMeCloudPCByIdTroubleshoot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCloudPCClient.DeleteMeCloudPCById`

```go
ctx := context.TODO()
id := mecloudpc.NewMeCloudPCID("cloudPCIdValue")

read, err := client.DeleteMeCloudPCById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCloudPCClient.GetMeCloudPCById`

```go
ctx := context.TODO()
id := mecloudpc.NewMeCloudPCID("cloudPCIdValue")

read, err := client.GetMeCloudPCById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCloudPCClient.GetMeCloudPCCount`

```go
ctx := context.TODO()


read, err := client.GetMeCloudPCCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCloudPCClient.ListMeCloudPCs`

```go
ctx := context.TODO()


// alternatively `client.ListMeCloudPCs(ctx)` can be used to do batched pagination
items, err := client.ListMeCloudPCsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeCloudPCClient.RestoreMeCloudPCById`

```go
ctx := context.TODO()
id := mecloudpc.NewMeCloudPCID("cloudPCIdValue")

payload := mecloudpc.RestoreMeCloudPCByIdRequest{
	// ...
}


read, err := client.RestoreMeCloudPCById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCloudPCClient.StartMeCloudPCById`

```go
ctx := context.TODO()
id := mecloudpc.NewMeCloudPCID("cloudPCIdValue")

read, err := client.StartMeCloudPCById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCloudPCClient.StopMeCloudPCById`

```go
ctx := context.TODO()
id := mecloudpc.NewMeCloudPCID("cloudPCIdValue")

read, err := client.StopMeCloudPCById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCloudPCClient.UpdateMeCloudPCById`

```go
ctx := context.TODO()
id := mecloudpc.NewMeCloudPCID("cloudPCIdValue")

payload := mecloudpc.CloudPC{
	// ...
}


read, err := client.UpdateMeCloudPCById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCloudPCClient.ValidateMeCloudPCsBulkResize`

```go
ctx := context.TODO()

payload := mecloudpc.ValidateMeCloudPCsBulkResizeRequest{
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
