
## `github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2024-03-01/dppjob` Documentation

The `dppjob` SDK allows for interaction with the Azure Resource Manager Service `dataprotection` (API Version `2024-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2024-03-01/dppjob"
```


### Client Initialization

```go
client := dppjob.NewDppJobClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DppJobClient.FetchCrossRegionRestoreJobGet`

```go
ctx := context.TODO()
id := dppjob.NewProviderLocationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationValue")

payload := dppjob.CrossRegionRestoreJobRequest{
	// ...
}


read, err := client.FetchCrossRegionRestoreJobGet(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DppJobClient.FetchCrossRegionRestoreJobsList`

```go
ctx := context.TODO()
id := dppjob.NewProviderLocationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationValue")

payload := dppjob.CrossRegionRestoreJobsRequest{
	// ...
}


// alternatively `client.FetchCrossRegionRestoreJobsList(ctx, id, payload, dppjob.DefaultFetchCrossRegionRestoreJobsListOperationOptions())` can be used to do batched pagination
items, err := client.FetchCrossRegionRestoreJobsListComplete(ctx, id, payload, dppjob.DefaultFetchCrossRegionRestoreJobsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
