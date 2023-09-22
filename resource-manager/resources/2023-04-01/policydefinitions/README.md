
## `github.com/hashicorp/go-azure-sdk/resource-manager/resources/2023-04-01/policydefinitions` Documentation

The `policydefinitions` SDK allows for interaction with the Azure Resource Manager Service `resources` (API Version `2023-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/resources/2023-04-01/policydefinitions"
```


### Client Initialization

```go
client := policydefinitions.NewPolicyDefinitionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PolicyDefinitionsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := policydefinitions.NewProviderPolicyDefinitionID("12345678-1234-9876-4563-123456789012", "policyDefinitionValue")

payload := policydefinitions.PolicyDefinition{
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


### Example Usage: `PolicyDefinitionsClient.CreateOrUpdateAtManagementGroup`

```go
ctx := context.TODO()
id := policydefinitions.NewProviders2PolicyDefinitionID("managementGroupValue", "policyDefinitionValue")

payload := policydefinitions.PolicyDefinition{
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


### Example Usage: `PolicyDefinitionsClient.Delete`

```go
ctx := context.TODO()
id := policydefinitions.NewProviderPolicyDefinitionID("12345678-1234-9876-4563-123456789012", "policyDefinitionValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyDefinitionsClient.DeleteAtManagementGroup`

```go
ctx := context.TODO()
id := policydefinitions.NewProviders2PolicyDefinitionID("managementGroupValue", "policyDefinitionValue")

read, err := client.DeleteAtManagementGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyDefinitionsClient.Get`

```go
ctx := context.TODO()
id := policydefinitions.NewProviderPolicyDefinitionID("12345678-1234-9876-4563-123456789012", "policyDefinitionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyDefinitionsClient.GetAtManagementGroup`

```go
ctx := context.TODO()
id := policydefinitions.NewProviders2PolicyDefinitionID("managementGroupValue", "policyDefinitionValue")

read, err := client.GetAtManagementGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyDefinitionsClient.GetBuiltIn`

```go
ctx := context.TODO()
id := policydefinitions.NewPolicyDefinitionID("policyDefinitionValue")

read, err := client.GetBuiltIn(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyDefinitionsClient.List`

```go
ctx := context.TODO()
id := policydefinitions.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id, policydefinitions.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, policydefinitions.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicyDefinitionsClient.ListBuiltIn`

```go
ctx := context.TODO()


// alternatively `client.ListBuiltIn(ctx, policydefinitions.DefaultListBuiltInOperationOptions())` can be used to do batched pagination
items, err := client.ListBuiltInComplete(ctx, policydefinitions.DefaultListBuiltInOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicyDefinitionsClient.ListByManagementGroup`

```go
ctx := context.TODO()
id := policydefinitions.NewManagementGroupID("groupIdValue")

// alternatively `client.ListByManagementGroup(ctx, id, policydefinitions.DefaultListByManagementGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListByManagementGroupComplete(ctx, id, policydefinitions.DefaultListByManagementGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
