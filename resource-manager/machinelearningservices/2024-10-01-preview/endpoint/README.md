
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/endpoint` Documentation

The `endpoint` SDK allows for interaction with Azure Resource Manager `machinelearningservices` (API Version `2024-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/endpoint"
```


### Client Initialization

```go
client := endpoint.NewEndpointClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EndpointClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := endpoint.NewEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "endpointName")

payload := endpoint.EndpointResourcePropertiesBasicResource{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `EndpointClient.DeploymentCreateOrUpdate`

```go
ctx := context.TODO()
id := endpoint.NewDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "endpointName", "deploymentName")

payload := endpoint.EndpointDeploymentResourcePropertiesBasicResource{
	// ...
}


if err := client.DeploymentCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `EndpointClient.DeploymentDelete`

```go
ctx := context.TODO()
id := endpoint.NewDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "endpointName", "deploymentName")

if err := client.DeploymentDeleteThenPoll(ctx, id, endpoint.DefaultDeploymentDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `EndpointClient.DeploymentGet`

```go
ctx := context.TODO()
id := endpoint.NewDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "endpointName", "deploymentName")

read, err := client.DeploymentGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EndpointClient.DeploymentGetInWorkspace`

```go
ctx := context.TODO()
id := endpoint.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

// alternatively `client.DeploymentGetInWorkspace(ctx, id, endpoint.DefaultDeploymentGetInWorkspaceOperationOptions())` can be used to do batched pagination
items, err := client.DeploymentGetInWorkspaceComplete(ctx, id, endpoint.DefaultDeploymentGetInWorkspaceOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EndpointClient.DeploymentList`

```go
ctx := context.TODO()
id := endpoint.NewEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "endpointName")

// alternatively `client.DeploymentList(ctx, id)` can be used to do batched pagination
items, err := client.DeploymentListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EndpointClient.Get`

```go
ctx := context.TODO()
id := endpoint.NewEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "endpointName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EndpointClient.GetModels`

```go
ctx := context.TODO()
id := endpoint.NewEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "endpointName")

// alternatively `client.GetModels(ctx, id)` can be used to do batched pagination
items, err := client.GetModelsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EndpointClient.List`

```go
ctx := context.TODO()
id := endpoint.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

// alternatively `client.List(ctx, id, endpoint.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, endpoint.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EndpointClient.ListKeys`

```go
ctx := context.TODO()
id := endpoint.NewEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "endpointName")

read, err := client.ListKeys(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EndpointClient.RaiPoliciesList`

```go
ctx := context.TODO()
id := endpoint.NewEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "endpointName")

// alternatively `client.RaiPoliciesList(ctx, id, endpoint.DefaultRaiPoliciesListOperationOptions())` can be used to do batched pagination
items, err := client.RaiPoliciesListComplete(ctx, id, endpoint.DefaultRaiPoliciesListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EndpointClient.RaiPolicyCreate`

```go
ctx := context.TODO()
id := endpoint.NewRaiPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "endpointName", "raiPolicyName")

payload := endpoint.RaiPolicyPropertiesBasicResource{
	// ...
}


if err := client.RaiPolicyCreateThenPoll(ctx, id, payload, endpoint.DefaultRaiPolicyCreateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `EndpointClient.RaiPolicyDelete`

```go
ctx := context.TODO()
id := endpoint.NewRaiPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "endpointName", "raiPolicyName")

if err := client.RaiPolicyDeleteThenPoll(ctx, id, endpoint.DefaultRaiPolicyDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `EndpointClient.RaiPolicyGet`

```go
ctx := context.TODO()
id := endpoint.NewRaiPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "endpointName", "raiPolicyName")

read, err := client.RaiPolicyGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EndpointClient.RegenerateKeys`

```go
ctx := context.TODO()
id := endpoint.NewEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "endpointName")

payload := endpoint.RegenerateServiceAccountKeyContent{
	// ...
}


read, err := client.RegenerateKeys(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
