
## `github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2022-01-10-preview/vcenters` Documentation

The `vcenters` SDK allows for interaction with Azure Resource Manager `connectedvmware` (API Version `2022-01-10-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2022-01-10-preview/vcenters"
```


### Client Initialization

```go
client := vcenters.NewVCentersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VCentersClient.Create`

```go
ctx := context.TODO()
id := vcenters.NewVCenterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vCenterName")

payload := vcenters.VCenter{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VCentersClient.Delete`

```go
ctx := context.TODO()
id := vcenters.NewVCenterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vCenterName")

if err := client.DeleteThenPoll(ctx, id, vcenters.DefaultDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `VCentersClient.Get`

```go
ctx := context.TODO()
id := vcenters.NewVCenterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vCenterName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VCentersClient.List`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VCentersClient.ListByResourceGroup`

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


### Example Usage: `VCentersClient.Update`

```go
ctx := context.TODO()
id := vcenters.NewVCenterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vCenterName")

payload := vcenters.ResourcePatch{
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
