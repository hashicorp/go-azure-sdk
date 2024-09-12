
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsautopilotdeploymentprofile` Documentation

The `windowsautopilotdeploymentprofile` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsautopilotdeploymentprofile"
```


### Client Initialization

```go
client := windowsautopilotdeploymentprofile.NewWindowsAutopilotDeploymentProfileClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WindowsAutopilotDeploymentProfileClient.AssignWindowsAutopilotDeploymentProfile`

```go
ctx := context.TODO()
id := windowsautopilotdeploymentprofile.NewDeviceManagementWindowsAutopilotDeploymentProfileID("windowsAutopilotDeploymentProfileIdValue")

payload := windowsautopilotdeploymentprofile.AssignWindowsAutopilotDeploymentProfileRequest{
	// ...
}


read, err := client.AssignWindowsAutopilotDeploymentProfile(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsAutopilotDeploymentProfileClient.CreateWindowsAutopilotDeploymentProfile`

```go
ctx := context.TODO()

payload := windowsautopilotdeploymentprofile.WindowsAutopilotDeploymentProfile{
	// ...
}


read, err := client.CreateWindowsAutopilotDeploymentProfile(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsAutopilotDeploymentProfileClient.DeleteWindowsAutopilotDeploymentProfile`

```go
ctx := context.TODO()
id := windowsautopilotdeploymentprofile.NewDeviceManagementWindowsAutopilotDeploymentProfileID("windowsAutopilotDeploymentProfileIdValue")

read, err := client.DeleteWindowsAutopilotDeploymentProfile(ctx, id, windowsautopilotdeploymentprofile.DefaultDeleteWindowsAutopilotDeploymentProfileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsAutopilotDeploymentProfileClient.GetWindowsAutopilotDeploymentProfile`

```go
ctx := context.TODO()
id := windowsautopilotdeploymentprofile.NewDeviceManagementWindowsAutopilotDeploymentProfileID("windowsAutopilotDeploymentProfileIdValue")

read, err := client.GetWindowsAutopilotDeploymentProfile(ctx, id, windowsautopilotdeploymentprofile.DefaultGetWindowsAutopilotDeploymentProfileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsAutopilotDeploymentProfileClient.GetWindowsAutopilotDeploymentProfilesCount`

```go
ctx := context.TODO()


read, err := client.GetWindowsAutopilotDeploymentProfilesCount(ctx, windowsautopilotdeploymentprofile.DefaultGetWindowsAutopilotDeploymentProfilesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsAutopilotDeploymentProfileClient.ListWindowsAutopilotDeploymentProfileHasPayloadLinks`

```go
ctx := context.TODO()

payload := windowsautopilotdeploymentprofile.ListWindowsAutopilotDeploymentProfileHasPayloadLinksRequest{
	// ...
}


// alternatively `client.ListWindowsAutopilotDeploymentProfileHasPayloadLinks(ctx, payload, windowsautopilotdeploymentprofile.DefaultListWindowsAutopilotDeploymentProfileHasPayloadLinksOperationOptions())` can be used to do batched pagination
items, err := client.ListWindowsAutopilotDeploymentProfileHasPayloadLinksComplete(ctx, payload, windowsautopilotdeploymentprofile.DefaultListWindowsAutopilotDeploymentProfileHasPayloadLinksOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WindowsAutopilotDeploymentProfileClient.ListWindowsAutopilotDeploymentProfiles`

```go
ctx := context.TODO()


// alternatively `client.ListWindowsAutopilotDeploymentProfiles(ctx, windowsautopilotdeploymentprofile.DefaultListWindowsAutopilotDeploymentProfilesOperationOptions())` can be used to do batched pagination
items, err := client.ListWindowsAutopilotDeploymentProfilesComplete(ctx, windowsautopilotdeploymentprofile.DefaultListWindowsAutopilotDeploymentProfilesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WindowsAutopilotDeploymentProfileClient.UpdateWindowsAutopilotDeploymentProfile`

```go
ctx := context.TODO()
id := windowsautopilotdeploymentprofile.NewDeviceManagementWindowsAutopilotDeploymentProfileID("windowsAutopilotDeploymentProfileIdValue")

payload := windowsautopilotdeploymentprofile.WindowsAutopilotDeploymentProfile{
	// ...
}


read, err := client.UpdateWindowsAutopilotDeploymentProfile(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
