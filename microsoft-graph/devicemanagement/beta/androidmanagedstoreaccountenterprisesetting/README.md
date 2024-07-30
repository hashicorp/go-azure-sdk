
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


### Example Usage: `AndroidManagedStoreAccountEnterpriseSettingClient.AddDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsApp`

```go
ctx := context.TODO()

payload := androidmanagedstoreaccountenterprisesetting.AddDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsAppRequest{
	// ...
}


read, err := client.AddDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsApp(ctx, payload)
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


### Example Usage: `AndroidManagedStoreAccountEnterpriseSettingClient.CreateAndroidManagedStoreAccountEnterpriseSettingCreateGooglePlayWebToken`

```go
ctx := context.TODO()

payload := androidmanagedstoreaccountenterprisesetting.CreateAndroidManagedStoreAccountEnterpriseSettingCreateGooglePlayWebTokenRequest{
	// ...
}


read, err := client.CreateAndroidManagedStoreAccountEnterpriseSettingCreateGooglePlayWebToken(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidManagedStoreAccountEnterpriseSettingClient.CreateAndroidManagedStoreAccountEnterpriseSettingRequestSignupUrl`

```go
ctx := context.TODO()

payload := androidmanagedstoreaccountenterprisesetting.CreateAndroidManagedStoreAccountEnterpriseSettingRequestSignupUrlRequest{
	// ...
}


read, err := client.CreateAndroidManagedStoreAccountEnterpriseSettingRequestSignupUrl(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidManagedStoreAccountEnterpriseSettingClient.CreateAndroidManagedStoreAccountEnterpriseSettingSyncApp`

```go
ctx := context.TODO()


read, err := client.CreateAndroidManagedStoreAccountEnterpriseSettingSyncApp(ctx)
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


### Example Usage: `AndroidManagedStoreAccountEnterpriseSettingClient.DeleteAndroidManagedStoreAccountEnterpriseSetting`

```go
ctx := context.TODO()


read, err := client.DeleteAndroidManagedStoreAccountEnterpriseSetting(ctx)
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


read, err := client.GetAndroidManagedStoreAccountEnterpriseSetting(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidManagedStoreAccountEnterpriseSettingClient.SetDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsAndroidDeviceOwnerFullyManagedEnrollmentState`

```go
ctx := context.TODO()

payload := androidmanagedstoreaccountenterprisesetting.SetDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsAndroidDeviceOwnerFullyManagedEnrollmentStateRequest{
	// ...
}


read, err := client.SetDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsAndroidDeviceOwnerFullyManagedEnrollmentState(ctx, payload)
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
