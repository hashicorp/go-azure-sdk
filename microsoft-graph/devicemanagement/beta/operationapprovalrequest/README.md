
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/operationapprovalrequest` Documentation

The `operationapprovalrequest` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/operationapprovalrequest"
```


### Client Initialization

```go
client := operationapprovalrequest.NewOperationApprovalRequestClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OperationApprovalRequestClient.CancelOperationApprovalRequestApproval`

```go
ctx := context.TODO()
id := operationapprovalrequest.NewDeviceManagementOperationApprovalRequestID("operationApprovalRequestIdValue")

payload := operationapprovalrequest.CancelOperationApprovalRequestApprovalRequest{
	// ...
}


read, err := client.CancelOperationApprovalRequestApproval(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OperationApprovalRequestClient.CancelOperationApprovalRequestsMyRequest`

```go
ctx := context.TODO()

payload := operationapprovalrequest.CancelOperationApprovalRequestsMyRequestRequest{
	// ...
}


read, err := client.CancelOperationApprovalRequestsMyRequest(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OperationApprovalRequestClient.CreateOperationApprovalRequest`

```go
ctx := context.TODO()

payload := operationapprovalrequest.OperationApprovalRequest{
	// ...
}


read, err := client.CreateOperationApprovalRequest(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OperationApprovalRequestClient.CreateOperationApprovalRequestApprove`

```go
ctx := context.TODO()
id := operationapprovalrequest.NewDeviceManagementOperationApprovalRequestID("operationApprovalRequestIdValue")

payload := operationapprovalrequest.CreateOperationApprovalRequestApproveRequest{
	// ...
}


read, err := client.CreateOperationApprovalRequestApprove(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OperationApprovalRequestClient.CreateOperationApprovalRequestReject`

```go
ctx := context.TODO()
id := operationapprovalrequest.NewDeviceManagementOperationApprovalRequestID("operationApprovalRequestIdValue")

payload := operationapprovalrequest.CreateOperationApprovalRequestRejectRequest{
	// ...
}


read, err := client.CreateOperationApprovalRequestReject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OperationApprovalRequestClient.CreateOperationApprovalRequestRetrieveRequestStatus`

```go
ctx := context.TODO()

payload := operationapprovalrequest.CreateOperationApprovalRequestRetrieveRequestStatusRequest{
	// ...
}


read, err := client.CreateOperationApprovalRequestRetrieveRequestStatus(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OperationApprovalRequestClient.DeleteOperationApprovalRequest`

```go
ctx := context.TODO()
id := operationapprovalrequest.NewDeviceManagementOperationApprovalRequestID("operationApprovalRequestIdValue")

read, err := client.DeleteOperationApprovalRequest(ctx, id, operationapprovalrequest.DefaultDeleteOperationApprovalRequestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OperationApprovalRequestClient.GetOperationApprovalRequest`

```go
ctx := context.TODO()
id := operationapprovalrequest.NewDeviceManagementOperationApprovalRequestID("operationApprovalRequestIdValue")

read, err := client.GetOperationApprovalRequest(ctx, id, operationapprovalrequest.DefaultGetOperationApprovalRequestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OperationApprovalRequestClient.GetOperationApprovalRequestsCount`

```go
ctx := context.TODO()


read, err := client.GetOperationApprovalRequestsCount(ctx, operationapprovalrequest.DefaultGetOperationApprovalRequestsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OperationApprovalRequestClient.ListOperationApprovalRequests`

```go
ctx := context.TODO()


// alternatively `client.ListOperationApprovalRequests(ctx, operationapprovalrequest.DefaultListOperationApprovalRequestsOperationOptions())` can be used to do batched pagination
items, err := client.ListOperationApprovalRequestsComplete(ctx, operationapprovalrequest.DefaultListOperationApprovalRequestsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OperationApprovalRequestClient.UpdateOperationApprovalRequest`

```go
ctx := context.TODO()
id := operationapprovalrequest.NewDeviceManagementOperationApprovalRequestID("operationApprovalRequestIdValue")

payload := operationapprovalrequest.OperationApprovalRequest{
	// ...
}


read, err := client.UpdateOperationApprovalRequest(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
