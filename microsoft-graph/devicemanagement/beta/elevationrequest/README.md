
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/elevationrequest` Documentation

The `elevationrequest` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/elevationrequest"
```


### Client Initialization

```go
client := elevationrequest.NewElevationRequestClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ElevationRequestClient.CreateElevationRequest`

```go
ctx := context.TODO()

payload := elevationrequest.PrivilegeManagementElevationRequest{
	// ...
}


read, err := client.CreateElevationRequest(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ElevationRequestClient.CreateElevationRequestApprove`

```go
ctx := context.TODO()
id := elevationrequest.NewDeviceManagementElevationRequestID("privilegeManagementElevationRequestIdValue")

payload := elevationrequest.CreateElevationRequestApproveRequest{
	// ...
}


read, err := client.CreateElevationRequestApprove(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ElevationRequestClient.CreateElevationRequestDeny`

```go
ctx := context.TODO()
id := elevationrequest.NewDeviceManagementElevationRequestID("privilegeManagementElevationRequestIdValue")

payload := elevationrequest.CreateElevationRequestDenyRequest{
	// ...
}


read, err := client.CreateElevationRequestDeny(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ElevationRequestClient.DeleteElevationRequest`

```go
ctx := context.TODO()
id := elevationrequest.NewDeviceManagementElevationRequestID("privilegeManagementElevationRequestIdValue")

read, err := client.DeleteElevationRequest(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ElevationRequestClient.GetDeviceManagementElevationRequestAllElevationRequests`

```go
ctx := context.TODO()
id := elevationrequest.NewDeviceManagementElevationRequestID("privilegeManagementElevationRequestIdValue")

// alternatively `client.GetDeviceManagementElevationRequestAllElevationRequests(ctx, id)` can be used to do batched pagination
items, err := client.GetDeviceManagementElevationRequestAllElevationRequestsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ElevationRequestClient.GetElevationRequest`

```go
ctx := context.TODO()
id := elevationrequest.NewDeviceManagementElevationRequestID("privilegeManagementElevationRequestIdValue")

read, err := client.GetElevationRequest(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ElevationRequestClient.GetElevationRequestCount`

```go
ctx := context.TODO()


read, err := client.GetElevationRequestCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ElevationRequestClient.ListElevationRequests`

```go
ctx := context.TODO()


// alternatively `client.ListElevationRequests(ctx)` can be used to do batched pagination
items, err := client.ListElevationRequestsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ElevationRequestClient.UpdateElevationRequest`

```go
ctx := context.TODO()
id := elevationrequest.NewDeviceManagementElevationRequestID("privilegeManagementElevationRequestIdValue")

payload := elevationrequest.PrivilegeManagementElevationRequest{
	// ...
}


read, err := client.UpdateElevationRequest(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
