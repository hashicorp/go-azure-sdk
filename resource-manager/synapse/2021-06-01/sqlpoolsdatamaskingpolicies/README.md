
## `github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsdatamaskingpolicies` Documentation

The `sqlpoolsdatamaskingpolicies` SDK allows for interaction with Azure Resource Manager `synapse` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsdatamaskingpolicies"
```


### Client Initialization

```go
client := sqlpoolsdatamaskingpolicies.NewSqlPoolsDataMaskingPoliciesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SqlPoolsDataMaskingPoliciesClient.DataMaskingPoliciesCreateOrUpdate`

```go
ctx := context.TODO()
id := sqlpoolsdatamaskingpolicies.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue")

payload := sqlpoolsdatamaskingpolicies.DataMaskingPolicy{
	// ...
}


read, err := client.DataMaskingPoliciesCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SqlPoolsDataMaskingPoliciesClient.DataMaskingPoliciesGet`

```go
ctx := context.TODO()
id := sqlpoolsdatamaskingpolicies.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue")

read, err := client.DataMaskingPoliciesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
