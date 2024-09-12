
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/me` Documentation

The `me` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/me"
```


### Client Initialization

```go
client := me.NewMeClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeClient.AssignLicense`

```go
ctx := context.TODO()

payload := me.AssignLicenseRequest{
	// ...
}


read, err := client.AssignLicense(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.ChangePassword`

```go
ctx := context.TODO()

payload := me.ChangePasswordRequest{
	// ...
}


read, err := client.ChangePassword(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CreateFindetingTime`

```go
ctx := context.TODO()

payload := me.CreateFindetingTimeRequest{
	// ...
}


read, err := client.CreateFindetingTime(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.ExportPersonalData`

```go
ctx := context.TODO()

payload := me.ExportPersonalDataRequest{
	// ...
}


read, err := client.ExportPersonalData(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.GetMailTips`

```go
ctx := context.TODO()

payload := me.GetMailTipsRequest{
	// ...
}


// alternatively `client.GetMailTips(ctx, payload, me.DefaultGetMailTipsOperationOptions())` can be used to do batched pagination
items, err := client.GetMailTipsComplete(ctx, payload, me.DefaultGetMailTipsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeClient.GetMe`

```go
ctx := context.TODO()


read, err := client.GetMe(ctx, me.DefaultGetMeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.ListCheckmberGroups`

```go
ctx := context.TODO()

payload := me.ListCheckmberGroupsRequest{
	// ...
}


// alternatively `client.ListCheckmberGroups(ctx, payload, me.DefaultListCheckmberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.ListCheckmberGroupsComplete(ctx, payload, me.DefaultListCheckmberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeClient.ListCheckmberObjects`

```go
ctx := context.TODO()

payload := me.ListCheckmberObjectsRequest{
	// ...
}


// alternatively `client.ListCheckmberObjects(ctx, payload, me.DefaultListCheckmberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.ListCheckmberObjectsComplete(ctx, payload, me.DefaultListCheckmberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeClient.ListGetmberGroups`

```go
ctx := context.TODO()

payload := me.ListGetmberGroupsRequest{
	// ...
}


// alternatively `client.ListGetmberGroups(ctx, payload, me.DefaultListGetmberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.ListGetmberGroupsComplete(ctx, payload, me.DefaultListGetmberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeClient.ListGetmberObjects`

```go
ctx := context.TODO()

payload := me.ListGetmberObjectsRequest{
	// ...
}


// alternatively `client.ListGetmberObjects(ctx, payload, me.DefaultListGetmberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.ListGetmberObjectsComplete(ctx, payload, me.DefaultListGetmberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeClient.RemoveAllDevicesFromManagement`

```go
ctx := context.TODO()


read, err := client.RemoveAllDevicesFromManagement(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.ReprocessLicenseAssignment`

```go
ctx := context.TODO()


read, err := client.ReprocessLicenseAssignment(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.Restore`

```go
ctx := context.TODO()


read, err := client.Restore(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.RetryServiceProvisioning`

```go
ctx := context.TODO()


read, err := client.RetryServiceProvisioning(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.RevokeSignInSession`

```go
ctx := context.TODO()


read, err := client.RevokeSignInSession(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.SendMail`

```go
ctx := context.TODO()

payload := me.SendMailRequest{
	// ...
}


read, err := client.SendMail(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.TranslateExchangeIds`

```go
ctx := context.TODO()

payload := me.TranslateExchangeIdsRequest{
	// ...
}


// alternatively `client.TranslateExchangeIds(ctx, payload, me.DefaultTranslateExchangeIdsOperationOptions())` can be used to do batched pagination
items, err := client.TranslateExchangeIdsComplete(ctx, payload, me.DefaultTranslateExchangeIdsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeClient.UpdateMe`

```go
ctx := context.TODO()

payload := me.User{
	// ...
}


read, err := client.UpdateMe(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.WipeManagedAppRegistrationsByDeviceTag`

```go
ctx := context.TODO()

payload := me.WipeManagedAppRegistrationsByDeviceTagRequest{
	// ...
}


read, err := client.WipeManagedAppRegistrationsByDeviceTag(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
