
## `github.com/hashicorp/go-azure-sdk/resource-manager/resources/2023-04-01/policysetdefinitionversions` Documentation

The `policysetdefinitionversions` SDK allows for interaction with the Azure Resource Manager Service `resources` (API Version `2023-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/resources/2023-04-01/policysetdefinitionversions"
```


### Client Initialization

```go
client := policysetdefinitionversions.NewPolicySetDefinitionVersionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PolicySetDefinitionVersionsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := policysetdefinitionversions.NewProviderPolicySetDefinitionVersionID("12345678-1234-9876-4563-123456789012", "policySetDefinitionValue", "versionValue")

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
id := policysetdefinitionversions.NewProviders2PolicySetDefinitionVersionID("managementGroupValue", "policySetDefinitionValue", "versionValue")

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
id := policysetdefinitionversions.NewProviderPolicySetDefinitionVersionID("12345678-1234-9876-4563-123456789012", "policySetDefinitionValue", "versionValue")

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
id := policysetdefinitionversions.NewProviders2PolicySetDefinitionVersionID("managementGroupValue", "policySetDefinitionValue", "versionValue")

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
id := policysetdefinitionversions.NewProviderPolicySetDefinitionVersionID("12345678-1234-9876-4563-123456789012", "policySetDefinitionValue", "versionValue")

read, err := client.Get(ctx, id)
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
id := policysetdefinitionversions.NewProviders2PolicySetDefinitionVersionID("managementGroupValue", "policySetDefinitionValue", "versionValue")

read, err := client.GetAtManagementGroup(ctx, id)
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
id := policysetdefinitionversions.NewPolicySetDefinitionVersionID("policySetDefinitionValue", "versionValue")

read, err := client.GetBuiltIn(ctx, id)
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
id := policysetdefinitionversions.NewProviderPolicySetDefinitionID("12345678-1234-9876-4563-123456789012", "policySetDefinitionValue")

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
id := policysetdefinitionversions.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.ListAll(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicySetDefinitionVersionsClient.ListAllAtManagementGroup`

```go
ctx := context.TODO()
id := policysetdefinitionversions.NewManagementGroupID("groupIdValue")

read, err := client.ListAllAtManagementGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicySetDefinitionVersionsClient.ListAllBuiltins`

```go
ctx := context.TODO()


read, err := client.ListAllBuiltins(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicySetDefinitionVersionsClient.ListBuiltIn`

```go
ctx := context.TODO()
id := policysetdefinitionversions.NewPolicySetDefinitionID("policySetDefinitionValue")

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
id := policysetdefinitionversions.NewProviders2PolicySetDefinitionID("managementGroupValue", "policySetDefinitionValue")

// alternatively `client.ListByManagementGroup(ctx, id, policysetdefinitionversions.DefaultListByManagementGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListByManagementGroupComplete(ctx, id, policysetdefinitionversions.DefaultListByManagementGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
