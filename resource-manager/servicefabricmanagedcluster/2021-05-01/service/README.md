
## `github.com/hashicorp/go-azure-sdk/resource-manager/servicefabricmanagedcluster/2021-05-01/service` Documentation

The `service` SDK allows for interaction with the Azure Resource Manager Service `servicefabricmanagedcluster` (API Version `2021-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/servicefabricmanagedcluster/2021-05-01/service"
```


### Client Initialization

```go
client := service.NewServiceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `ServiceClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := service.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue", "applicationValue", "serviceValue")

payload := service.ServiceResource{
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


### Example Usage: `ServiceClient.Delete`

```go
ctx := context.TODO()
id := service.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue", "applicationValue", "serviceValue")
future, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `ServiceClient.Get`

```go
ctx := context.TODO()
id := service.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue", "applicationValue", "serviceValue")
read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServiceClient.ListByApplications`

```go
ctx := context.TODO()
id := service.NewApplicationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue", "applicationValue")
// alternatively `client.ListByApplications(ctx, id)` can be used to do batched pagination
items, err := client.ListByApplicationsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
