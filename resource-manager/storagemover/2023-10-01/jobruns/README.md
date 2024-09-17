
## `github.com/hashicorp/go-azure-sdk/resource-manager/storagemover/2023-10-01/jobruns` Documentation

The `jobruns` SDK allows for interaction with Azure Resource Manager `storagemover` (API Version `2023-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/storagemover/2023-10-01/jobruns"
```


### Client Initialization

```go
client := jobruns.NewJobRunsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JobRunsClient.Get`

```go
ctx := context.TODO()
id := jobruns.NewJobRunID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageMoverValue", "projectValue", "jobDefinitionValue", "jobRunValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JobRunsClient.List`

```go
ctx := context.TODO()
id := jobruns.NewJobDefinitionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageMoverValue", "projectValue", "jobDefinitionValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
