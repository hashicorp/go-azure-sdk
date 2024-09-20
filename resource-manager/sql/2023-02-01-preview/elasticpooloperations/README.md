
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-02-01-preview/elasticpooloperations` Documentation

The `elasticpooloperations` SDK allows for interaction with Azure Resource Manager `sql` (API Version `2023-02-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-02-01-preview/elasticpooloperations"
```


### Client Initialization

```go
client := elasticpooloperations.NewElasticPoolOperationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ElasticPoolOperationsClient.Cancel`

```go
ctx := context.TODO()
id := elasticpooloperations.NewOperationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "elasticPoolName", "operationId")

read, err := client.Cancel(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ElasticPoolOperationsClient.ListByElasticPool`

```go
ctx := context.TODO()
id := commonids.NewSqlElasticPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "elasticPoolName")

// alternatively `client.ListByElasticPool(ctx, id)` can be used to do batched pagination
items, err := client.ListByElasticPoolComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
