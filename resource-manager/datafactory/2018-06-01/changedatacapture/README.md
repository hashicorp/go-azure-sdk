
## `github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/changedatacapture` Documentation

The `changedatacapture` SDK allows for interaction with the Azure Resource Manager Service `datafactory` (API Version `2018-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/changedatacapture"
```


### Client Initialization

```go
client := changedatacapture.NewChangeDataCaptureClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ChangeDataCaptureClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := changedatacapture.NewAdfcdcID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue", "adfcdcValue")

payload := changedatacapture.ChangeDataCaptureResource{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload, changedatacapture.DefaultCreateOrUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChangeDataCaptureClient.Delete`

```go
ctx := context.TODO()
id := changedatacapture.NewAdfcdcID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue", "adfcdcValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChangeDataCaptureClient.Get`

```go
ctx := context.TODO()
id := changedatacapture.NewAdfcdcID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue", "adfcdcValue")

read, err := client.Get(ctx, id, changedatacapture.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChangeDataCaptureClient.ListByFactory`

```go
ctx := context.TODO()
id := changedatacapture.NewFactoryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue")

// alternatively `client.ListByFactory(ctx, id)` can be used to do batched pagination
items, err := client.ListByFactoryComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ChangeDataCaptureClient.Start`

```go
ctx := context.TODO()
id := changedatacapture.NewAdfcdcID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue", "adfcdcValue")

read, err := client.Start(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChangeDataCaptureClient.Status`

```go
ctx := context.TODO()
id := changedatacapture.NewAdfcdcID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue", "adfcdcValue")

read, err := client.Status(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChangeDataCaptureClient.Stop`

```go
ctx := context.TODO()
id := changedatacapture.NewAdfcdcID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue", "adfcdcValue")

read, err := client.Stop(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
