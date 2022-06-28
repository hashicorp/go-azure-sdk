
## `github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2021-10-01-preview/labplan` Documentation

The `labplan` SDK allows for interaction with the Azure Resource Manager Service `labservices` (API Version `2021-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2021-10-01-preview/labplan"
```


### Client Initialization

```go
client := labplan.NewLabPlanClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `LabPlanClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := labplan.NewLabPlanID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labPlanValue")

payload := labplan.LabPlan{
	// ...
}

future, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `LabPlanClient.Delete`

```go
ctx := context.TODO()
id := labplan.NewLabPlanID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labPlanValue")
future, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `LabPlanClient.Get`

```go
ctx := context.TODO()
id := labplan.NewLabPlanID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labPlanValue")
read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LabPlanClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := labplan.NewResourceGroupID()
// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `LabPlanClient.ListBySubscription`

```go
ctx := context.TODO()
id := labplan.NewSubscriptionID()
// alternatively `client.ListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `LabPlanClient.Update`

```go
ctx := context.TODO()
id := labplan.NewLabPlanID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labPlanValue")

payload := labplan.LabPlanUpdate{
	// ...
}

future, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```
