
## `github.com/hashicorp/go-azure-sdk/resource-manager/resources/2023-04-01/policydefinitionversions` Documentation

The `policydefinitionversions` SDK allows for interaction with the Azure Resource Manager Service `resources` (API Version `2023-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/resources/2023-04-01/policydefinitionversions"
```


### Client Initialization

```go
client := policydefinitionversions.NewPolicyDefinitionVersionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PolicyDefinitionVersionsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := policydefinitionversions.NewPolicyDefinitionVersionID("12345678-1234-9876-4563-123456789012", "policyDefinitionValue", "versionValue")

payload := policydefinitionversions.PolicyDefinitionVersion{
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


### Example Usage: `PolicyDefinitionVersionsClient.CreateOrUpdateAtManagementGroup`

```go
ctx := context.TODO()
id := policydefinitionversions.NewProviders2PolicyDefinitionVersionID("managementGroupValue", "policyDefinitionValue", "versionValue")

payload := policydefinitionversions.PolicyDefinitionVersion{
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


### Example Usage: `PolicyDefinitionVersionsClient.Delete`

```go
ctx := context.TODO()
id := policydefinitionversions.NewPolicyDefinitionVersionID("12345678-1234-9876-4563-123456789012", "policyDefinitionValue", "versionValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyDefinitionVersionsClient.DeleteAtManagementGroup`

```go
ctx := context.TODO()
id := policydefinitionversions.NewProviders2PolicyDefinitionVersionID("managementGroupValue", "policyDefinitionValue", "versionValue")

read, err := client.DeleteAtManagementGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyDefinitionVersionsClient.Get`

```go
ctx := context.TODO()
id := policydefinitionversions.NewPolicyDefinitionVersionID("12345678-1234-9876-4563-123456789012", "policyDefinitionValue", "versionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyDefinitionVersionsClient.GetAtManagementGroup`

```go
ctx := context.TODO()
id := policydefinitionversions.NewProviders2PolicyDefinitionVersionID("managementGroupValue", "policyDefinitionValue", "versionValue")

read, err := client.GetAtManagementGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyDefinitionVersionsClient.GetBuiltIn`

```go
ctx := context.TODO()
id := policydefinitionversions.NewVersionID("policyDefinitionValue", "versionValue")

read, err := client.GetBuiltIn(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyDefinitionVersionsClient.List`

```go
ctx := context.TODO()
id := policydefinitionversions.NewProviderPolicyDefinitionID("12345678-1234-9876-4563-123456789012", "policyDefinitionValue")

// alternatively `client.List(ctx, id, policydefinitionversions.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, policydefinitionversions.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicyDefinitionVersionsClient.ListAll`

```go
ctx := context.TODO()
id := policydefinitionversions.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.ListAll(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyDefinitionVersionsClient.ListAllAtManagementGroup`

```go
ctx := context.TODO()
id := policydefinitionversions.NewManagementGroupID("groupIdValue")

read, err := client.ListAllAtManagementGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyDefinitionVersionsClient.ListAllBuiltins`

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


### Example Usage: `PolicyDefinitionVersionsClient.ListBuiltIn`

```go
ctx := context.TODO()
id := policydefinitionversions.NewPolicyDefinitionID("policyDefinitionValue")

// alternatively `client.ListBuiltIn(ctx, id, policydefinitionversions.DefaultListBuiltInOperationOptions())` can be used to do batched pagination
items, err := client.ListBuiltInComplete(ctx, id, policydefinitionversions.DefaultListBuiltInOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicyDefinitionVersionsClient.ListByManagementGroup`

```go
ctx := context.TODO()
id := policydefinitionversions.NewProviders2PolicyDefinitionID("managementGroupValue", "policyDefinitionValue")

// alternatively `client.ListByManagementGroup(ctx, id, policydefinitionversions.DefaultListByManagementGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListByManagementGroupComplete(ctx, id, policydefinitionversions.DefaultListByManagementGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
