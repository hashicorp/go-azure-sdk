
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/chromeosonboardingsetting` Documentation

The `chromeosonboardingsetting` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/chromeosonboardingsetting"
```


### Client Initialization

```go
client := chromeosonboardingsetting.NewChromeOSOnboardingSettingClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ChromeOSOnboardingSettingClient.CreateChromeOSOnboardingSetting`

```go
ctx := context.TODO()

payload := chromeosonboardingsetting.ChromeOSOnboardingSettings{
	// ...
}


read, err := client.CreateChromeOSOnboardingSetting(ctx, payload)
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


read, err := client.CreateChromeOSOnboardingSettingConnect(ctx, payload)
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


read, err := client.CreateChromeOSOnboardingSettingDisconnect(ctx)
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
id := chromeosonboardingsetting.NewDeviceManagementChromeOSOnboardingSettingID("chromeOSOnboardingSettingsIdValue")

read, err := client.DeleteChromeOSOnboardingSetting(ctx, id)
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
id := chromeosonboardingsetting.NewDeviceManagementChromeOSOnboardingSettingID("chromeOSOnboardingSettingsIdValue")

read, err := client.GetChromeOSOnboardingSetting(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChromeOSOnboardingSettingClient.GetChromeOSOnboardingSettingCount`

```go
ctx := context.TODO()


read, err := client.GetChromeOSOnboardingSettingCount(ctx)
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


// alternatively `client.ListChromeOSOnboardingSettings(ctx)` can be used to do batched pagination
items, err := client.ListChromeOSOnboardingSettingsComplete(ctx)
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
id := chromeosonboardingsetting.NewDeviceManagementChromeOSOnboardingSettingID("chromeOSOnboardingSettingsIdValue")

payload := chromeosonboardingsetting.ChromeOSOnboardingSettings{
	// ...
}


read, err := client.UpdateChromeOSOnboardingSetting(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
