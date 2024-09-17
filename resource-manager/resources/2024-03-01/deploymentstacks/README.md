
## `github.com/hashicorp/go-azure-sdk/resource-manager/resources/2024-03-01/deploymentstacks` Documentation

The `deploymentstacks` SDK allows for interaction with Azure Resource Manager `resources` (API Version `2024-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/resources/2024-03-01/deploymentstacks"
```


### Client Initialization

```go
client := deploymentstacks.NewDeploymentStacksClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeploymentStacksClient.CreateOrUpdateAtManagementGroup`

```go
ctx := context.TODO()
id := deploymentstacks.NewProviders2DeploymentStackID("managementGroupIdValue", "deploymentStackValue")

payload := deploymentstacks.DeploymentStack{
	// ...
}


if err := client.CreateOrUpdateAtManagementGroupThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DeploymentStacksClient.CreateOrUpdateAtResourceGroup`

```go
ctx := context.TODO()
id := deploymentstacks.NewProviderDeploymentStackID("12345678-1234-9876-4563-123456789012", "example-resource-group", "deploymentStackValue")

payload := deploymentstacks.DeploymentStack{
	// ...
}


if err := client.CreateOrUpdateAtResourceGroupThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DeploymentStacksClient.CreateOrUpdateAtSubscription`

```go
ctx := context.TODO()
id := deploymentstacks.NewDeploymentStackID("12345678-1234-9876-4563-123456789012", "deploymentStackValue")

payload := deploymentstacks.DeploymentStack{
	// ...
}


if err := client.CreateOrUpdateAtSubscriptionThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DeploymentStacksClient.DeleteAtManagementGroup`

```go
ctx := context.TODO()
id := deploymentstacks.NewProviders2DeploymentStackID("managementGroupIdValue", "deploymentStackValue")

if err := client.DeleteAtManagementGroupThenPoll(ctx, id, deploymentstacks.DefaultDeleteAtManagementGroupOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `DeploymentStacksClient.DeleteAtResourceGroup`

```go
ctx := context.TODO()
id := deploymentstacks.NewProviderDeploymentStackID("12345678-1234-9876-4563-123456789012", "example-resource-group", "deploymentStackValue")

if err := client.DeleteAtResourceGroupThenPoll(ctx, id, deploymentstacks.DefaultDeleteAtResourceGroupOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `DeploymentStacksClient.DeleteAtSubscription`

```go
ctx := context.TODO()
id := deploymentstacks.NewDeploymentStackID("12345678-1234-9876-4563-123456789012", "deploymentStackValue")

if err := client.DeleteAtSubscriptionThenPoll(ctx, id, deploymentstacks.DefaultDeleteAtSubscriptionOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `DeploymentStacksClient.ExportTemplateAtManagementGroup`

```go
ctx := context.TODO()
id := deploymentstacks.NewProviders2DeploymentStackID("managementGroupIdValue", "deploymentStackValue")

read, err := client.ExportTemplateAtManagementGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeploymentStacksClient.ExportTemplateAtResourceGroup`

```go
ctx := context.TODO()
id := deploymentstacks.NewProviderDeploymentStackID("12345678-1234-9876-4563-123456789012", "example-resource-group", "deploymentStackValue")

read, err := client.ExportTemplateAtResourceGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeploymentStacksClient.ExportTemplateAtSubscription`

```go
ctx := context.TODO()
id := deploymentstacks.NewDeploymentStackID("12345678-1234-9876-4563-123456789012", "deploymentStackValue")

read, err := client.ExportTemplateAtSubscription(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeploymentStacksClient.GetAtManagementGroup`

```go
ctx := context.TODO()
id := deploymentstacks.NewProviders2DeploymentStackID("managementGroupIdValue", "deploymentStackValue")

read, err := client.GetAtManagementGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeploymentStacksClient.GetAtResourceGroup`

```go
ctx := context.TODO()
id := deploymentstacks.NewProviderDeploymentStackID("12345678-1234-9876-4563-123456789012", "example-resource-group", "deploymentStackValue")

read, err := client.GetAtResourceGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeploymentStacksClient.GetAtSubscription`

```go
ctx := context.TODO()
id := deploymentstacks.NewDeploymentStackID("12345678-1234-9876-4563-123456789012", "deploymentStackValue")

read, err := client.GetAtSubscription(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeploymentStacksClient.ListAtManagementGroup`

```go
ctx := context.TODO()
id := commonids.NewManagementGroupID("groupIdValue")

// alternatively `client.ListAtManagementGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListAtManagementGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeploymentStacksClient.ListAtResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListAtResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListAtResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeploymentStacksClient.ListAtSubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListAtSubscription(ctx, id)` can be used to do batched pagination
items, err := client.ListAtSubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeploymentStacksClient.ValidateStackAtManagementGroup`

```go
ctx := context.TODO()
id := deploymentstacks.NewProviders2DeploymentStackID("managementGroupIdValue", "deploymentStackValue")

payload := deploymentstacks.DeploymentStack{
	// ...
}


if err := client.ValidateStackAtManagementGroupThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DeploymentStacksClient.ValidateStackAtResourceGroup`

```go
ctx := context.TODO()
id := deploymentstacks.NewProviderDeploymentStackID("12345678-1234-9876-4563-123456789012", "example-resource-group", "deploymentStackValue")

payload := deploymentstacks.DeploymentStack{
	// ...
}


if err := client.ValidateStackAtResourceGroupThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DeploymentStacksClient.ValidateStackAtSubscription`

```go
ctx := context.TODO()
id := deploymentstacks.NewDeploymentStackID("12345678-1234-9876-4563-123456789012", "deploymentStackValue")

payload := deploymentstacks.DeploymentStack{
	// ...
}


if err := client.ValidateStackAtSubscriptionThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
