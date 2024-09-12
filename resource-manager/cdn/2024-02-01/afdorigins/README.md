
## `github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2024-02-01/afdorigins` Documentation

The `afdorigins` SDK allows for interaction with the Azure Resource Manager Service `cdn` (API Version `2024-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2024-02-01/afdorigins"
```


### Client Initialization

```go
client := afdorigins.NewAFDOriginsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AFDOriginsClient.Create`

```go
ctx := context.TODO()
id := afdorigins.NewOriginGroupOriginID("12345678-1234-9876-4563-123456789012", "example-resource-group", "profileValue", "originGroupValue", "originValue")

payload := afdorigins.AFDOrigin{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AFDOriginsClient.Delete`

```go
ctx := context.TODO()
id := afdorigins.NewOriginGroupOriginID("12345678-1234-9876-4563-123456789012", "example-resource-group", "profileValue", "originGroupValue", "originValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `AFDOriginsClient.Get`

```go
ctx := context.TODO()
id := afdorigins.NewOriginGroupOriginID("12345678-1234-9876-4563-123456789012", "example-resource-group", "profileValue", "originGroupValue", "originValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AFDOriginsClient.ListByOriginGroup`

```go
ctx := context.TODO()
id := afdorigins.NewOriginGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "profileValue", "originGroupValue")

// alternatively `client.ListByOriginGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByOriginGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AFDOriginsClient.Update`

```go
ctx := context.TODO()
id := afdorigins.NewOriginGroupOriginID("12345678-1234-9876-4563-123456789012", "example-resource-group", "profileValue", "originGroupValue", "originValue")

payload := afdorigins.AFDOriginUpdateParameters{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
