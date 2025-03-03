
## `github.com/hashicorp/go-azure-sdk/resource-manager/resources/2025-03-01/policysetdefinitionversions` Documentation

The `policysetdefinitionversions` SDK allows for interaction with Azure Resource Manager `resources` (API Version `2025-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/resources/2025-03-01/policysetdefinitionversions"
```


### Client Initialization

```go
client := policysetdefinitionversions.NewPolicySetDefinitionVersionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PolicySetDefinitionVersionsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := policysetdefinitionversions.NewProviderPolicySetDefinitionVersionID("12345678-1234-9876-4563-123456789012", "policySetDefinitionName", "versionName")

payload := policysetdefinitionversions.PolicySetDefinitionVersion{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicySetDefinitionVersionsClient.CreateOrUpdateAtManagementGroup`

```go
ctx := context.TODO()
id := policysetdefinitionversions.NewProviders2PolicySetDefinitionVersionID("managementGroupName", "policySetDefinitionName", "versionName")

payload := policysetdefinitionversions.PolicySetDefinitionVersion{
	// ...
}


read, err := client.CreateOrUpdateAtManagementGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicySetDefinitionVersionsClient.Delete`

```go
ctx := context.TODO()
id := policysetdefinitionversions.NewProviderPolicySetDefinitionVersionID("12345678-1234-9876-4563-123456789012", "policySetDefinitionName", "versionName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicySetDefinitionVersionsClient.DeleteAtManagementGroup`

```go
ctx := context.TODO()
id := policysetdefinitionversions.NewProviders2PolicySetDefinitionVersionID("managementGroupName", "policySetDefinitionName", "versionName")

read, err := client.DeleteAtManagementGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicySetDefinitionVersionsClient.Get`

```go
ctx := context.TODO()
id := policysetdefinitionversions.NewProviderPolicySetDefinitionVersionID("12345678-1234-9876-4563-123456789012", "policySetDefinitionName", "versionName")

read, err := client.Get(ctx, id, policysetdefinitionversions.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicySetDefinitionVersionsClient.GetAtManagementGroup`

```go
ctx := context.TODO()
id := policysetdefinitionversions.NewProviders2PolicySetDefinitionVersionID("managementGroupName", "policySetDefinitionName", "versionName")

read, err := client.GetAtManagementGroup(ctx, id, policysetdefinitionversions.DefaultGetAtManagementGroupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicySetDefinitionVersionsClient.GetBuiltIn`

```go
ctx := context.TODO()
id := policysetdefinitionversions.NewPolicySetDefinitionVersionID("policySetDefinitionName", "versionName")

read, err := client.GetBuiltIn(ctx, id, policysetdefinitionversions.DefaultGetBuiltInOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicySetDefinitionVersionsClient.List`

```go
ctx := context.TODO()
id := policysetdefinitionversions.NewProviderPolicySetDefinitionID("12345678-1234-9876-4563-123456789012", "policySetDefinitionName")

// alternatively `client.List(ctx, id, policysetdefinitionversions.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, policysetdefinitionversions.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicySetDefinitionVersionsClient.ListAll`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListAll(ctx, id)` can be used to do batched pagination
items, err := client.ListAllComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicySetDefinitionVersionsClient.ListAllAtManagementGroup`

```go
ctx := context.TODO()
id := commonids.NewManagementGroupID("groupId")

// alternatively `client.ListAllAtManagementGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListAllAtManagementGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicySetDefinitionVersionsClient.ListAllBuiltins`

```go
ctx := context.TODO()


// alternatively `client.ListAllBuiltins(ctx)` can be used to do batched pagination
items, err := client.ListAllBuiltinsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicySetDefinitionVersionsClient.ListBuiltIn`

```go
ctx := context.TODO()
id := policysetdefinitionversions.NewPolicySetDefinitionID("policySetDefinitionName")

// alternatively `client.ListBuiltIn(ctx, id, policysetdefinitionversions.DefaultListBuiltInOperationOptions())` can be used to do batched pagination
items, err := client.ListBuiltInComplete(ctx, id, policysetdefinitionversions.DefaultListBuiltInOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicySetDefinitionVersionsClient.ListByManagementGroup`

```go
ctx := context.TODO()
id := policysetdefinitionversions.NewProviders2PolicySetDefinitionID("managementGroupName", "policySetDefinitionName")

// alternatively `client.ListByManagementGroup(ctx, id, policysetdefinitionversions.DefaultListByManagementGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListByManagementGroupComplete(ctx, id, policysetdefinitionversions.DefaultListByManagementGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
