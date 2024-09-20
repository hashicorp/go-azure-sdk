
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/resourceaccessprofile` Documentation

The `resourceaccessprofile` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/resourceaccessprofile"
```


### Client Initialization

```go
client := resourceaccessprofile.NewResourceAccessProfileClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ResourceAccessProfileClient.AssignResourceAccessProfiles`

```go
ctx := context.TODO()
id := resourceaccessprofile.NewDeviceManagementResourceAccessProfileID("deviceManagementResourceAccessProfileBaseId")

payload := resourceaccessprofile.AssignResourceAccessProfilesRequest{
	// ...
}


// alternatively `client.AssignResourceAccessProfiles(ctx, id, payload, resourceaccessprofile.DefaultAssignResourceAccessProfilesOperationOptions())` can be used to do batched pagination
items, err := client.AssignResourceAccessProfilesComplete(ctx, id, payload, resourceaccessprofile.DefaultAssignResourceAccessProfilesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ResourceAccessProfileClient.CreateResourceAccessProfile`

```go
ctx := context.TODO()

payload := resourceaccessprofile.DeviceManagementResourceAccessProfileBase{
	// ...
}


read, err := client.CreateResourceAccessProfile(ctx, payload, resourceaccessprofile.DefaultCreateResourceAccessProfileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ResourceAccessProfileClient.DeleteResourceAccessProfile`

```go
ctx := context.TODO()
id := resourceaccessprofile.NewDeviceManagementResourceAccessProfileID("deviceManagementResourceAccessProfileBaseId")

read, err := client.DeleteResourceAccessProfile(ctx, id, resourceaccessprofile.DefaultDeleteResourceAccessProfileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ResourceAccessProfileClient.GetResourceAccessProfile`

```go
ctx := context.TODO()
id := resourceaccessprofile.NewDeviceManagementResourceAccessProfileID("deviceManagementResourceAccessProfileBaseId")

read, err := client.GetResourceAccessProfile(ctx, id, resourceaccessprofile.DefaultGetResourceAccessProfileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ResourceAccessProfileClient.GetResourceAccessProfilesCount`

```go
ctx := context.TODO()


read, err := client.GetResourceAccessProfilesCount(ctx, resourceaccessprofile.DefaultGetResourceAccessProfilesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ResourceAccessProfileClient.ListResourceAccessProfileQueryByPlatformTypes`

```go
ctx := context.TODO()

payload := resourceaccessprofile.ListResourceAccessProfileQueryByPlatformTypesRequest{
	// ...
}


// alternatively `client.ListResourceAccessProfileQueryByPlatformTypes(ctx, payload, resourceaccessprofile.DefaultListResourceAccessProfileQueryByPlatformTypesOperationOptions())` can be used to do batched pagination
items, err := client.ListResourceAccessProfileQueryByPlatformTypesComplete(ctx, payload, resourceaccessprofile.DefaultListResourceAccessProfileQueryByPlatformTypesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ResourceAccessProfileClient.ListResourceAccessProfiles`

```go
ctx := context.TODO()


// alternatively `client.ListResourceAccessProfiles(ctx, resourceaccessprofile.DefaultListResourceAccessProfilesOperationOptions())` can be used to do batched pagination
items, err := client.ListResourceAccessProfilesComplete(ctx, resourceaccessprofile.DefaultListResourceAccessProfilesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ResourceAccessProfileClient.UpdateResourceAccessProfile`

```go
ctx := context.TODO()
id := resourceaccessprofile.NewDeviceManagementResourceAccessProfileID("deviceManagementResourceAccessProfileBaseId")

payload := resourceaccessprofile.DeviceManagementResourceAccessProfileBase{
	// ...
}


read, err := client.UpdateResourceAccessProfile(ctx, id, payload, resourceaccessprofile.DefaultUpdateResourceAccessProfileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
