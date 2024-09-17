
## `github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsblobauditing` Documentation

The `sqlpoolsblobauditing` SDK allows for interaction with Azure Resource Manager `synapse` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsblobauditing"
```


### Client Initialization

```go
client := sqlpoolsblobauditing.NewSqlPoolsBlobAuditingClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SqlPoolsBlobAuditingClient.ExtendedSqlPoolBlobAuditingPoliciesListBySqlPool`

```go
ctx := context.TODO()
id := sqlpoolsblobauditing.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue")

// alternatively `client.ExtendedSqlPoolBlobAuditingPoliciesListBySqlPool(ctx, id)` can be used to do batched pagination
items, err := client.ExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
