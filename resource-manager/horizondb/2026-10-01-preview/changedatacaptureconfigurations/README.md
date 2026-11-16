
## `github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/changedatacaptureconfigurations` Documentation

The `changedatacaptureconfigurations` SDK allows for interaction with Azure Resource Manager `horizondb` (API Version `2026-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/changedatacaptureconfigurations"
```


### Client Initialization

```go
client := changedatacaptureconfigurations.NewChangeDataCaptureConfigurationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ChangeDataCaptureConfigurationsClient.CreateOrReplace`

```go
ctx := context.TODO()
id := changedatacaptureconfigurations.NewChangeDataCaptureConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "changeDataCaptureConfigurationName")

payload := changedatacaptureconfigurations.ChangeDataCaptureConfiguration{
	// ...
}


if err := client.CreateOrReplaceThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ChangeDataCaptureConfigurationsClient.Delete`

```go
ctx := context.TODO()
id := changedatacaptureconfigurations.NewChangeDataCaptureConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "changeDataCaptureConfigurationName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ChangeDataCaptureConfigurationsClient.Get`

```go
ctx := context.TODO()
id := changedatacaptureconfigurations.NewChangeDataCaptureConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "changeDataCaptureConfigurationName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChangeDataCaptureConfigurationsClient.List`

```go
ctx := context.TODO()
id := changedatacaptureconfigurations.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ChangeDataCaptureConfigurationsClient.Update`

```go
ctx := context.TODO()
id := changedatacaptureconfigurations.NewChangeDataCaptureConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "changeDataCaptureConfigurationName")

payload := changedatacaptureconfigurations.ChangeDataCaptureConfigurationForPatchUpdate{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
