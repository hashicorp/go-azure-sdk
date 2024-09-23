
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/chromeosonboardingsetting` Documentation

The `chromeosonboardingsetting` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/chromeosonboardingsetting"
```


### Client Initialization

```go
client := chromeosonboardingsetting.NewChromeOSOnboardingSettingClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ChromeOSOnboardingSettingClient.CreateChromeOSOnboardingSetting`

```go
ctx := context.TODO()

payload := chromeosonboardingsetting.ChromeOSOnboardingSettings{
	// ...
}


read, err := client.CreateChromeOSOnboardingSetting(ctx, payload, chromeosonboardingsetting.DefaultCreateChromeOSOnboardingSettingOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChromeOSOnboardingSettingClient.CreateChromeOSOnboardingSettingConnect`

```go
ctx := context.TODO()

payload := chromeosonboardingsetting.CreateChromeOSOnboardingSettingConnectRequest{
	// ...
}


read, err := client.CreateChromeOSOnboardingSettingConnect(ctx, payload, chromeosonboardingsetting.DefaultCreateChromeOSOnboardingSettingConnectOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChromeOSOnboardingSettingClient.CreateChromeOSOnboardingSettingDisconnect`

```go
ctx := context.TODO()


read, err := client.CreateChromeOSOnboardingSettingDisconnect(ctx, chromeosonboardingsetting.DefaultCreateChromeOSOnboardingSettingDisconnectOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChromeOSOnboardingSettingClient.DeleteChromeOSOnboardingSetting`

```go
ctx := context.TODO()
id := chromeosonboardingsetting.NewDeviceManagementChromeOSOnboardingSettingID("chromeOSOnboardingSettingsId")

read, err := client.DeleteChromeOSOnboardingSetting(ctx, id, chromeosonboardingsetting.DefaultDeleteChromeOSOnboardingSettingOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChromeOSOnboardingSettingClient.GetChromeOSOnboardingSetting`

```go
ctx := context.TODO()
id := chromeosonboardingsetting.NewDeviceManagementChromeOSOnboardingSettingID("chromeOSOnboardingSettingsId")

read, err := client.GetChromeOSOnboardingSetting(ctx, id, chromeosonboardingsetting.DefaultGetChromeOSOnboardingSettingOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChromeOSOnboardingSettingClient.GetChromeOSOnboardingSettingsCount`

```go
ctx := context.TODO()


read, err := client.GetChromeOSOnboardingSettingsCount(ctx, chromeosonboardingsetting.DefaultGetChromeOSOnboardingSettingsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChromeOSOnboardingSettingClient.ListChromeOSOnboardingSettings`

```go
ctx := context.TODO()


// alternatively `client.ListChromeOSOnboardingSettings(ctx, chromeosonboardingsetting.DefaultListChromeOSOnboardingSettingsOperationOptions())` can be used to do batched pagination
items, err := client.ListChromeOSOnboardingSettingsComplete(ctx, chromeosonboardingsetting.DefaultListChromeOSOnboardingSettingsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ChromeOSOnboardingSettingClient.UpdateChromeOSOnboardingSetting`

```go
ctx := context.TODO()
id := chromeosonboardingsetting.NewDeviceManagementChromeOSOnboardingSettingID("chromeOSOnboardingSettingsId")

payload := chromeosonboardingsetting.ChromeOSOnboardingSettings{
	// ...
}


read, err := client.UpdateChromeOSOnboardingSetting(ctx, id, payload, chromeosonboardingsetting.DefaultUpdateChromeOSOnboardingSettingOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
