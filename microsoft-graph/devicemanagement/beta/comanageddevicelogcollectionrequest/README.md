
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/comanageddevicelogcollectionrequest` Documentation

The `comanageddevicelogcollectionrequest` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/comanageddevicelogcollectionrequest"
```


### Client Initialization

```go
client := comanageddevicelogcollectionrequest.NewComanagedDeviceLogCollectionRequestClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ComanagedDeviceLogCollectionRequestClient.CreateComanagedDeviceLogCollectionRequest`

```go
ctx := context.TODO()
id := comanageddevicelogcollectionrequest.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevicelogcollectionrequest.DeviceLogCollectionResponse{
	// ...
}


read, err := client.CreateComanagedDeviceLogCollectionRequest(ctx, id, payload)
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
id := comanageddevicelogcollectionrequest.NewDeviceManagementComanagedDeviceIdLogCollectionRequestID("managedDeviceIdValue", "deviceLogCollectionResponseIdValue")

read, err := client.CreateComanagedDeviceLogCollectionRequestDownloadUrl(ctx, id)
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
id := comanageddevicelogcollectionrequest.NewDeviceManagementComanagedDeviceIdLogCollectionRequestID("managedDeviceIdValue", "deviceLogCollectionResponseIdValue")

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
id := comanageddevicelogcollectionrequest.NewDeviceManagementComanagedDeviceIdLogCollectionRequestID("managedDeviceIdValue", "deviceLogCollectionResponseIdValue")

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
id := comanageddevicelogcollectionrequest.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

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
id := comanageddevicelogcollectionrequest.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

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
id := comanageddevicelogcollectionrequest.NewDeviceManagementComanagedDeviceIdLogCollectionRequestID("managedDeviceIdValue", "deviceLogCollectionResponseIdValue")

payload := comanageddevicelogcollectionrequest.DeviceLogCollectionResponse{
	// ...
}


read, err := client.UpdateComanagedDeviceLogCollectionRequest(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
