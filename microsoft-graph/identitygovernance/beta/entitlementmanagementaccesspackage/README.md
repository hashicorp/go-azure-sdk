
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackage` Documentation

The `entitlementmanagementaccesspackage` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackage"
```


### Client Initialization

```go
client := entitlementmanagementaccesspackage.NewEntitlementManagementAccessPackageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EntitlementManagementAccessPackageClient.CreateEntitlementManagementAccessPackage`

```go
ctx := context.TODO()

payload := entitlementmanagementaccesspackage.AccessPackage{
	// ...
}


read, err := client.CreateEntitlementManagementAccessPackage(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAccessPackageClient.DeleteEntitlementManagementAccessPackage`

```go
ctx := context.TODO()
id := entitlementmanagementaccesspackage.NewIdentityGovernanceEntitlementManagementAccessPackageID("accessPackageIdValue")

read, err := client.DeleteEntitlementManagementAccessPackage(ctx, id, entitlementmanagementaccesspackage.DefaultDeleteEntitlementManagementAccessPackageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAccessPackageClient.GetEntitlementManagementAccessPackage`

```go
ctx := context.TODO()
id := entitlementmanagementaccesspackage.NewIdentityGovernanceEntitlementManagementAccessPackageID("accessPackageIdValue")

read, err := client.GetEntitlementManagementAccessPackage(ctx, id, entitlementmanagementaccesspackage.DefaultGetEntitlementManagementAccessPackageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAccessPackageClient.GetEntitlementManagementAccessPackageApplicablePolicyRequirements`

```go
ctx := context.TODO()
id := entitlementmanagementaccesspackage.NewIdentityGovernanceEntitlementManagementAccessPackageID("accessPackageIdValue")

// alternatively `client.GetEntitlementManagementAccessPackageApplicablePolicyRequirements(ctx, id, entitlementmanagementaccesspackage.DefaultGetEntitlementManagementAccessPackageApplicablePolicyRequirementsOperationOptions())` can be used to do batched pagination
items, err := client.GetEntitlementManagementAccessPackageApplicablePolicyRequirementsComplete(ctx, id, entitlementmanagementaccesspackage.DefaultGetEntitlementManagementAccessPackageApplicablePolicyRequirementsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EntitlementManagementAccessPackageClient.GetEntitlementManagementAccessPackagesCount`

```go
ctx := context.TODO()


read, err := client.GetEntitlementManagementAccessPackagesCount(ctx, entitlementmanagementaccesspackage.DefaultGetEntitlementManagementAccessPackagesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAccessPackageClient.ListEntitlementManagementAccessPackages`

```go
ctx := context.TODO()


// alternatively `client.ListEntitlementManagementAccessPackages(ctx, entitlementmanagementaccesspackage.DefaultListEntitlementManagementAccessPackagesOperationOptions())` can be used to do batched pagination
items, err := client.ListEntitlementManagementAccessPackagesComplete(ctx, entitlementmanagementaccesspackage.DefaultListEntitlementManagementAccessPackagesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EntitlementManagementAccessPackageClient.MoveEntitlementManagementAccessPackageToCatalog`

```go
ctx := context.TODO()
id := entitlementmanagementaccesspackage.NewIdentityGovernanceEntitlementManagementAccessPackageID("accessPackageIdValue")

payload := entitlementmanagementaccesspackage.MoveEntitlementManagementAccessPackageToCatalogRequest{
	// ...
}


read, err := client.MoveEntitlementManagementAccessPackageToCatalog(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementAccessPackageClient.UpdateEntitlementManagementAccessPackage`

```go
ctx := context.TODO()
id := entitlementmanagementaccesspackage.NewIdentityGovernanceEntitlementManagementAccessPackageID("accessPackageIdValue")

payload := entitlementmanagementaccesspackage.AccessPackage{
	// ...
}


read, err := client.UpdateEntitlementManagementAccessPackage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
