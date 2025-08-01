
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2025-06-01/job` Documentation

The `job` SDK allows for interaction with Azure Resource Manager `machinelearningservices` (API Version `2025-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2025-06-01/job"
```


### Client Initialization

```go
client := job.NewJobClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JobClient.Cancel`

```go
ctx := context.TODO()
id := job.NewJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "jobName")

if err := client.CancelThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `JobClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := job.NewJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "jobName")

payload := job.JobBaseResource{
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


### Example Usage: `JobClient.Delete`

```go
ctx := context.TODO()
id := job.NewJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "jobName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `JobClient.Get`

```go
ctx := context.TODO()
id := job.NewJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "jobName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JobClient.List`

```go
ctx := context.TODO()
id := job.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

// alternatively `client.List(ctx, id, job.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, job.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
