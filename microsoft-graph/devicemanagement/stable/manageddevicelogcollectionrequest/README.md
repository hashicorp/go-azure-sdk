
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/manageddevicelogcollectionrequest` Documentation

The `manageddevicelogcollectionrequest` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/manageddevicelogcollectionrequest"
```


### Client Initialization

```go
client := manageddevicelogcollectionrequest.NewManagedDeviceLogCollectionRequestClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedDeviceLogCollectionRequestClient.CreateManagedDeviceLogCollectionRequest`

```go
ctx := context.TODO()
id := manageddevicelogcollectionrequest.NewDeviceManagementManagedDeviceID("managedDeviceId")

payload := manageddevicelogcollectionrequest.DeviceLogCollectionResponse{
	// ...
}


read, err := client.CreateManagedDeviceLogCollectionRequest(ctx, id, payload, manageddevicelogcollectionrequest.DefaultCreateManagedDeviceLogCollectionRequestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceLogCollectionRequestClient.CreateManagedDeviceLogCollectionRequestDownloadUrl`

```go
ctx := context.TODO()
id := manageddevicelogcollectionrequest.NewDeviceManagementManagedDeviceIdLogCollectionRequestID("managedDeviceId", "deviceLogCollectionResponseId")

read, err := client.CreateManagedDeviceLogCollectionRequestDownloadUrl(ctx, id, manageddevicelogcollectionrequest.DefaultCreateManagedDeviceLogCollectionRequestDownloadUrlOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceLogCollectionRequestClient.DeleteManagedDeviceLogCollectionRequest`

```go
ctx := context.TODO()
id := manageddevicelogcollectionrequest.NewDeviceManagementManagedDeviceIdLogCollectionRequestID("managedDeviceId", "deviceLogCollectionResponseId")

read, err := client.DeleteManagedDeviceLogCollectionRequest(ctx, id, manageddevicelogcollectionrequest.DefaultDeleteManagedDeviceLogCollectionRequestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceLogCollectionRequestClient.GetManagedDeviceLogCollectionRequest`

```go
ctx := context.TODO()
id := manageddevicelogcollectionrequest.NewDeviceManagementManagedDeviceIdLogCollectionRequestID("managedDeviceId", "deviceLogCollectionResponseId")

read, err := client.GetManagedDeviceLogCollectionRequest(ctx, id, manageddevicelogcollectionrequest.DefaultGetManagedDeviceLogCollectionRequestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceLogCollectionRequestClient.GetManagedDeviceLogCollectionRequestsCount`

```go
ctx := context.TODO()
id := manageddevicelogcollectionrequest.NewDeviceManagementManagedDeviceID("managedDeviceId")

read, err := client.GetManagedDeviceLogCollectionRequestsCount(ctx, id, manageddevicelogcollectionrequest.DefaultGetManagedDeviceLogCollectionRequestsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceLogCollectionRequestClient.ListManagedDeviceLogCollectionRequests`

```go
ctx := context.TODO()
id := manageddevicelogcollectionrequest.NewDeviceManagementManagedDeviceID("managedDeviceId")

// alternatively `client.ListManagedDeviceLogCollectionRequests(ctx, id, manageddevicelogcollectionrequest.DefaultListManagedDeviceLogCollectionRequestsOperationOptions())` can be used to do batched pagination
items, err := client.ListManagedDeviceLogCollectionRequestsComplete(ctx, id, manageddevicelogcollectionrequest.DefaultListManagedDeviceLogCollectionRequestsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagedDeviceLogCollectionRequestClient.UpdateManagedDeviceLogCollectionRequest`

```go
ctx := context.TODO()
id := manageddevicelogcollectionrequest.NewDeviceManagementManagedDeviceIdLogCollectionRequestID("managedDeviceId", "deviceLogCollectionResponseId")

payload := manageddevicelogcollectionrequest.DeviceLogCollectionResponse{
	// ...
}


read, err := client.UpdateManagedDeviceLogCollectionRequest(ctx, id, payload, manageddevicelogcollectionrequest.DefaultUpdateManagedDeviceLogCollectionRequestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
