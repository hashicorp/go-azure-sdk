
## `github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-06-01-preview/dicomservices` Documentation

The `dicomservices` SDK allows for interaction with the Azure Resource Manager Service `healthcareapis` (API Version `2021-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-06-01-preview/dicomservices"
```


### Client Initialization

```go
client := dicomservices.NewDicomServicesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `DicomServicesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := dicomservices.NewDicomServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "dicomServiceValue")

payload := dicomservices.DicomService{
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


### Example Usage: `DicomServicesClient.Delete`

```go
ctx := context.TODO()
id := dicomservices.NewDicomServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "dicomServiceValue")
future, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `DicomServicesClient.Get`

```go
ctx := context.TODO()
id := dicomservices.NewDicomServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "dicomServiceValue")
read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DicomServicesClient.ListByWorkspace`

```go
ctx := context.TODO()
id := dicomservices.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")
// alternatively `client.ListByWorkspace(ctx, id)` can be used to do batched pagination
items, err := client.ListByWorkspaceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DicomServicesClient.Update`

```go
ctx := context.TODO()
id := dicomservices.NewDicomServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "dicomServiceValue")

payload := dicomservices.ResourceTags{
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
