
## `github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/baremetalmachinekeysets` Documentation

The `baremetalmachinekeysets` SDK allows for interaction with Azure Resource Manager `networkcloud` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/baremetalmachinekeysets"
```


### Client Initialization

```go
client := baremetalmachinekeysets.NewBareMetalMachineKeySetsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BareMetalMachineKeySetsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := baremetalmachinekeysets.NewBareMetalMachineKeySetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "bareMetalMachineKeySetName")

payload := baremetalmachinekeysets.BareMetalMachineKeySet{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload, baremetalmachinekeysets.DefaultCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `BareMetalMachineKeySetsClient.Delete`

```go
ctx := context.TODO()
id := baremetalmachinekeysets.NewBareMetalMachineKeySetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "bareMetalMachineKeySetName")

if err := client.DeleteThenPoll(ctx, id, baremetalmachinekeysets.DefaultDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `BareMetalMachineKeySetsClient.Get`

```go
ctx := context.TODO()
id := baremetalmachinekeysets.NewBareMetalMachineKeySetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "bareMetalMachineKeySetName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BareMetalMachineKeySetsClient.ListByCluster`

```go
ctx := context.TODO()
id := baremetalmachinekeysets.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

// alternatively `client.ListByCluster(ctx, id, baremetalmachinekeysets.DefaultListByClusterOperationOptions())` can be used to do batched pagination
items, err := client.ListByClusterComplete(ctx, id, baremetalmachinekeysets.DefaultListByClusterOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BareMetalMachineKeySetsClient.Update`

```go
ctx := context.TODO()
id := baremetalmachinekeysets.NewBareMetalMachineKeySetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "bareMetalMachineKeySetName")

payload := baremetalmachinekeysets.BareMetalMachineKeySetPatchParameters{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload, baremetalmachinekeysets.DefaultUpdateOperationOptions()); err != nil {
	// handle the error
}
```
