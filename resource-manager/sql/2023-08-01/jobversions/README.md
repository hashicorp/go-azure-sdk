
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01/jobversions` Documentation

The `jobversions` SDK allows for interaction with Azure Resource Manager `sql` (API Version `2023-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01/jobversions"
```


### Client Initialization

```go
client := jobversions.NewJobVersionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JobVersionsClient.Get`

```go
ctx := context.TODO()
id := jobversions.NewVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "jobAgentName", "jobName", "versionName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JobVersionsClient.ListByJob`

```go
ctx := context.TODO()
id := jobversions.NewJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "jobAgentName", "jobName")

// alternatively `client.ListByJob(ctx, id)` can be used to do batched pagination
items, err := client.ListByJobComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
