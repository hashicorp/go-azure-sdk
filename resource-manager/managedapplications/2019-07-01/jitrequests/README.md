
## `github.com/hashicorp/go-azure-sdk/resource-manager/managedapplications/2019-07-01/jitrequests` Documentation

The `jitrequests` SDK allows for interaction with Azure Resource Manager `managedapplications` (API Version `2019-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/managedapplications/2019-07-01/jitrequests"
```


### Client Initialization

```go
client := jitrequests.NewJitRequestsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JitRequestsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := jitrequests.NewJitRequestID("12345678-1234-9876-4563-123456789012", "example-resource-group", "jitRequestName")

payload := jitrequests.JitRequestDefinition{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `JitRequestsClient.Delete`

```go
ctx := context.TODO()
id := jitrequests.NewJitRequestID("12345678-1234-9876-4563-123456789012", "example-resource-group", "jitRequestName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JitRequestsClient.Get`

```go
ctx := context.TODO()
id := jitrequests.NewJitRequestID("12345678-1234-9876-4563-123456789012", "example-resource-group", "jitRequestName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JitRequestsClient.ListByResourceGroup`

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


### Example Usage: `JitRequestsClient.ListBySubscription`

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


### Example Usage: `JitRequestsClient.Update`

```go
ctx := context.TODO()
id := jitrequests.NewJitRequestID("12345678-1234-9876-4563-123456789012", "example-resource-group", "jitRequestName")

payload := jitrequests.JitRequestPatchable{
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
