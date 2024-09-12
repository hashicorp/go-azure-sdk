
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/manageddevicelogcollectionrequest` Documentation

The `manageddevicelogcollectionrequest` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/manageddevicelogcollectionrequest"
```


### Client Initialization

```go
client := manageddevicelogcollectionrequest.NewManagedDeviceLogCollectionRequestClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedDeviceLogCollectionRequestClient.CreateManagedDeviceLogCollectionRequest`

```go
ctx := context.TODO()
id := manageddevicelogcollectionrequest.NewMeManagedDeviceID("managedDeviceIdValue")

payload := manageddevicelogcollectionrequest.DeviceLogCollectionResponse{
	// ...
}


read, err := client.CreateManagedDeviceLogCollectionRequest(ctx, id, payload)
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
id := manageddevicelogcollectionrequest.NewMeManagedDeviceIdLogCollectionRequestID("managedDeviceIdValue", "deviceLogCollectionResponseIdValue")

read, err := client.CreateManagedDeviceLogCollectionRequestDownloadUrl(ctx, id)
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
id := manageddevicelogcollectionrequest.NewMeManagedDeviceIdLogCollectionRequestID("managedDeviceIdValue", "deviceLogCollectionResponseIdValue")

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
id := manageddevicelogcollectionrequest.NewMeManagedDeviceIdLogCollectionRequestID("managedDeviceIdValue", "deviceLogCollectionResponseIdValue")

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
id := manageddevicelogcollectionrequest.NewMeManagedDeviceID("managedDeviceIdValue")

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
id := manageddevicelogcollectionrequest.NewMeManagedDeviceID("managedDeviceIdValue")

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
id := manageddevicelogcollectionrequest.NewMeManagedDeviceIdLogCollectionRequestID("managedDeviceIdValue", "deviceLogCollectionResponseIdValue")

payload := manageddevicelogcollectionrequest.DeviceLogCollectionResponse{
	// ...
}


read, err := client.UpdateManagedDeviceLogCollectionRequest(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
