
## `github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2018-06-01/waitstatistics` Documentation

The `waitstatistics` SDK allows for interaction with the Azure Resource Manager Service `postgresql` (API Version `2018-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2018-06-01/waitstatistics"
```


### Client Initialization

```go
client := waitstatistics.NewWaitStatisticsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WaitStatisticsClient.Get`

```go
ctx := context.TODO()
id := waitstatistics.NewWaitStatisticID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "waitStatisticsIdValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WaitStatisticsClient.ListByServer`

```go
ctx := context.TODO()
id := waitstatistics.NewServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue")

payload := waitstatistics.WaitStatisticsInput{
	// ...
}


// alternatively `client.ListByServer(ctx, id, payload)` can be used to do batched pagination
items, err := client.ListByServerComplete(ctx, id, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
