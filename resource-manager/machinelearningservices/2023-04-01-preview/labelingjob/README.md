
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/labelingjob` Documentation

The `labelingjob` SDK allows for interaction with the Azure Resource Manager Service `machinelearningservices` (API Version `2023-04-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/labelingjob"
```


### Client Initialization

```go
client := labelingjob.NewLabelingJobClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LabelingJobClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := labelingjob.NewLabelingJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "labelingJobValue")

payload := labelingjob.LabelingJobResource{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `LabelingJobClient.Delete`

```go
ctx := context.TODO()
id := labelingjob.NewLabelingJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "labelingJobValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LabelingJobClient.ExportLabels`

```go
ctx := context.TODO()
id := labelingjob.NewLabelingJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "labelingJobValue")

payload := labelingjob.ExportSummary{
	// ...
}


if err := client.ExportLabelsThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `LabelingJobClient.Get`

```go
ctx := context.TODO()
id := labelingjob.NewLabelingJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "labelingJobValue")

read, err := client.Get(ctx, id, labelingjob.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LabelingJobClient.List`

```go
ctx := context.TODO()
id := labelingjob.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

// alternatively `client.List(ctx, id, labelingjob.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, labelingjob.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `LabelingJobClient.Pause`

```go
ctx := context.TODO()
id := labelingjob.NewLabelingJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "labelingJobValue")

read, err := client.Pause(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LabelingJobClient.Resume`

```go
ctx := context.TODO()
id := labelingjob.NewLabelingJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "labelingJobValue")

if err := client.ResumeThenPoll(ctx, id); err != nil {
	// handle the error
}
```
