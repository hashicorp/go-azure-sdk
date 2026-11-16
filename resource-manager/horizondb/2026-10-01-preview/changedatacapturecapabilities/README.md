
## `github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/changedatacapturecapabilities` Documentation

The `changedatacapturecapabilities` SDK allows for interaction with Azure Resource Manager `horizondb` (API Version `2026-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/changedatacapturecapabilities"
```


### Client Initialization

```go
client := changedatacapturecapabilities.NewChangeDataCaptureCapabilitiesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ChangeDataCaptureCapabilitiesClient.CreateOrReplace`

```go
ctx := context.TODO()
id := changedatacapturecapabilities.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

payload := changedatacapturecapabilities.ChangeDataCaptureCapability{
	// ...
}


if err := client.CreateOrReplaceThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ChangeDataCaptureCapabilitiesClient.Delete`

```go
ctx := context.TODO()
id := changedatacapturecapabilities.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ChangeDataCaptureCapabilitiesClient.Get`

```go
ctx := context.TODO()
id := changedatacapturecapabilities.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChangeDataCaptureCapabilitiesClient.List`

```go
ctx := context.TODO()
id := changedatacapturecapabilities.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ChangeDataCaptureCapabilitiesClient.Update`

```go
ctx := context.TODO()
id := changedatacapturecapabilities.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

payload := changedatacapturecapabilities.ChangeDataCaptureCapabilityForPatchUpdate{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
