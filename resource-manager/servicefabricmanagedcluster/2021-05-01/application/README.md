
## `github.com/hashicorp/go-azure-sdk/resource-manager/servicefabricmanagedcluster/2021-05-01/application` Documentation

The `application` SDK allows for interaction with Azure Resource Manager `servicefabricmanagedcluster` (API Version `2021-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/servicefabricmanagedcluster/2021-05-01/application"
```


### Client Initialization

```go
client := application.NewApplicationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ApplicationClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := application.NewApplicationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "applicationName")

payload := application.ApplicationResource{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ApplicationClient.Delete`

```go
ctx := context.TODO()
id := application.NewApplicationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "applicationName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ApplicationClient.Get`

```go
ctx := context.TODO()
id := application.NewApplicationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "applicationName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.List`

```go
ctx := context.TODO()
id := application.NewManagedClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ApplicationClient.Update`

```go
ctx := context.TODO()
id := application.NewApplicationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "applicationName")

payload := application.ApplicationUpdateParameters{
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
