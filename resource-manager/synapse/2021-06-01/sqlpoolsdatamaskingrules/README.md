
## `github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsdatamaskingrules` Documentation

The `sqlpoolsdatamaskingrules` SDK allows for interaction with Azure Resource Manager `synapse` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsdatamaskingrules"
```


### Client Initialization

```go
client := sqlpoolsdatamaskingrules.NewSqlPoolsDataMaskingRulesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SqlPoolsDataMaskingRulesClient.DataMaskingRulesCreateOrUpdate`

```go
ctx := context.TODO()
id := sqlpoolsdatamaskingrules.NewRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue", "ruleValue")

payload := sqlpoolsdatamaskingrules.DataMaskingRule{
	// ...
}


read, err := client.DataMaskingRulesCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SqlPoolsDataMaskingRulesClient.DataMaskingRulesGet`

```go
ctx := context.TODO()
id := sqlpoolsdatamaskingrules.NewRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue", "ruleValue")

read, err := client.DataMaskingRulesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SqlPoolsDataMaskingRulesClient.DataMaskingRulesListBySqlPool`

```go
ctx := context.TODO()
id := sqlpoolsdatamaskingrules.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue")

read, err := client.DataMaskingRulesListBySqlPool(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
