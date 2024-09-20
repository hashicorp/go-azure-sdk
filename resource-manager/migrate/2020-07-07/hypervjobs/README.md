
## `github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-07-07/hypervjobs` Documentation

The `hypervjobs` SDK allows for interaction with Azure Resource Manager `migrate` (API Version `2020-07-07`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-07-07/hypervjobs"
```


### Client Initialization

```go
client := hypervjobs.NewHyperVJobsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `HyperVJobsClient.GetAllJobsInSite`

```go
ctx := context.TODO()
id := hypervjobs.NewHyperVSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

// alternatively `client.GetAllJobsInSite(ctx, id)` can be used to do batched pagination
items, err := client.GetAllJobsInSiteComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `HyperVJobsClient.GetJob`

```go
ctx := context.TODO()
id := commonids.NewHyperVSiteJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hyperVSiteName", "jobName")

read, err := client.GetJob(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
