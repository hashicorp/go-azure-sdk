
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementassignmentrequest` Documentation

The `entitlementmanagementassignmentrequest` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementassignmentrequest"
```


### Client Initialization

```go
client := entitlementmanagementassignmentrequest.NewEntitlementManagementAssignmentRequestClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EntitlementManagementAssignmentRequestClient.CreateEntitlementManagementAssignmentRequest`

```go
ctx := context.TODO()

payload := entitlementmanagementassignmentrequest.AccessPackageAssignmentRequest{
	// ...
}


read, err := client.CreateEntitlementManagementAssignmentRequest(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAssignmentRequestClient.CreateEntitlementManagementAssignmentRequestCancel`

```go
ctx := context.TODO()
id := entitlementmanagementassignmentrequest.NewIdentityGovernanceEntitlementManagementAssignmentRequestID("accessPackageAssignmentRequestIdValue")

read, err := client.CreateEntitlementManagementAssignmentRequestCancel(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAssignmentRequestClient.CreateEntitlementManagementAssignmentRequestReproces`

```go
ctx := context.TODO()
id := entitlementmanagementassignmentrequest.NewIdentityGovernanceEntitlementManagementAssignmentRequestID("accessPackageAssignmentRequestIdValue")

read, err := client.CreateEntitlementManagementAssignmentRequestReproces(ctx, id)
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
id := entitlementmanagementassignmentrequest.NewIdentityGovernanceEntitlementManagementAssignmentRequestID("accessPackageAssignmentRequestIdValue")

payload := entitlementmanagementassignmentrequest.CreateEntitlementManagementAssignmentRequestResumeRequest{
	// ...
}


read, err := client.CreateEntitlementManagementAssignmentRequestResume(ctx, id, payload)
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
id := entitlementmanagementassignmentrequest.NewIdentityGovernanceEntitlementManagementAssignmentRequestID("accessPackageAssignmentRequestIdValue")

read, err := client.DeleteEntitlementManagementAssignmentRequest(ctx, id)
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
id := entitlementmanagementassignmentrequest.NewIdentityGovernanceEntitlementManagementAssignmentRequestID("accessPackageAssignmentRequestIdValue")

read, err := client.GetEntitlementManagementAssignmentRequest(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAssignmentRequestClient.GetEntitlementManagementAssignmentRequestCount`

```go
ctx := context.TODO()


read, err := client.GetEntitlementManagementAssignmentRequestCount(ctx)
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


// alternatively `client.ListEntitlementManagementAssignmentRequests(ctx)` can be used to do batched pagination
items, err := client.ListEntitlementManagementAssignmentRequestsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EntitlementManagementAssignmentRequestClient.UpdateEntitlementManagementAssignmentRequest`

```go
ctx := context.TODO()
id := entitlementmanagementassignmentrequest.NewIdentityGovernanceEntitlementManagementAssignmentRequestID("accessPackageAssignmentRequestIdValue")

payload := entitlementmanagementassignmentrequest.AccessPackageAssignmentRequest{
	// ...
}


read, err := client.UpdateEntitlementManagementAssignmentRequest(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
