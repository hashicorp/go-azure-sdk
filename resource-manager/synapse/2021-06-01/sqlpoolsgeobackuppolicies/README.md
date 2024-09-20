
## `github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsgeobackuppolicies` Documentation

The `sqlpoolsgeobackuppolicies` SDK allows for interaction with Azure Resource Manager `synapse` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsgeobackuppolicies"
```


### Client Initialization

```go
client := sqlpoolsgeobackuppolicies.NewSqlPoolsGeoBackupPoliciesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SqlPoolsGeoBackupPoliciesClient.SqlPoolGeoBackupPoliciesCreateOrUpdate`

```go
ctx := context.TODO()
id := sqlpoolsgeobackuppolicies.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "sqlPoolName")

payload := sqlpoolsgeobackuppolicies.GeoBackupPolicy{
	// ...
}


read, err := client.SqlPoolGeoBackupPoliciesCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SqlPoolsGeoBackupPoliciesClient.SqlPoolGeoBackupPoliciesGet`

```go
ctx := context.TODO()
id := sqlpoolsgeobackuppolicies.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "sqlPoolName")

read, err := client.SqlPoolGeoBackupPoliciesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SqlPoolsGeoBackupPoliciesClient.SqlPoolGeoBackupPoliciesList`

```go
ctx := context.TODO()
id := sqlpoolsgeobackuppolicies.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "sqlPoolName")

read, err := client.SqlPoolGeoBackupPoliciesList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
