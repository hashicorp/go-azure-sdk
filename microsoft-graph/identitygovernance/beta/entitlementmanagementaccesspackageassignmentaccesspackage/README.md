
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackageassignmentaccesspackage` Documentation

The `entitlementmanagementaccesspackageassignmentaccesspackage` SDK allows for interaction with Microsoft Graph `identitygovernance` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackageassignmentaccesspackage"
```


### Client Initialization

```go
client := entitlementmanagementaccesspackageassignmentaccesspackage.NewEntitlementManagementAccessPackageAssignmentAccessPackageClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EntitlementManagementAccessPackageAssignmentAccessPackageClient.DeleteEntitlementManagementAccessPackageAssignmentAccessPackage`

```go
ctx := context.TODO()
id := entitlementmanagementaccesspackageassignmentaccesspackage.NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentID("accessPackageAssignmentId")

read, err := client.DeleteEntitlementManagementAccessPackageAssignmentAccessPackage(ctx, id, entitlementmanagementaccesspackageassignmentaccesspackage.DefaultDeleteEntitlementManagementAccessPackageAssignmentAccessPackageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAccessPackageAssignmentAccessPackageClient.GetEntitlementManagementAccessPackageAssignmentAccessPackage`

```go
ctx := context.TODO()
id := entitlementmanagementaccesspackageassignmentaccesspackage.NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentID("accessPackageAssignmentId")

read, err := client.GetEntitlementManagementAccessPackageAssignmentAccessPackage(ctx, id, entitlementmanagementaccesspackageassignmentaccesspackage.DefaultGetEntitlementManagementAccessPackageAssignmentAccessPackageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAccessPackageAssignmentAccessPackageClient.GetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirements`

```go
ctx := context.TODO()
id := entitlementmanagementaccesspackageassignmentaccesspackage.NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentID("accessPackageAssignmentId")

// alternatively `client.GetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirements(ctx, id, entitlementmanagementaccesspackageassignmentaccesspackage.DefaultGetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirementsOperationOptions())` can be used to do batched pagination
items, err := client.GetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirementsComplete(ctx, id, entitlementmanagementaccesspackageassignmentaccesspackage.DefaultGetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirementsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EntitlementManagementAccessPackageAssignmentAccessPackageClient.MoveEntitlementManagementAccessPackageAssignmentToCatalog`

```go
ctx := context.TODO()
id := entitlementmanagementaccesspackageassignmentaccesspackage.NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentID("accessPackageAssignmentId")

payload := entitlementmanagementaccesspackageassignmentaccesspackage.MoveEntitlementManagementAccessPackageAssignmentToCatalogRequest{
	// ...
}


read, err := client.MoveEntitlementManagementAccessPackageAssignmentToCatalog(ctx, id, payload, entitlementmanagementaccesspackageassignmentaccesspackage.DefaultMoveEntitlementManagementAccessPackageAssignmentToCatalogOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAccessPackageAssignmentAccessPackageClient.UpdateEntitlementManagementAccessPackageAssignmentAccessPackage`

```go
ctx := context.TODO()
id := entitlementmanagementaccesspackageassignmentaccesspackage.NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentID("accessPackageAssignmentId")

payload := entitlementmanagementaccesspackageassignmentaccesspackage.AccessPackage{
	// ...
}


read, err := client.UpdateEntitlementManagementAccessPackageAssignmentAccessPackage(ctx, id, payload, entitlementmanagementaccesspackageassignmentaccesspackage.DefaultUpdateEntitlementManagementAccessPackageAssignmentAccessPackageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
