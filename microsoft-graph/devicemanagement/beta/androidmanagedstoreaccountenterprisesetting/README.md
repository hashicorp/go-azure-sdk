
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/androidmanagedstoreaccountenterprisesetting` Documentation

The `androidmanagedstoreaccountenterprisesetting` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/androidmanagedstoreaccountenterprisesetting"
```


### Client Initialization

```go
client := androidmanagedstoreaccountenterprisesetting.NewAndroidManagedStoreAccountEnterpriseSettingClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AndroidManagedStoreAccountEnterpriseSettingClient.AddAndroidManagedStoreAccountEnterpriseSettingsApp`

```go
ctx := context.TODO()

payload := androidmanagedstoreaccountenterprisesetting.AddAndroidManagedStoreAccountEnterpriseSettingsAppRequest{
	// ...
}


read, err := client.AddAndroidManagedStoreAccountEnterpriseSettingsApp(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidManagedStoreAccountEnterpriseSettingClient.CreateAndroidManagedStoreAccountEnterpriseSettingApproveApp`

```go
ctx := context.TODO()

payload := androidmanagedstoreaccountenterprisesetting.CreateAndroidManagedStoreAccountEnterpriseSettingApproveAppRequest{
	// ...
}


read, err := client.CreateAndroidManagedStoreAccountEnterpriseSettingApproveApp(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidManagedStoreAccountEnterpriseSettingClient.CreateAndroidManagedStoreAccountEnterpriseSettingCompleteSignup`

```go
ctx := context.TODO()

payload := androidmanagedstoreaccountenterprisesetting.CreateAndroidManagedStoreAccountEnterpriseSettingCompleteSignupRequest{
	// ...
}


read, err := client.CreateAndroidManagedStoreAccountEnterpriseSettingCompleteSignup(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidManagedStoreAccountEnterpriseSettingClient.CreateAndroidManagedStoreAccountEnterpriseSettingUnbind`

```go
ctx := context.TODO()


read, err := client.CreateAndroidManagedStoreAccountEnterpriseSettingUnbind(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidManagedStoreAccountEnterpriseSettingClient.CreateAndroidManagedStoreAccountEnterpriseSettingsGooglePlayWebToken`

```go
ctx := context.TODO()

payload := androidmanagedstoreaccountenterprisesetting.CreateAndroidManagedStoreAccountEnterpriseSettingsGooglePlayWebTokenRequest{
	// ...
}


read, err := client.CreateAndroidManagedStoreAccountEnterpriseSettingsGooglePlayWebToken(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidManagedStoreAccountEnterpriseSettingClient.DeleteAndroidManagedStoreAccountEnterpriseSetting`

```go
ctx := context.TODO()


read, err := client.DeleteAndroidManagedStoreAccountEnterpriseSetting(ctx, androidmanagedstoreaccountenterprisesetting.DefaultDeleteAndroidManagedStoreAccountEnterpriseSettingOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidManagedStoreAccountEnterpriseSettingClient.GetAndroidManagedStoreAccountEnterpriseSetting`

```go
ctx := context.TODO()


read, err := client.GetAndroidManagedStoreAccountEnterpriseSetting(ctx, androidmanagedstoreaccountenterprisesetting.DefaultGetAndroidManagedStoreAccountEnterpriseSettingOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidManagedStoreAccountEnterpriseSettingClient.RequestAndroidManagedStoreAccountEnterpriseSettingsSignupUrl`

```go
ctx := context.TODO()

payload := androidmanagedstoreaccountenterprisesetting.RequestAndroidManagedStoreAccountEnterpriseSettingsSignupUrlRequest{
	// ...
}


read, err := client.RequestAndroidManagedStoreAccountEnterpriseSettingsSignupUrl(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidManagedStoreAccountEnterpriseSettingClient.SetAndroidManagedStoreAccountEnterpriseSettingsAndroidDeviceOwnerFullyManagedEnrollmentState`

```go
ctx := context.TODO()

payload := androidmanagedstoreaccountenterprisesetting.SetAndroidManagedStoreAccountEnterpriseSettingsAndroidDeviceOwnerFullyManagedEnrollmentStateRequest{
	// ...
}


read, err := client.SetAndroidManagedStoreAccountEnterpriseSettingsAndroidDeviceOwnerFullyManagedEnrollmentState(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidManagedStoreAccountEnterpriseSettingClient.SyncAndroidManagedStoreAccountEnterpriseSettingsApp`

```go
ctx := context.TODO()


read, err := client.SyncAndroidManagedStoreAccountEnterpriseSettingsApp(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidManagedStoreAccountEnterpriseSettingClient.UpdateAndroidManagedStoreAccountEnterpriseSetting`

```go
ctx := context.TODO()

payload := androidmanagedstoreaccountenterprisesetting.AndroidManagedStoreAccountEnterpriseSettings{
	// ...
}


read, err := client.UpdateAndroidManagedStoreAccountEnterpriseSetting(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
