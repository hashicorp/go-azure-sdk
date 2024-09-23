
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackageassignmentaccesspackageassignmentrequest` Documentation

The `entitlementmanagementaccesspackageassignmentaccesspackageassignmentrequest` SDK allows for interaction with Microsoft Graph `identitygovernance` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackageassignmentaccesspackageassignmentrequest"
```


### Client Initialization

```go
client := entitlementmanagementaccesspackageassignmentaccesspackageassignmentrequest.NewEntitlementManagementAccessPackageAssignmentAccessPackageAssignmentRequestClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentRequestClient.CancelEntitlementManagementAccessPackageAssignmentRequest`

```go
ctx := context.TODO()
id := entitlementmanagementaccesspackageassignmentaccesspackageassignmentrequest.NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestID("accessPackageAssignmentId", "accessPackageAssignmentRequestId")

read, err := client.CancelEntitlementManagementAccessPackageAssignmentRequest(ctx, id, entitlementmanagementaccesspackageassignmentaccesspackageassignmentrequest.DefaultCancelEntitlementManagementAccessPackageAssignmentRequestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentRequestClient.CreateEntitlementManagementAccessPackageAssignmentRequest`

```go
ctx := context.TODO()
id := entitlementmanagementaccesspackageassignmentaccesspackageassignmentrequest.NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentID("accessPackageAssignmentId")

payload := entitlementmanagementaccesspackageassignmentaccesspackageassignmentrequest.AccessPackageAssignmentRequest{
	// ...
}


read, err := client.CreateEntitlementManagementAccessPackageAssignmentRequest(ctx, id, payload, entitlementmanagementaccesspackageassignmentaccesspackageassignmentrequest.DefaultCreateEntitlementManagementAccessPackageAssignmentRequestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentRequestClient.CreateEntitlementManagementAccessPackageAssignmentRequestResume`

```go
ctx := context.TODO()
id := entitlementmanagementaccesspackageassignmentaccesspackageassignmentrequest.NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestID("accessPackageAssignmentId", "accessPackageAssignmentRequestId")

payload := entitlementmanagementaccesspackageassignmentaccesspackageassignmentrequest.CreateEntitlementManagementAccessPackageAssignmentRequestResumeRequest{
	// ...
}


read, err := client.CreateEntitlementManagementAccessPackageAssignmentRequestResume(ctx, id, payload, entitlementmanagementaccesspackageassignmentaccesspackageassignmentrequest.DefaultCreateEntitlementManagementAccessPackageAssignmentRequestResumeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentRequestClient.DeleteEntitlementManagementAccessPackageAssignmentRequest`

```go
ctx := context.TODO()
id := entitlementmanagementaccesspackageassignmentaccesspackageassignmentrequest.NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestID("accessPackageAssignmentId", "accessPackageAssignmentRequestId")

read, err := client.DeleteEntitlementManagementAccessPackageAssignmentRequest(ctx, id, entitlementmanagementaccesspackageassignmentaccesspackageassignmentrequest.DefaultDeleteEntitlementManagementAccessPackageAssignmentRequestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentRequestClient.GetEntitlementManagementAccessPackageAssignmentRequest`

```go
ctx := context.TODO()
id := entitlementmanagementaccesspackageassignmentaccesspackageassignmentrequest.NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestID("accessPackageAssignmentId", "accessPackageAssignmentRequestId")

read, err := client.GetEntitlementManagementAccessPackageAssignmentRequest(ctx, id, entitlementmanagementaccesspackageassignmentaccesspackageassignmentrequest.DefaultGetEntitlementManagementAccessPackageAssignmentRequestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentRequestClient.GetEntitlementManagementAccessPackageAssignmentRequestsCount`

```go
ctx := context.TODO()
id := entitlementmanagementaccesspackageassignmentaccesspackageassignmentrequest.NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentID("accessPackageAssignmentId")

read, err := client.GetEntitlementManagementAccessPackageAssignmentRequestsCount(ctx, id, entitlementmanagementaccesspackageassignmentaccesspackageassignmentrequest.DefaultGetEntitlementManagementAccessPackageAssignmentRequestsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentRequestClient.ListEntitlementManagementAccessPackageAssignmentRequests`

```go
ctx := context.TODO()
id := entitlementmanagementaccesspackageassignmentaccesspackageassignmentrequest.NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentID("accessPackageAssignmentId")

// alternatively `client.ListEntitlementManagementAccessPackageAssignmentRequests(ctx, id, entitlementmanagementaccesspackageassignmentaccesspackageassignmentrequest.DefaultListEntitlementManagementAccessPackageAssignmentRequestsOperationOptions())` can be used to do batched pagination
items, err := client.ListEntitlementManagementAccessPackageAssignmentRequestsComplete(ctx, id, entitlementmanagementaccesspackageassignmentaccesspackageassignmentrequest.DefaultListEntitlementManagementAccessPackageAssignmentRequestsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentRequestClient.ReprocessEntitlementManagementAccessPackageAssignmentRequest`

```go
ctx := context.TODO()
id := entitlementmanagementaccesspackageassignmentaccesspackageassignmentrequest.NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestID("accessPackageAssignmentId", "accessPackageAssignmentRequestId")

read, err := client.ReprocessEntitlementManagementAccessPackageAssignmentRequest(ctx, id, entitlementmanagementaccesspackageassignmentaccesspackageassignmentrequest.DefaultReprocessEntitlementManagementAccessPackageAssignmentRequestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentRequestClient.UpdateEntitlementManagementAccessPackageAssignmentRequest`

```go
ctx := context.TODO()
id := entitlementmanagementaccesspackageassignmentaccesspackageassignmentrequest.NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestID("accessPackageAssignmentId", "accessPackageAssignmentRequestId")

payload := entitlementmanagementaccesspackageassignmentaccesspackageassignmentrequest.AccessPackageAssignmentRequest{
	// ...
}


read, err := client.UpdateEntitlementManagementAccessPackageAssignmentRequest(ctx, id, payload, entitlementmanagementaccesspackageassignmentaccesspackageassignmentrequest.DefaultUpdateEntitlementManagementAccessPackageAssignmentRequestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
