
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementassignmentrequest` Documentation

The `entitlementmanagementassignmentrequest` SDK allows for interaction with Microsoft Graph `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementassignmentrequest"
```


### Client Initialization

```go
client := entitlementmanagementassignmentrequest.NewEntitlementManagementAssignmentRequestClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EntitlementManagementAssignmentRequestClient.CancelEntitlementManagementAssignmentRequest`

```go
ctx := context.TODO()
id := entitlementmanagementassignmentrequest.NewIdentityGovernanceEntitlementManagementAssignmentRequestID("accessPackageAssignmentRequestId")

read, err := client.CancelEntitlementManagementAssignmentRequest(ctx, id, entitlementmanagementassignmentrequest.DefaultCancelEntitlementManagementAssignmentRequestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAssignmentRequestClient.CreateEntitlementManagementAssignmentRequest`

```go
ctx := context.TODO()

payload := entitlementmanagementassignmentrequest.AccessPackageAssignmentRequest{
	// ...
}


read, err := client.CreateEntitlementManagementAssignmentRequest(ctx, payload, entitlementmanagementassignmentrequest.DefaultCreateEntitlementManagementAssignmentRequestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAssignmentRequestClient.CreateEntitlementManagementAssignmentRequestResume`

```go
ctx := context.TODO()
id := entitlementmanagementassignmentrequest.NewIdentityGovernanceEntitlementManagementAssignmentRequestID("accessPackageAssignmentRequestId")

payload := entitlementmanagementassignmentrequest.CreateEntitlementManagementAssignmentRequestResumeRequest{
	// ...
}


read, err := client.CreateEntitlementManagementAssignmentRequestResume(ctx, id, payload, entitlementmanagementassignmentrequest.DefaultCreateEntitlementManagementAssignmentRequestResumeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAssignmentRequestClient.DeleteEntitlementManagementAssignmentRequest`

```go
ctx := context.TODO()
id := entitlementmanagementassignmentrequest.NewIdentityGovernanceEntitlementManagementAssignmentRequestID("accessPackageAssignmentRequestId")

read, err := client.DeleteEntitlementManagementAssignmentRequest(ctx, id, entitlementmanagementassignmentrequest.DefaultDeleteEntitlementManagementAssignmentRequestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAssignmentRequestClient.GetEntitlementManagementAssignmentRequest`

```go
ctx := context.TODO()
id := entitlementmanagementassignmentrequest.NewIdentityGovernanceEntitlementManagementAssignmentRequestID("accessPackageAssignmentRequestId")

read, err := client.GetEntitlementManagementAssignmentRequest(ctx, id, entitlementmanagementassignmentrequest.DefaultGetEntitlementManagementAssignmentRequestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAssignmentRequestClient.GetEntitlementManagementAssignmentRequestsCount`

```go
ctx := context.TODO()


read, err := client.GetEntitlementManagementAssignmentRequestsCount(ctx, entitlementmanagementassignmentrequest.DefaultGetEntitlementManagementAssignmentRequestsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAssignmentRequestClient.ListEntitlementManagementAssignmentRequests`

```go
ctx := context.TODO()


// alternatively `client.ListEntitlementManagementAssignmentRequests(ctx, entitlementmanagementassignmentrequest.DefaultListEntitlementManagementAssignmentRequestsOperationOptions())` can be used to do batched pagination
items, err := client.ListEntitlementManagementAssignmentRequestsComplete(ctx, entitlementmanagementassignmentrequest.DefaultListEntitlementManagementAssignmentRequestsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EntitlementManagementAssignmentRequestClient.ReprocessEntitlementManagementAssignmentRequest`

```go
ctx := context.TODO()
id := entitlementmanagementassignmentrequest.NewIdentityGovernanceEntitlementManagementAssignmentRequestID("accessPackageAssignmentRequestId")

read, err := client.ReprocessEntitlementManagementAssignmentRequest(ctx, id, entitlementmanagementassignmentrequest.DefaultReprocessEntitlementManagementAssignmentRequestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAssignmentRequestClient.UpdateEntitlementManagementAssignmentRequest`

```go
ctx := context.TODO()
id := entitlementmanagementassignmentrequest.NewIdentityGovernanceEntitlementManagementAssignmentRequestID("accessPackageAssignmentRequestId")

payload := entitlementmanagementassignmentrequest.AccessPackageAssignmentRequest{
	// ...
}


read, err := client.UpdateEntitlementManagementAssignmentRequest(ctx, id, payload, entitlementmanagementassignmentrequest.DefaultUpdateEntitlementManagementAssignmentRequestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
