
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/androidforworksetting` Documentation

The `androidforworksetting` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/androidforworksetting"
```


### Client Initialization

```go
client := androidforworksetting.NewAndroidForWorkSettingClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AndroidForWorkSettingClient.CreateAndroidForWorkSettingCompleteSignup`

```go
ctx := context.TODO()

payload := androidforworksetting.CreateAndroidForWorkSettingCompleteSignupRequest{
	// ...
}


read, err := client.CreateAndroidForWorkSettingCompleteSignup(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidForWorkSettingClient.CreateAndroidForWorkSettingUnbind`

```go
ctx := context.TODO()


read, err := client.CreateAndroidForWorkSettingUnbind(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidForWorkSettingClient.DeleteAndroidForWorkSetting`

```go
ctx := context.TODO()


read, err := client.DeleteAndroidForWorkSetting(ctx, androidforworksetting.DefaultDeleteAndroidForWorkSettingOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidForWorkSettingClient.GetAndroidForWorkSetting`

```go
ctx := context.TODO()


read, err := client.GetAndroidForWorkSetting(ctx, androidforworksetting.DefaultGetAndroidForWorkSettingOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidForWorkSettingClient.RequestAndroidForWorkSettingsSignupUrl`

```go
ctx := context.TODO()

payload := androidforworksetting.RequestAndroidForWorkSettingsSignupUrlRequest{
	// ...
}


read, err := client.RequestAndroidForWorkSettingsSignupUrl(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidForWorkSettingClient.SyncAndroidForWorkSettingsApp`

```go
ctx := context.TODO()


read, err := client.SyncAndroidForWorkSettingsApp(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidForWorkSettingClient.UpdateAndroidForWorkSetting`

```go
ctx := context.TODO()

payload := androidforworksetting.AndroidForWorkSettings{
	// ...
}


read, err := client.UpdateAndroidForWorkSetting(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
