
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackageassignmentrequest` Documentation

The `entitlementmanagementaccesspackageassignmentrequest` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackageassignmentrequest"
```


### Client Initialization

```go
client := entitlementmanagementaccesspackageassignmentrequest.NewEntitlementManagementAccessPackageAssignmentRequestClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EntitlementManagementAccessPackageAssignmentRequestClient.CancelEntitlementManagementAccessPackageAssignmentRequest`

```go
ctx := context.TODO()
id := entitlementmanagementaccesspackageassignmentrequest.NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestID("accessPackageAssignmentRequestIdValue")

read, err := client.CancelEntitlementManagementAccessPackageAssignmentRequest(ctx, id)
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


read, err := client.CreateEntitlementManagementAccessPackageAssignmentRequest(ctx, payload)
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
id := entitlementmanagementaccesspackageassignmentrequest.NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestID("accessPackageAssignmentRequestIdValue")

payload := entitlementmanagementaccesspackageassignmentrequest.CreateEntitlementManagementAccessPackageAssignmentRequestResumeRequest{
	// ...
}


read, err := client.CreateEntitlementManagementAccessPackageAssignmentRequestResume(ctx, id, payload)
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
id := entitlementmanagementaccesspackageassignmentrequest.NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestID("accessPackageAssignmentRequestIdValue")

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
id := entitlementmanagementaccesspackageassignmentrequest.NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestID("accessPackageAssignmentRequestIdValue")

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
id := entitlementmanagementaccesspackageassignmentrequest.NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestID("accessPackageAssignmentRequestIdValue")

read, err := client.ReprocessEntitlementManagementAccessPackageAssignmentRequest(ctx, id)
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
id := entitlementmanagementaccesspackageassignmentrequest.NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestID("accessPackageAssignmentRequestIdValue")

payload := entitlementmanagementaccesspackageassignmentrequest.AccessPackageAssignmentRequest{
	// ...
}


read, err := client.UpdateEntitlementManagementAccessPackageAssignmentRequest(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
