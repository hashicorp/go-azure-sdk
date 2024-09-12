
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


### Example Usage: `DeviceManagementClient.EnableAndroidDeviceAdministratorEnrollment`

```go
ctx := context.TODO()


read, err := client.EnableAndroidDeviceAdministratorEnrollment(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementClient.EnableLegacyPcManagement`

```go
ctx := context.TODO()


read, err := client.EnableLegacyPcManagement(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementClient.EnableUnlicensedAdminstrator`

```go
ctx := context.TODO()


read, err := client.EnableUnlicensedAdminstrator(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementClient.GetAssignmentFiltersStatusDetail`

```go
ctx := context.TODO()

payload := devicemanagement.GetAssignmentFiltersStatusDetailRequest{
	// ...
}


read, err := client.GetAssignmentFiltersStatusDetail(ctx, payload)
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


read, err := client.GetDeviceManagement(ctx, devicemanagement.DefaultGetDeviceManagementOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementClient.SendCustomNotificationToCompanyPortal`

```go
ctx := context.TODO()

payload := devicemanagement.SendCustomNotificationToCompanyPortalRequest{
	// ...
}


read, err := client.SendCustomNotificationToCompanyPortal(ctx, payload)
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
