
## `github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-09-01/iscsipaths` Documentation

The `iscsipaths` SDK allows for interaction with Azure Resource Manager `vmware` (API Version `2023-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-09-01/iscsipaths"
```


### Client Initialization

```go
client := iscsipaths.NewIscsiPathsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IscsiPathsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := iscsipaths.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue")

payload := iscsipaths.IscsiPath{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `IscsiPathsClient.Delete`

```go
ctx := context.TODO()
id := iscsipaths.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `IscsiPathsClient.Get`

```go
ctx := context.TODO()
id := iscsipaths.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IscsiPathsClient.ListByPrivateCloud`

```go
ctx := context.TODO()
id := iscsipaths.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue")

// alternatively `client.ListByPrivateCloud(ctx, id)` can be used to do batched pagination
items, err := client.ListByPrivateCloudComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
