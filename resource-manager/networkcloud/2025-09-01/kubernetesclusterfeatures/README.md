
## `github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/kubernetesclusterfeatures` Documentation

The `kubernetesclusterfeatures` SDK allows for interaction with Azure Resource Manager `networkcloud` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/kubernetesclusterfeatures"
```


### Client Initialization

```go
client := kubernetesclusterfeatures.NewKubernetesClusterFeaturesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `KubernetesClusterFeaturesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := kubernetesclusterfeatures.NewFeatureID("12345678-1234-9876-4563-123456789012", "example-resource-group", "kubernetesClusterName", "featureName")

payload := kubernetesclusterfeatures.KubernetesClusterFeature{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload, kubernetesclusterfeatures.DefaultCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `KubernetesClusterFeaturesClient.Delete`

```go
ctx := context.TODO()
id := kubernetesclusterfeatures.NewFeatureID("12345678-1234-9876-4563-123456789012", "example-resource-group", "kubernetesClusterName", "featureName")

if err := client.DeleteThenPoll(ctx, id, kubernetesclusterfeatures.DefaultDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `KubernetesClusterFeaturesClient.Get`

```go
ctx := context.TODO()
id := kubernetesclusterfeatures.NewFeatureID("12345678-1234-9876-4563-123456789012", "example-resource-group", "kubernetesClusterName", "featureName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `KubernetesClusterFeaturesClient.ListByKubernetesCluster`

```go
ctx := context.TODO()
id := kubernetesclusterfeatures.NewKubernetesClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "kubernetesClusterName")

// alternatively `client.ListByKubernetesCluster(ctx, id, kubernetesclusterfeatures.DefaultListByKubernetesClusterOperationOptions())` can be used to do batched pagination
items, err := client.ListByKubernetesClusterComplete(ctx, id, kubernetesclusterfeatures.DefaultListByKubernetesClusterOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `KubernetesClusterFeaturesClient.Update`

```go
ctx := context.TODO()
id := kubernetesclusterfeatures.NewFeatureID("12345678-1234-9876-4563-123456789012", "example-resource-group", "kubernetesClusterName", "featureName")

payload := kubernetesclusterfeatures.KubernetesClusterFeaturePatchParameters{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload, kubernetesclusterfeatures.DefaultUpdateOperationOptions()); err != nil {
	// handle the error
}
```
