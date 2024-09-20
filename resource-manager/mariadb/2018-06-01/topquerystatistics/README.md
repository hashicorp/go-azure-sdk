
## `github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/topquerystatistics` Documentation

The `topquerystatistics` SDK allows for interaction with Azure Resource Manager `mariadb` (API Version `2018-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/topquerystatistics"
```


### Client Initialization

```go
client := topquerystatistics.NewTopQueryStatisticsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TopQueryStatisticsClient.Get`

```go
ctx := context.TODO()
id := topquerystatistics.NewTopQueryStatisticID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "queryStatisticId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TopQueryStatisticsClient.ListByServer`

```go
ctx := context.TODO()
id := topquerystatistics.NewServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName")

payload := topquerystatistics.TopQueryStatisticsInput{
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
