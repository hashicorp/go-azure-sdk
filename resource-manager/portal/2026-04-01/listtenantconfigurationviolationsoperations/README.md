
## `github.com/hashicorp/go-azure-sdk/resource-manager/portal/2026-04-01/listtenantconfigurationviolationsoperations` Documentation

The `listtenantconfigurationviolationsoperations` SDK allows for interaction with Azure Resource Manager `portal` (API Version `2026-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/portal/2026-04-01/listtenantconfigurationviolationsoperations"
```


### Client Initialization

```go
client := listtenantconfigurationviolationsoperations.NewListTenantConfigurationViolationsOperationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ListTenantConfigurationViolationsOperationsClient.ListTenantConfigurationViolationsList`

```go
ctx := context.TODO()


// alternatively `client.ListTenantConfigurationViolationsList(ctx)` can be used to do batched pagination
items, err := client.ListTenantConfigurationViolationsListComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
