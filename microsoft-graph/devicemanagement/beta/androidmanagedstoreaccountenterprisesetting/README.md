
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/androidmanagedstoreaccountenterprisesetting` Documentation

The `androidmanagedstoreaccountenterprisesetting` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/androidmanagedstoreaccountenterprisesetting"
```


### Client Initialization

```go
client := androidmanagedstoreaccountenterprisesetting.NewAndroidManagedStoreAccountEnterpriseSettingClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AndroidManagedStoreAccountEnterpriseSettingClient.AddAndroidManagedStoreAccountEnterpriseSettingsApps`

```go
ctx := context.TODO()

payload := androidmanagedstoreaccountenterprisesetting.AddAndroidManagedStoreAccountEnterpriseSettingsAppsRequest{
	// ...
}


read, err := client.AddAndroidManagedStoreAccountEnterpriseSettingsApps(ctx, payload, androidmanagedstoreaccountenterprisesetting.DefaultAddAndroidManagedStoreAccountEnterpriseSettingsAppsOperationOptions())
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


read, err := client.CreateAndroidManagedStoreAccountEnterpriseSettingApproveApp(ctx, payload, androidmanagedstoreaccountenterprisesetting.DefaultCreateAndroidManagedStoreAccountEnterpriseSettingApproveAppOperationOptions())
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


read, err := client.CreateAndroidManagedStoreAccountEnterpriseSettingCompleteSignup(ctx, payload, androidmanagedstoreaccountenterprisesetting.DefaultCreateAndroidManagedStoreAccountEnterpriseSettingCompleteSignupOperationOptions())
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


read, err := client.CreateAndroidManagedStoreAccountEnterpriseSettingUnbind(ctx, androidmanagedstoreaccountenterprisesetting.DefaultCreateAndroidManagedStoreAccountEnterpriseSettingUnbindOperationOptions())
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


read, err := client.CreateAndroidManagedStoreAccountEnterpriseSettingsGooglePlayWebToken(ctx, payload, androidmanagedstoreaccountenterprisesetting.DefaultCreateAndroidManagedStoreAccountEnterpriseSettingsGooglePlayWebTokenOperationOptions())
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


read, err := client.RequestAndroidManagedStoreAccountEnterpriseSettingsSignupUrl(ctx, payload, androidmanagedstoreaccountenterprisesetting.DefaultRequestAndroidManagedStoreAccountEnterpriseSettingsSignupUrlOperationOptions())
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


read, err := client.SetAndroidManagedStoreAccountEnterpriseSettingsAndroidDeviceOwnerFullyManagedEnrollmentState(ctx, payload, androidmanagedstoreaccountenterprisesetting.DefaultSetAndroidManagedStoreAccountEnterpriseSettingsAndroidDeviceOwnerFullyManagedEnrollmentStateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidManagedStoreAccountEnterpriseSettingClient.SyncAndroidManagedStoreAccountEnterpriseSettingsApps`

```go
ctx := context.TODO()


read, err := client.SyncAndroidManagedStoreAccountEnterpriseSettingsApps(ctx, androidmanagedstoreaccountenterprisesetting.DefaultSyncAndroidManagedStoreAccountEnterpriseSettingsAppsOperationOptions())
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


read, err := client.UpdateAndroidManagedStoreAccountEnterpriseSetting(ctx, payload, androidmanagedstoreaccountenterprisesetting.DefaultUpdateAndroidManagedStoreAccountEnterpriseSettingOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
