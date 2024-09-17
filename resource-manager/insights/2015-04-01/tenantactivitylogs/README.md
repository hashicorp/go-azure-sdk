
## `github.com/hashicorp/go-azure-sdk/resource-manager/insights/2015-04-01/tenantactivitylogs` Documentation

The `tenantactivitylogs` SDK allows for interaction with Azure Resource Manager `insights` (API Version `2015-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/insights/2015-04-01/tenantactivitylogs"
```


### Client Initialization

```go
client := tenantactivitylogs.NewTenantActivityLogsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TenantActivityLogsClient.List`

```go
ctx := context.TODO()


// alternatively `client.List(ctx, tenantactivitylogs.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, tenantactivitylogs.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
