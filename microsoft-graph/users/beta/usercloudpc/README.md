
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercloudpc` Documentation

The `usercloudpc` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercloudpc"
```


### Client Initialization

```go
client := usercloudpc.NewUserCloudPCClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCloudPCClient.CreateUserByIdCloudPC`

```go
ctx := context.TODO()
id := usercloudpc.NewUserID("userIdValue")

payload := usercloudpc.CloudPC{
	// ...
}


read, err := client.CreateUserByIdCloudPC(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCloudPCClient.CreateUserByIdCloudPCBulkResize`

```go
ctx := context.TODO()
id := usercloudpc.NewUserID("userIdValue")

payload := usercloudpc.CreateUserByIdCloudPCBulkResizeRequest{
	// ...
}


read, err := client.CreateUserByIdCloudPCBulkResize(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCloudPCClient.CreateUserByIdCloudPCByIdChangeUserAccountType`

```go
ctx := context.TODO()
id := usercloudpc.NewUserCloudPCID("userIdValue", "cloudPCIdValue")

payload := usercloudpc.CreateUserByIdCloudPCByIdChangeUserAccountTypeRequest{
	// ...
}


read, err := client.CreateUserByIdCloudPCByIdChangeUserAccountType(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCloudPCClient.CreateUserByIdCloudPCByIdEndGracePeriod`

```go
ctx := context.TODO()
id := usercloudpc.NewUserCloudPCID("userIdValue", "cloudPCIdValue")

read, err := client.CreateUserByIdCloudPCByIdEndGracePeriod(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCloudPCClient.CreateUserByIdCloudPCByIdPowerOff`

```go
ctx := context.TODO()
id := usercloudpc.NewUserCloudPCID("userIdValue", "cloudPCIdValue")

read, err := client.CreateUserByIdCloudPCByIdPowerOff(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCloudPCClient.CreateUserByIdCloudPCByIdPowerOn`

```go
ctx := context.TODO()
id := usercloudpc.NewUserCloudPCID("userIdValue", "cloudPCIdValue")

read, err := client.CreateUserByIdCloudPCByIdPowerOn(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCloudPCClient.CreateUserByIdCloudPCByIdReboot`

```go
ctx := context.TODO()
id := usercloudpc.NewUserCloudPCID("userIdValue", "cloudPCIdValue")

read, err := client.CreateUserByIdCloudPCByIdReboot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCloudPCClient.CreateUserByIdCloudPCByIdRename`

```go
ctx := context.TODO()
id := usercloudpc.NewUserCloudPCID("userIdValue", "cloudPCIdValue")

payload := usercloudpc.CreateUserByIdCloudPCByIdRenameRequest{
	// ...
}


read, err := client.CreateUserByIdCloudPCByIdRename(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCloudPCClient.CreateUserByIdCloudPCByIdReprovision`

```go
ctx := context.TODO()
id := usercloudpc.NewUserCloudPCID("userIdValue", "cloudPCIdValue")

payload := usercloudpc.CreateUserByIdCloudPCByIdReprovisionRequest{
	// ...
}


read, err := client.CreateUserByIdCloudPCByIdReprovision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCloudPCClient.CreateUserByIdCloudPCByIdRetryPartnerAgentInstallation`

```go
ctx := context.TODO()
id := usercloudpc.NewUserCloudPCID("userIdValue", "cloudPCIdValue")

read, err := client.CreateUserByIdCloudPCByIdRetryPartnerAgentInstallation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCloudPCClient.CreateUserByIdCloudPCByIdTroubleshoot`

```go
ctx := context.TODO()
id := usercloudpc.NewUserCloudPCID("userIdValue", "cloudPCIdValue")

read, err := client.CreateUserByIdCloudPCByIdTroubleshoot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCloudPCClient.DeleteUserByIdCloudPCById`

```go
ctx := context.TODO()
id := usercloudpc.NewUserCloudPCID("userIdValue", "cloudPCIdValue")

read, err := client.DeleteUserByIdCloudPCById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCloudPCClient.GetUserByIdCloudPCById`

```go
ctx := context.TODO()
id := usercloudpc.NewUserCloudPCID("userIdValue", "cloudPCIdValue")

read, err := client.GetUserByIdCloudPCById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCloudPCClient.GetUserByIdCloudPCCount`

```go
ctx := context.TODO()
id := usercloudpc.NewUserID("userIdValue")

read, err := client.GetUserByIdCloudPCCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCloudPCClient.ListUserByIdCloudPCs`

```go
ctx := context.TODO()
id := usercloudpc.NewUserID("userIdValue")

// alternatively `client.ListUserByIdCloudPCs(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCloudPCsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserCloudPCClient.RestoreUserByIdCloudPCById`

```go
ctx := context.TODO()
id := usercloudpc.NewUserCloudPCID("userIdValue", "cloudPCIdValue")

payload := usercloudpc.RestoreUserByIdCloudPCByIdRequest{
	// ...
}


read, err := client.RestoreUserByIdCloudPCById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCloudPCClient.StartUserByIdCloudPCById`

```go
ctx := context.TODO()
id := usercloudpc.NewUserCloudPCID("userIdValue", "cloudPCIdValue")

read, err := client.StartUserByIdCloudPCById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCloudPCClient.StopUserByIdCloudPCById`

```go
ctx := context.TODO()
id := usercloudpc.NewUserCloudPCID("userIdValue", "cloudPCIdValue")

read, err := client.StopUserByIdCloudPCById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCloudPCClient.UpdateUserByIdCloudPCById`

```go
ctx := context.TODO()
id := usercloudpc.NewUserCloudPCID("userIdValue", "cloudPCIdValue")

payload := usercloudpc.CloudPC{
	// ...
}


read, err := client.UpdateUserByIdCloudPCById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCloudPCClient.ValidateUserByIdCloudPCsBulkResize`

```go
ctx := context.TODO()
id := usercloudpc.NewUserID("userIdValue")

payload := usercloudpc.ValidateUserByIdCloudPCsBulkResizeRequest{
	// ...
}


read, err := client.ValidateUserByIdCloudPCsBulkResize(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
