
## `github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/locationbasedrecommendedactionsessionsresult` Documentation

The `locationbasedrecommendedactionsessionsresult` SDK allows for interaction with the Azure Resource Manager Service `mariadb` (API Version `2018-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/locationbasedrecommendedactionsessionsresult"
```


### Client Initialization

```go
client := locationbasedrecommendedactionsessionsresult.NewLocationBasedRecommendedActionSessionsResultClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LocationBasedRecommendedActionSessionsResultClient.List`

```go
ctx := context.TODO()
id := locationbasedrecommendedactionsessionsresult.NewRecommendedActionSessionsOperationResultID("12345678-1234-9876-4563-123456789012", "locationValue", "operationIdValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
