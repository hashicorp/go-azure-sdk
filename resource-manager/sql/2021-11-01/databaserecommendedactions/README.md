
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-11-01/databaserecommendedactions` Documentation

The `databaserecommendedactions` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2021-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-11-01/databaserecommendedactions"
```


### Client Initialization

```go
client := databaserecommendedactions.NewDatabaseRecommendedActionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DatabaseRecommendedActionsClient.Get`

```go
ctx := context.TODO()
id := databaserecommendedactions.NewRecommendedActionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "advisorValue", "recommendedActionValue")

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
id := databaserecommendedactions.NewDatabaseAdvisorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "advisorValue")

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
id := databaserecommendedactions.NewRecommendedActionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "advisorValue", "recommendedActionValue")

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
