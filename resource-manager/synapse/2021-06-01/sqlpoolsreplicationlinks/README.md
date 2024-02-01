
## `github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsreplicationlinks` Documentation

The `sqlpoolsreplicationlinks` SDK allows for interaction with the Azure Resource Manager Service `synapse` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsreplicationlinks"
```


### Client Initialization

```go
client := sqlpoolsreplicationlinks.NewSqlPoolsReplicationLinksClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SqlPoolsReplicationLinksClient.SqlPoolReplicationLinksGetByName`

```go
ctx := context.TODO()
id := sqlpoolsreplicationlinks.NewReplicationLinkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue", "linkIdValue")

read, err := client.SqlPoolReplicationLinksGetByName(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SqlPoolsReplicationLinksClient.SqlPoolReplicationLinksList`

```go
ctx := context.TODO()
id := sqlpoolsreplicationlinks.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue")

// alternatively `client.SqlPoolReplicationLinksList(ctx, id)` can be used to do batched pagination
items, err := client.SqlPoolReplicationLinksListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
