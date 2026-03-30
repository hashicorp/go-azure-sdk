
## `github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/bmckeysets` Documentation

The `bmckeysets` SDK allows for interaction with Azure Resource Manager `networkcloud` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/bmckeysets"
```


### Client Initialization

```go
client := bmckeysets.NewBmcKeySetsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BmcKeySetsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := bmckeysets.NewBmcKeySetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "bmcKeySetName")

payload := bmckeysets.BmcKeySet{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload, bmckeysets.DefaultCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `BmcKeySetsClient.Delete`

```go
ctx := context.TODO()
id := bmckeysets.NewBmcKeySetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "bmcKeySetName")

if err := client.DeleteThenPoll(ctx, id, bmckeysets.DefaultDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `BmcKeySetsClient.Get`

```go
ctx := context.TODO()
id := bmckeysets.NewBmcKeySetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "bmcKeySetName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BmcKeySetsClient.ListByCluster`

```go
ctx := context.TODO()
id := bmckeysets.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

// alternatively `client.ListByCluster(ctx, id, bmckeysets.DefaultListByClusterOperationOptions())` can be used to do batched pagination
items, err := client.ListByClusterComplete(ctx, id, bmckeysets.DefaultListByClusterOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BmcKeySetsClient.Update`

```go
ctx := context.TODO()
id := bmckeysets.NewBmcKeySetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "bmcKeySetName")

payload := bmckeysets.BmcKeySetPatchParameters{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload, bmckeysets.DefaultUpdateOperationOptions()); err != nil {
	// handle the error
}
```
