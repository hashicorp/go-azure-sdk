
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackageassignmentrequest` Documentation

The `entitlementmanagementaccesspackageassignmentrequest` SDK allows for interaction with Microsoft Graph `identitygovernance` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackageassignmentrequest"
```


### Client Initialization

```go
client := entitlementmanagementaccesspackageassignmentrequest.NewEntitlementManagementAccessPackageAssignmentRequestClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EntitlementManagementAccessPackageAssignmentRequestClient.CancelEntitlementManagementAccessPackageAssignmentRequest`

```go
ctx := context.TODO()
id := entitlementmanagementaccesspackageassignmentrequest.NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestID("accessPackageAssignmentRequestId")

read, err := client.CancelEntitlementManagementAccessPackageAssignmentRequest(ctx, id, entitlementmanagementaccesspackageassignmentrequest.DefaultCancelEntitlementManagementAccessPackageAssignmentRequestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAccessPackageAssignmentRequestClient.CreateEntitlementManagementAccessPackageAssignmentRequest`

```go
ctx := context.TODO()

payload := entitlementmanagementaccesspackageassignmentrequest.AccessPackageAssignmentRequest{
	// ...
}


read, err := client.CreateEntitlementManagementAccessPackageAssignmentRequest(ctx, payload, entitlementmanagementaccesspackageassignmentrequest.DefaultCreateEntitlementManagementAccessPackageAssignmentRequestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAccessPackageAssignmentRequestClient.CreateEntitlementManagementAccessPackageAssignmentRequestResume`

```go
ctx := context.TODO()
id := entitlementmanagementaccesspackageassignmentrequest.NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestID("accessPackageAssignmentRequestId")

payload := entitlementmanagementaccesspackageassignmentrequest.CreateEntitlementManagementAccessPackageAssignmentRequestResumeRequest{
	// ...
}


read, err := client.CreateEntitlementManagementAccessPackageAssignmentRequestResume(ctx, id, payload, entitlementmanagementaccesspackageassignmentrequest.DefaultCreateEntitlementManagementAccessPackageAssignmentRequestResumeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAccessPackageAssignmentRequestClient.DeleteEntitlementManagementAccessPackageAssignmentRequest`

```go
ctx := context.TODO()
id := entitlementmanagementaccesspackageassignmentrequest.NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestID("accessPackageAssignmentRequestId")

read, err := client.DeleteEntitlementManagementAccessPackageAssignmentRequest(ctx, id, entitlementmanagementaccesspackageassignmentrequest.DefaultDeleteEntitlementManagementAccessPackageAssignmentRequestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAccessPackageAssignmentRequestClient.GetEntitlementManagementAccessPackageAssignmentRequest`

```go
ctx := context.TODO()
id := entitlementmanagementaccesspackageassignmentrequest.NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestID("accessPackageAssignmentRequestId")

read, err := client.GetEntitlementManagementAccessPackageAssignmentRequest(ctx, id, entitlementmanagementaccesspackageassignmentrequest.DefaultGetEntitlementManagementAccessPackageAssignmentRequestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAccessPackageAssignmentRequestClient.GetEntitlementManagementAccessPackageAssignmentRequestsCount`

```go
ctx := context.TODO()


read, err := client.GetEntitlementManagementAccessPackageAssignmentRequestsCount(ctx, entitlementmanagementaccesspackageassignmentrequest.DefaultGetEntitlementManagementAccessPackageAssignmentRequestsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAccessPackageAssignmentRequestClient.ListEntitlementManagementAccessPackageAssignmentRequests`

```go
ctx := context.TODO()


// alternatively `client.ListEntitlementManagementAccessPackageAssignmentRequests(ctx, entitlementmanagementaccesspackageassignmentrequest.DefaultListEntitlementManagementAccessPackageAssignmentRequestsOperationOptions())` can be used to do batched pagination
items, err := client.ListEntitlementManagementAccessPackageAssignmentRequestsComplete(ctx, entitlementmanagementaccesspackageassignmentrequest.DefaultListEntitlementManagementAccessPackageAssignmentRequestsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EntitlementManagementAccessPackageAssignmentRequestClient.ReprocessEntitlementManagementAccessPackageAssignmentRequest`

```go
ctx := context.TODO()
id := entitlementmanagementaccesspackageassignmentrequest.NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestID("accessPackageAssignmentRequestId")

read, err := client.ReprocessEntitlementManagementAccessPackageAssignmentRequest(ctx, id, entitlementmanagementaccesspackageassignmentrequest.DefaultReprocessEntitlementManagementAccessPackageAssignmentRequestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAccessPackageAssignmentRequestClient.UpdateEntitlementManagementAccessPackageAssignmentRequest`

```go
ctx := context.TODO()
id := entitlementmanagementaccesspackageassignmentrequest.NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestID("accessPackageAssignmentRequestId")

payload := entitlementmanagementaccesspackageassignmentrequest.AccessPackageAssignmentRequest{
	// ...
}


read, err := client.UpdateEntitlementManagementAccessPackageAssignmentRequest(ctx, id, payload, entitlementmanagementaccesspackageassignmentrequest.DefaultUpdateEntitlementManagementAccessPackageAssignmentRequestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
