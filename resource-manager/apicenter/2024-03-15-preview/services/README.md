
## `github.com/hashicorp/go-azure-sdk/resource-manager/apicenter/2024-03-15-preview/services` Documentation

The `services` SDK allows for interaction with the Azure Resource Manager Service `apicenter` (API Version `2024-03-15-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/apicenter/2024-03-15-preview/services"
```


### Client Initialization

```go
client := services.NewServicesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ServicesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := services.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue")

payload := services.Service{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicesClient.Delete`

```go
ctx := context.TODO()
id := services.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicesClient.ExportMetadataSchema`

```go
ctx := context.TODO()
id := services.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue")

payload := services.MetadataSchemaExportRequest{
	// ...
}


if err := client.ExportMetadataSchemaThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ServicesClient.Get`

```go
ctx := context.TODO()
id := services.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicesClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ServicesClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ServicesClient.Update`

```go
ctx := context.TODO()
id := services.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue")

payload := services.ServiceUpdate{
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
