
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01-preview/databaserecommendedactions` Documentation

The `databaserecommendedactions` SDK allows for interaction with Azure Resource Manager `sql` (API Version `2023-08-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01-preview/databaserecommendedactions"
```


### Client Initialization

```go
client := databaserecommendedactions.NewDatabaseRecommendedActionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DatabaseRecommendedActionsClient.Get`

```go
ctx := context.TODO()
id := databaserecommendedactions.NewRecommendedActionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseName", "advisorName", "recommendedActionName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DatabaseRecommendedActionsClient.ListByDatabaseAdvisor`

```go
ctx := context.TODO()
id := databaserecommendedactions.NewDatabaseAdvisorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseName", "advisorName")

read, err := client.ListByDatabaseAdvisor(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DatabaseRecommendedActionsClient.Update`

```go
ctx := context.TODO()
id := databaserecommendedactions.NewRecommendedActionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseName", "advisorName", "recommendedActionName")

payload := databaserecommendedactions.RecommendedAction{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
