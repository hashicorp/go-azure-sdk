
## `github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/logfiles` Documentation

The `logfiles` SDK allows for interaction with Azure Resource Manager `horizondb` (API Version `2026-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/logfiles"
```


### Client Initialization

```go
client := logfiles.NewLogFilesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LogFilesClient.Get`

```go
ctx := context.TODO()
id := logfiles.NewLogFileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "logFileName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LogFilesClient.List`

```go
ctx := context.TODO()
id := logfiles.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `LogFilesClient.ListDownloadURLs`

```go
ctx := context.TODO()
id := logfiles.NewLogFileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "logFileName")

read, err := client.ListDownloadURLs(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
