
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/resourceaccessprofile` Documentation

The `resourceaccessprofile` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/resourceaccessprofile"
```


### Client Initialization

```go
client := resourceaccessprofile.NewResourceAccessProfileClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ResourceAccessProfileClient.AssignDeviceManagementResourceAccessProfiles`

```go
ctx := context.TODO()
id := resourceaccessprofile.NewDeviceManagementResourceAccessProfileID("deviceManagementResourceAccessProfileBaseIdValue")

payload := resourceaccessprofile.AssignDeviceManagementResourceAccessProfilesRequest{
	// ...
}


// alternatively `client.AssignDeviceManagementResourceAccessProfiles(ctx, id, payload)` can be used to do batched pagination
items, err := client.AssignDeviceManagementResourceAccessProfilesComplete(ctx, id, payload)
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


read, err := client.CreateResourceAccessProfile(ctx, payload)
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
id := resourceaccessprofile.NewDeviceManagementResourceAccessProfileID("deviceManagementResourceAccessProfileBaseIdValue")

read, err := client.DeleteResourceAccessProfile(ctx, id)
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
id := resourceaccessprofile.NewDeviceManagementResourceAccessProfileID("deviceManagementResourceAccessProfileBaseIdValue")

read, err := client.GetResourceAccessProfile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ResourceAccessProfileClient.GetResourceAccessProfileCount`

```go
ctx := context.TODO()


read, err := client.GetResourceAccessProfileCount(ctx)
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


// alternatively `client.ListResourceAccessProfileQueryByPlatformTypes(ctx, payload)` can be used to do batched pagination
items, err := client.ListResourceAccessProfileQueryByPlatformTypesComplete(ctx, payload)
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


// alternatively `client.ListResourceAccessProfiles(ctx)` can be used to do batched pagination
items, err := client.ListResourceAccessProfilesComplete(ctx)
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
id := resourceaccessprofile.NewDeviceManagementResourceAccessProfileID("deviceManagementResourceAccessProfileBaseIdValue")

payload := resourceaccessprofile.DeviceManagementResourceAccessProfileBase{
	// ...
}


read, err := client.UpdateResourceAccessProfile(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
