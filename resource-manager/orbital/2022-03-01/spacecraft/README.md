
## `github.com/hashicorp/go-azure-sdk/resource-manager/orbital/2022-03-01/spacecraft` Documentation

The `spacecraft` SDK allows for interaction with the Azure Resource Manager Service `orbital` (API Version `2022-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/orbital/2022-03-01/spacecraft"
```


### Client Initialization

```go
client := spacecraft.NewSpacecraftClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `SpacecraftClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := spacecraft.NewSpacecraftID("12345678-1234-9876-4563-123456789012", "example-resource-group", "spacecraftValue")

payload := spacecraft.Spacecraft{
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


### Example Usage: `SpacecraftClient.Delete`

```go
ctx := context.TODO()
id := spacecraft.NewSpacecraftID("12345678-1234-9876-4563-123456789012", "example-resource-group", "spacecraftValue")
future, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `SpacecraftClient.Get`

```go
ctx := context.TODO()
id := spacecraft.NewSpacecraftID("12345678-1234-9876-4563-123456789012", "example-resource-group", "spacecraftValue")
read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SpacecraftClient.List`

```go
ctx := context.TODO()
id := spacecraft.NewResourceGroupID()
// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SpacecraftClient.ListBySubscription`

```go
ctx := context.TODO()
id := spacecraft.NewSubscriptionID()
// alternatively `client.ListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SpacecraftClient.UpdateTags`

```go
ctx := context.TODO()
id := spacecraft.NewSpacecraftID("12345678-1234-9876-4563-123456789012", "example-resource-group", "spacecraftValue")

payload := spacecraft.TagsObject{
	// ...
}

future, err := client.UpdateTags(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```
