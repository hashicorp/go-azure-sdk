
## `github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/changedatacapturedestinations` Documentation

The `changedatacapturedestinations` SDK allows for interaction with Azure Resource Manager `horizondb` (API Version `2026-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/changedatacapturedestinations"
```


### Client Initialization

```go
client := changedatacapturedestinations.NewChangeDataCaptureDestinationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ChangeDataCaptureDestinationsClient.ActivateIdentity`

```go
ctx := context.TODO()
id := changedatacapturedestinations.NewChangeDataCaptureDestinationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "changeDataCaptureDestinationName")

payload := changedatacapturedestinations.ActivateChangeDataCaptureIdentityRequest{
	// ...
}


if err := client.ActivateIdentityThenPoll(ctx, id, payload, changedatacapturedestinations.DefaultActivateIdentityOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `ChangeDataCaptureDestinationsClient.CreateOrReplace`

```go
ctx := context.TODO()
id := changedatacapturedestinations.NewChangeDataCaptureDestinationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "changeDataCaptureDestinationName")

payload := changedatacapturedestinations.ChangeDataCaptureDestination{
	// ...
}


if err := client.CreateOrReplaceThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ChangeDataCaptureDestinationsClient.Delete`

```go
ctx := context.TODO()
id := changedatacapturedestinations.NewChangeDataCaptureDestinationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "changeDataCaptureDestinationName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ChangeDataCaptureDestinationsClient.Get`

```go
ctx := context.TODO()
id := changedatacapturedestinations.NewChangeDataCaptureDestinationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "changeDataCaptureDestinationName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChangeDataCaptureDestinationsClient.List`

```go
ctx := context.TODO()
id := changedatacapturedestinations.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ChangeDataCaptureDestinationsClient.Update`

```go
ctx := context.TODO()
id := changedatacapturedestinations.NewChangeDataCaptureDestinationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "changeDataCaptureDestinationName")

payload := changedatacapturedestinations.ChangeDataCaptureDestinationForPatchUpdate{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
