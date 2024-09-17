
## `github.com/hashicorp/go-azure-sdk/resource-manager/servicelinker/2024-04-01/linkers` Documentation

The `linkers` SDK allows for interaction with Azure Resource Manager `servicelinker` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/servicelinker/2024-04-01/linkers"
```


### Client Initialization

```go
client := linkers.NewLinkersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LinkersClient.CreateDryrun`

```go
ctx := context.TODO()
id := linkers.NewScopedDryrunID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "dryrunValue")

payload := linkers.DryrunResource{
	// ...
}


if err := client.CreateDryrunThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `LinkersClient.DeleteDryrun`

```go
ctx := context.TODO()
id := linkers.NewScopedDryrunID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "dryrunValue")

read, err := client.DeleteDryrun(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LinkersClient.GenerateConfigurations`

```go
ctx := context.TODO()
id := linkers.NewScopedLinkerID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "linkerValue")

payload := linkers.ConfigurationInfo{
	// ...
}


read, err := client.GenerateConfigurations(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LinkersClient.GetDryrun`

```go
ctx := context.TODO()
id := linkers.NewScopedDryrunID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "dryrunValue")

read, err := client.GetDryrun(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LinkersClient.LinkerDelete`

```go
ctx := context.TODO()
id := linkers.NewScopedLinkerID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "linkerValue")

if err := client.LinkerDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `LinkersClient.LinkerListConfigurations`

```go
ctx := context.TODO()
id := linkers.NewScopedLinkerID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "linkerValue")

read, err := client.LinkerListConfigurations(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LinkersClient.LinkerUpdate`

```go
ctx := context.TODO()
id := linkers.NewScopedLinkerID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "linkerValue")

payload := linkers.LinkerPatch{
	// ...
}


if err := client.LinkerUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `LinkersClient.LinkerValidate`

```go
ctx := context.TODO()
id := linkers.NewScopedLinkerID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "linkerValue")

if err := client.LinkerValidateThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `LinkersClient.ListDaprConfigurations`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.ListDaprConfigurations(ctx, id)` can be used to do batched pagination
items, err := client.ListDaprConfigurationsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `LinkersClient.ListDryrun`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.ListDryrun(ctx, id)` can be used to do batched pagination
items, err := client.ListDryrunComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `LinkersClient.UpdateDryrun`

```go
ctx := context.TODO()
id := linkers.NewScopedDryrunID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "dryrunValue")

payload := linkers.DryrunPatch{
	// ...
}


if err := client.UpdateDryrunThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
