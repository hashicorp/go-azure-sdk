
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/daprcomponentresiliencypolicies` Documentation

The `daprcomponentresiliencypolicies` SDK allows for interaction with Azure Resource Manager `containerapps` (API Version `2024-02-02-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/daprcomponentresiliencypolicies"
```


### Client Initialization

```go
client := daprcomponentresiliencypolicies.NewDaprComponentResiliencyPoliciesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DaprComponentResiliencyPoliciesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := daprcomponentresiliencypolicies.NewDaprComponentResiliencyPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "environmentName", "componentName", "name")

payload := daprcomponentresiliencypolicies.DaprComponentResiliencyPolicy{
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


### Example Usage: `DaprComponentResiliencyPoliciesClient.Delete`

```go
ctx := context.TODO()
id := daprcomponentresiliencypolicies.NewDaprComponentResiliencyPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "environmentName", "componentName", "name")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DaprComponentResiliencyPoliciesClient.Get`

```go
ctx := context.TODO()
id := daprcomponentresiliencypolicies.NewDaprComponentResiliencyPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "environmentName", "componentName", "name")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DaprComponentResiliencyPoliciesClient.List`

```go
ctx := context.TODO()
id := daprcomponentresiliencypolicies.NewDaprComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "environmentName", "componentName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
