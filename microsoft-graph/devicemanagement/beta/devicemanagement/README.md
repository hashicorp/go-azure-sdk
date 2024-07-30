
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicemanagement` Documentation

The `devicemanagement` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicemanagement"
```


### Client Initialization

```go
client := devicemanagement.NewDeviceManagementClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeviceManagementClient.CreateEnableAndroidDeviceAdministratorEnrollment`

```go
ctx := context.TODO()


read, err := client.CreateEnableAndroidDeviceAdministratorEnrollment(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementClient.CreateEnableLegacyPcManagement`

```go
ctx := context.TODO()


read, err := client.CreateEnableLegacyPcManagement(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementClient.CreateEnableUnlicensedAdminstrator`

```go
ctx := context.TODO()


read, err := client.CreateEnableUnlicensedAdminstrator(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementClient.CreateEvaluateAssignmentFilter`

```go
ctx := context.TODO()

payload := devicemanagement.CreateEvaluateAssignmentFilterRequest{
	// ...
}


read, err := client.CreateEvaluateAssignmentFilter(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementClient.CreateSendCustomNotificationToCompanyPortal`

```go
ctx := context.TODO()

payload := devicemanagement.CreateSendCustomNotificationToCompanyPortalRequest{
	// ...
}


read, err := client.CreateSendCustomNotificationToCompanyPortal(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementClient.GetDeviceManagement`

```go
ctx := context.TODO()


read, err := client.GetDeviceManagement(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementClient.GetDeviceManagementAssignmentFiltersStatusDetail`

```go
ctx := context.TODO()

payload := devicemanagement.GetDeviceManagementAssignmentFiltersStatusDetailRequest{
	// ...
}


read, err := client.GetDeviceManagementAssignmentFiltersStatusDetail(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementClient.UpdateDeviceManagement`

```go
ctx := context.TODO()

payload := devicemanagement.DeviceManagement{
	// ...
}


read, err := client.UpdateDeviceManagement(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
