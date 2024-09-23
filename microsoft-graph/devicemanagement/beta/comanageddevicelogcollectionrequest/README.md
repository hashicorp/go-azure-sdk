
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/comanageddevicelogcollectionrequest` Documentation

The `comanageddevicelogcollectionrequest` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/comanageddevicelogcollectionrequest"
```


### Client Initialization

```go
client := comanageddevicelogcollectionrequest.NewComanagedDeviceLogCollectionRequestClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ComanagedDeviceLogCollectionRequestClient.CreateComanagedDeviceLogCollectionRequest`

```go
ctx := context.TODO()
id := comanageddevicelogcollectionrequest.NewDeviceManagementComanagedDeviceID("managedDeviceId")

payload := comanageddevicelogcollectionrequest.DeviceLogCollectionResponse{
	// ...
}


read, err := client.CreateComanagedDeviceLogCollectionRequest(ctx, id, payload, comanageddevicelogcollectionrequest.DefaultCreateComanagedDeviceLogCollectionRequestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceLogCollectionRequestClient.CreateComanagedDeviceLogCollectionRequestDownloadUrl`

```go
ctx := context.TODO()
id := comanageddevicelogcollectionrequest.NewDeviceManagementComanagedDeviceIdLogCollectionRequestID("managedDeviceId", "deviceLogCollectionResponseId")

read, err := client.CreateComanagedDeviceLogCollectionRequestDownloadUrl(ctx, id, comanageddevicelogcollectionrequest.DefaultCreateComanagedDeviceLogCollectionRequestDownloadUrlOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceLogCollectionRequestClient.DeleteComanagedDeviceLogCollectionRequest`

```go
ctx := context.TODO()
id := comanageddevicelogcollectionrequest.NewDeviceManagementComanagedDeviceIdLogCollectionRequestID("managedDeviceId", "deviceLogCollectionResponseId")

read, err := client.DeleteComanagedDeviceLogCollectionRequest(ctx, id, comanageddevicelogcollectionrequest.DefaultDeleteComanagedDeviceLogCollectionRequestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceLogCollectionRequestClient.GetComanagedDeviceLogCollectionRequest`

```go
ctx := context.TODO()
id := comanageddevicelogcollectionrequest.NewDeviceManagementComanagedDeviceIdLogCollectionRequestID("managedDeviceId", "deviceLogCollectionResponseId")

read, err := client.GetComanagedDeviceLogCollectionRequest(ctx, id, comanageddevicelogcollectionrequest.DefaultGetComanagedDeviceLogCollectionRequestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceLogCollectionRequestClient.GetComanagedDeviceLogCollectionRequestsCount`

```go
ctx := context.TODO()
id := comanageddevicelogcollectionrequest.NewDeviceManagementComanagedDeviceID("managedDeviceId")

read, err := client.GetComanagedDeviceLogCollectionRequestsCount(ctx, id, comanageddevicelogcollectionrequest.DefaultGetComanagedDeviceLogCollectionRequestsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceLogCollectionRequestClient.ListComanagedDeviceLogCollectionRequests`

```go
ctx := context.TODO()
id := comanageddevicelogcollectionrequest.NewDeviceManagementComanagedDeviceID("managedDeviceId")

// alternatively `client.ListComanagedDeviceLogCollectionRequests(ctx, id, comanageddevicelogcollectionrequest.DefaultListComanagedDeviceLogCollectionRequestsOperationOptions())` can be used to do batched pagination
items, err := client.ListComanagedDeviceLogCollectionRequestsComplete(ctx, id, comanageddevicelogcollectionrequest.DefaultListComanagedDeviceLogCollectionRequestsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ComanagedDeviceLogCollectionRequestClient.UpdateComanagedDeviceLogCollectionRequest`

```go
ctx := context.TODO()
id := comanageddevicelogcollectionrequest.NewDeviceManagementComanagedDeviceIdLogCollectionRequestID("managedDeviceId", "deviceLogCollectionResponseId")

payload := comanageddevicelogcollectionrequest.DeviceLogCollectionResponse{
	// ...
}


read, err := client.UpdateComanagedDeviceLogCollectionRequest(ctx, id, payload, comanageddevicelogcollectionrequest.DefaultUpdateComanagedDeviceLogCollectionRequestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
