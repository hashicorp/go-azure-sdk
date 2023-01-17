
## `github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/disks` Documentation

The `disks` SDK allows for interaction with the Azure Resource Manager Service `devtestlab` (API Version `2018-09-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/disks"
```


### Client Initialization

```go
client := disks.NewDisksClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DisksClient.Attach`

```go
ctx := context.TODO()
id := disks.NewDiskID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "userValue", "nameValue")

payload := disks.AttachDiskProperties{
	// ...
}


if err := client.AttachThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DisksClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := disks.NewDiskID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "userValue", "nameValue")

payload := disks.Disk{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DisksClient.Delete`

```go
ctx := context.TODO()
id := disks.NewDiskID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "userValue", "nameValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DisksClient.Detach`

```go
ctx := context.TODO()
id := disks.NewDiskID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "userValue", "nameValue")

payload := disks.DetachDiskProperties{
	// ...
}


if err := client.DetachThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DisksClient.Get`

```go
ctx := context.TODO()
id := disks.NewDiskID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "userValue", "nameValue")

read, err := client.Get(ctx, id, disks.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DisksClient.List`

```go
ctx := context.TODO()
id := disks.NewUserID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "nameValue")

// alternatively `client.List(ctx, id, disks.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, disks.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DisksClient.Update`

```go
ctx := context.TODO()
id := disks.NewDiskID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "userValue", "nameValue")

payload := disks.UpdateResource{
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
