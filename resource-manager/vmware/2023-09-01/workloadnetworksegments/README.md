
## `github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-09-01/workloadnetworksegments` Documentation

The `workloadnetworksegments` SDK allows for interaction with Azure Resource Manager `vmware` (API Version `2023-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-09-01/workloadnetworksegments"
```


### Client Initialization

```go
client := workloadnetworksegments.NewWorkloadNetworkSegmentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkloadNetworkSegmentsClient.WorkloadNetworksCreateSegments`

```go
ctx := context.TODO()
id := workloadnetworksegments.NewSegmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "segmentIdValue")

payload := workloadnetworksegments.WorkloadNetworkSegment{
	// ...
}


if err := client.WorkloadNetworksCreateSegmentsThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworkSegmentsClient.WorkloadNetworksDeleteSegment`

```go
ctx := context.TODO()
id := workloadnetworksegments.NewSegmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "segmentIdValue")

if err := client.WorkloadNetworksDeleteSegmentThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworkSegmentsClient.WorkloadNetworksGetSegment`

```go
ctx := context.TODO()
id := workloadnetworksegments.NewSegmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "segmentIdValue")

read, err := client.WorkloadNetworksGetSegment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkloadNetworkSegmentsClient.WorkloadNetworksListSegments`

```go
ctx := context.TODO()
id := workloadnetworksegments.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue")

// alternatively `client.WorkloadNetworksListSegments(ctx, id)` can be used to do batched pagination
items, err := client.WorkloadNetworksListSegmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkloadNetworkSegmentsClient.WorkloadNetworksUpdateSegments`

```go
ctx := context.TODO()
id := workloadnetworksegments.NewSegmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "segmentIdValue")

payload := workloadnetworksegments.WorkloadNetworkSegment{
	// ...
}


if err := client.WorkloadNetworksUpdateSegmentsThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
