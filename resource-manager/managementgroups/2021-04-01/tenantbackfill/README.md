
## `github.com/hashicorp/go-azure-sdk/resource-manager/managementgroups/2021-04-01/tenantbackfill` Documentation

The `tenantbackfill` SDK allows for interaction with Azure Resource Manager `managementgroups` (API Version `2021-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/managementgroups/2021-04-01/tenantbackfill"
```


### Client Initialization

```go
client := tenantbackfill.NewTenantBackfillClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TenantBackfillClient.StartTenantBackfill`

```go
ctx := context.TODO()


read, err := client.StartTenantBackfill(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TenantBackfillClient.Status`

```go
ctx := context.TODO()


read, err := client.Status(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
