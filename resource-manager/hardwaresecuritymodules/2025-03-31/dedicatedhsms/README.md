
## `github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/dedicatedhsms` Documentation

The `dedicatedhsms` SDK allows for interaction with Azure Resource Manager `hardwaresecuritymodules` (API Version `2025-03-31`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/dedicatedhsms"
```


### Client Initialization

```go
client := dedicatedhsms.NewDedicatedHsmsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DedicatedHsmsClient.DedicatedHsmCreateOrUpdate`

```go
ctx := context.TODO()
id := dedicatedhsms.NewDedicatedHSMID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dedicatedHSMName")

payload := dedicatedhsms.DedicatedHsm{
	// ...
}


if err := client.DedicatedHsmCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DedicatedHsmsClient.DedicatedHsmDelete`

```go
ctx := context.TODO()
id := dedicatedhsms.NewDedicatedHSMID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dedicatedHSMName")

if err := client.DedicatedHsmDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DedicatedHsmsClient.DedicatedHsmGet`

```go
ctx := context.TODO()
id := dedicatedhsms.NewDedicatedHSMID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dedicatedHSMName")

read, err := client.DedicatedHsmGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DedicatedHsmsClient.DedicatedHsmListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.DedicatedHsmListByResourceGroup(ctx, id, dedicatedhsms.DefaultDedicatedHsmListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.DedicatedHsmListByResourceGroupComplete(ctx, id, dedicatedhsms.DefaultDedicatedHsmListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DedicatedHsmsClient.DedicatedHsmListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.DedicatedHsmListBySubscription(ctx, id, dedicatedhsms.DefaultDedicatedHsmListBySubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.DedicatedHsmListBySubscriptionComplete(ctx, id, dedicatedhsms.DefaultDedicatedHsmListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DedicatedHsmsClient.DedicatedHsmListOutboundNetworkDependenciesEndpoints`

```go
ctx := context.TODO()
id := dedicatedhsms.NewDedicatedHSMID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dedicatedHSMName")

// alternatively `client.DedicatedHsmListOutboundNetworkDependenciesEndpoints(ctx, id)` can be used to do batched pagination
items, err := client.DedicatedHsmListOutboundNetworkDependenciesEndpointsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DedicatedHsmsClient.DedicatedHsmUpdate`

```go
ctx := context.TODO()
id := dedicatedhsms.NewDedicatedHSMID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dedicatedHSMName")

payload := dedicatedhsms.DedicatedHsmPatchParameters{
	// ...
}


if err := client.DedicatedHsmUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
