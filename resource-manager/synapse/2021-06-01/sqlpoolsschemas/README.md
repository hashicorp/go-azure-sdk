
## `github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsschemas` Documentation

The `sqlpoolsschemas` SDK allows for interaction with Azure Resource Manager `synapse` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsschemas"
```


### Client Initialization

```go
client := sqlpoolsschemas.NewSqlPoolsSchemasClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SqlPoolsSchemasClient.SqlPoolSchemasList`

```go
ctx := context.TODO()
id := sqlpoolsschemas.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue")

// alternatively `client.SqlPoolSchemasList(ctx, id, sqlpoolsschemas.DefaultSqlPoolSchemasListOperationOptions())` can be used to do batched pagination
items, err := client.SqlPoolSchemasListComplete(ctx, id, sqlpoolsschemas.DefaultSqlPoolSchemasListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SqlPoolsSchemasClient.SqlPoolVulnerabilityAssessmentScansList`

```go
ctx := context.TODO()
id := sqlpoolsschemas.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue")

// alternatively `client.SqlPoolVulnerabilityAssessmentScansList(ctx, id)` can be used to do batched pagination
items, err := client.SqlPoolVulnerabilityAssessmentScansListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SqlPoolsSchemasClient.SqlPoolVulnerabilityAssessmentsList`

```go
ctx := context.TODO()
id := sqlpoolsschemas.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue")

// alternatively `client.SqlPoolVulnerabilityAssessmentsList(ctx, id)` can be used to do batched pagination
items, err := client.SqlPoolVulnerabilityAssessmentsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
