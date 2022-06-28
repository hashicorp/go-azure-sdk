
## `github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2021-10-01-preview/lab` Documentation

The `lab` SDK allows for interaction with the Azure Resource Manager Service `labservices` (API Version `2021-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2021-10-01-preview/lab"
```


### Client Initialization

```go
client := lab.NewLabClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `LabClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := lab.NewLabID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue")

payload := lab.Lab{
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


### Example Usage: `LabClient.Delete`

```go
ctx := context.TODO()
id := lab.NewLabID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue")
future, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `LabClient.Get`

```go
ctx := context.TODO()
id := lab.NewLabID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue")
read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LabClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := lab.NewResourceGroupID()
// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `LabClient.ListBySubscription`

```go
ctx := context.TODO()
id := lab.NewSubscriptionID()
// alternatively `client.ListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `LabClient.Publish`

```go
ctx := context.TODO()
id := lab.NewLabID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue")
future, err := client.Publish(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `LabClient.SyncGroup`

```go
ctx := context.TODO()
id := lab.NewLabID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue")
future, err := client.SyncGroup(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `LabClient.Update`

```go
ctx := context.TODO()
id := lab.NewLabID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue")

payload := lab.LabUpdate{
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
