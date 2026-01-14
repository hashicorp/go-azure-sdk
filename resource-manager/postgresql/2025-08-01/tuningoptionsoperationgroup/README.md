
## `github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/tuningoptionsoperationgroup` Documentation

The `tuningoptionsoperationgroup` SDK allows for interaction with Azure Resource Manager `postgresql` (API Version `2025-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/tuningoptionsoperationgroup"
```


### Client Initialization

```go
client := tuningoptionsoperationgroup.NewTuningOptionsOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TuningOptionsOperationGroupClient.TuningOptionsGet`

```go
ctx := context.TODO()
id := tuningoptionsoperationgroup.NewTuningOptionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerName", "index")

read, err := client.TuningOptionsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TuningOptionsOperationGroupClient.TuningOptionsListByServer`

```go
ctx := context.TODO()
id := tuningoptionsoperationgroup.NewFlexibleServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerName")

// alternatively `client.TuningOptionsListByServer(ctx, id)` can be used to do batched pagination
items, err := client.TuningOptionsListByServerComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TuningOptionsOperationGroupClient.TuningOptionsListRecommendations`

```go
ctx := context.TODO()
id := tuningoptionsoperationgroup.NewTuningOptionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerName", "index")

// alternatively `client.TuningOptionsListRecommendations(ctx, id, tuningoptionsoperationgroup.DefaultTuningOptionsListRecommendationsOperationOptions())` can be used to do batched pagination
items, err := client.TuningOptionsListRecommendationsComplete(ctx, id, tuningoptionsoperationgroup.DefaultTuningOptionsListRecommendationsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
