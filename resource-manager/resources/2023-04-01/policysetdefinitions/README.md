
## `github.com/hashicorp/go-azure-sdk/resource-manager/resources/2023-04-01/policysetdefinitions` Documentation

The `policysetdefinitions` SDK allows for interaction with the Azure Resource Manager Service `resources` (API Version `2023-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/resources/2023-04-01/policysetdefinitions"
```


### Client Initialization

```go
client := policysetdefinitions.NewPolicySetDefinitionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PolicySetDefinitionsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := policysetdefinitions.NewProviderPolicySetDefinitionID("12345678-1234-9876-4563-123456789012", "policySetDefinitionValue")

payload := policysetdefinitions.PolicySetDefinition{
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


### Example Usage: `PolicySetDefinitionsClient.CreateOrUpdateAtManagementGroup`

```go
ctx := context.TODO()
id := policysetdefinitions.NewProviders2PolicySetDefinitionID("managementGroupValue", "policySetDefinitionValue")

payload := policysetdefinitions.PolicySetDefinition{
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


### Example Usage: `PolicySetDefinitionsClient.Delete`

```go
ctx := context.TODO()
id := policysetdefinitions.NewProviderPolicySetDefinitionID("12345678-1234-9876-4563-123456789012", "policySetDefinitionValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicySetDefinitionsClient.DeleteAtManagementGroup`

```go
ctx := context.TODO()
id := policysetdefinitions.NewProviders2PolicySetDefinitionID("managementGroupValue", "policySetDefinitionValue")

read, err := client.DeleteAtManagementGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicySetDefinitionsClient.Get`

```go
ctx := context.TODO()
id := policysetdefinitions.NewProviderPolicySetDefinitionID("12345678-1234-9876-4563-123456789012", "policySetDefinitionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicySetDefinitionsClient.GetAtManagementGroup`

```go
ctx := context.TODO()
id := policysetdefinitions.NewProviders2PolicySetDefinitionID("managementGroupValue", "policySetDefinitionValue")

read, err := client.GetAtManagementGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicySetDefinitionsClient.GetBuiltIn`

```go
ctx := context.TODO()
id := policysetdefinitions.NewPolicySetDefinitionID("policySetDefinitionValue")

read, err := client.GetBuiltIn(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicySetDefinitionsClient.List`

```go
ctx := context.TODO()
id := policysetdefinitions.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id, policysetdefinitions.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, policysetdefinitions.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicySetDefinitionsClient.ListBuiltIn`

```go
ctx := context.TODO()


// alternatively `client.ListBuiltIn(ctx, policysetdefinitions.DefaultListBuiltInOperationOptions())` can be used to do batched pagination
items, err := client.ListBuiltInComplete(ctx, policysetdefinitions.DefaultListBuiltInOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicySetDefinitionsClient.ListByManagementGroup`

```go
ctx := context.TODO()
id := policysetdefinitions.NewManagementGroupID("groupIdValue")

// alternatively `client.ListByManagementGroup(ctx, id, policysetdefinitions.DefaultListByManagementGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListByManagementGroupComplete(ctx, id, policysetdefinitions.DefaultListByManagementGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
