
## `github.com/hashicorp/go-azure-sdk/resource-manager/systemcentervirtualmachinemanager/2025-03-13/guestagents` Documentation

The `guestagents` SDK allows for interaction with Azure Resource Manager `systemcentervirtualmachinemanager` (API Version `2025-03-13`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/systemcentervirtualmachinemanager/2025-03-13/guestagents"
```


### Client Initialization

```go
client := guestagents.NewGuestAgentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GuestAgentsClient.Create`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

payload := guestagents.GuestAgent{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `GuestAgentsClient.Delete`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GuestAgentsClient.Get`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GuestAgentsClient.ListByVirtualMachineInstance`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.ListByVirtualMachineInstance(ctx, id)` can be used to do batched pagination
items, err := client.ListByVirtualMachineInstanceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
