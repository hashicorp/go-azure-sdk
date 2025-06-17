
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicemanagement` Documentation

The `devicemanagement` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicemanagement"
```


### Client Initialization

```go
client := devicemanagement.NewDeviceManagementClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeviceManagementClient.EnableAndroidDeviceAdministratorEnrollment`

```go
ctx := context.TODO()


read, err := client.EnableAndroidDeviceAdministratorEnrollment(ctx, devicemanagement.DefaultEnableAndroidDeviceAdministratorEnrollmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementClient.EnableEndpointPrivilegeManagement`

```go
ctx := context.TODO()


read, err := client.EnableEndpointPrivilegeManagement(ctx, devicemanagement.DefaultEnableEndpointPrivilegeManagementOperationOptions())
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


read, err := client.EnableLegacyPcManagement(ctx, devicemanagement.DefaultEnableLegacyPcManagementOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementClient.EnableUnlicensedAdminstrators`

```go
ctx := context.TODO()


read, err := client.EnableUnlicensedAdminstrators(ctx, devicemanagement.DefaultEnableUnlicensedAdminstratorsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementClient.EvaluateAssignmentFilter`

```go
ctx := context.TODO()

payload := devicemanagement.EvaluateAssignmentFilterRequest{
	// ...
}


read, err := client.EvaluateAssignmentFilter(ctx, payload, devicemanagement.DefaultEvaluateAssignmentFilterOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementClient.GetAssignmentFiltersStatusDetails`

```go
ctx := context.TODO()

payload := devicemanagement.GetAssignmentFiltersStatusDetailsRequest{
	// ...
}


read, err := client.GetAssignmentFiltersStatusDetails(ctx, payload, devicemanagement.DefaultGetAssignmentFiltersStatusDetailsOperationOptions())
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


read, err := client.SendCustomNotificationToCompanyPortal(ctx, payload, devicemanagement.DefaultSendCustomNotificationToCompanyPortalOperationOptions())
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


read, err := client.UpdateDeviceManagement(ctx, payload, devicemanagement.DefaultUpdateDeviceManagementOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
