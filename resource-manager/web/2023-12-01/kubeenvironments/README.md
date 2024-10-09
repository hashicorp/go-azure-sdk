
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2023-12-01/kubeenvironments` Documentation

The `kubeenvironments` SDK allows for interaction with Azure Resource Manager `web` (API Version `2023-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2023-12-01/kubeenvironments"
```


### Client Initialization

```go
client := kubeenvironments.NewKubeEnvironmentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `KubeEnvironmentsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := kubeenvironments.NewKubeEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "kubeEnvironmentName")

payload := kubeenvironments.KubeEnvironment{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `KubeEnvironmentsClient.Delete`

```go
ctx := context.TODO()
id := kubeenvironments.NewKubeEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "kubeEnvironmentName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `KubeEnvironmentsClient.Get`

```go
ctx := context.TODO()
id := kubeenvironments.NewKubeEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "kubeEnvironmentName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `KubeEnvironmentsClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `KubeEnvironmentsClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `KubeEnvironmentsClient.Update`

```go
ctx := context.TODO()
id := kubeenvironments.NewKubeEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "kubeEnvironmentName")

payload := kubeenvironments.KubeEnvironmentPatchResource{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
