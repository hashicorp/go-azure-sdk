
## `github.com/hashicorp/go-azure-sdk/resource-manager/portal/2020-09-01-preview/listtenantconfigurationviolations` Documentation

The `listtenantconfigurationviolations` SDK allows for interaction with the Azure Resource Manager Service `portal` (API Version `2020-09-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/portal/2020-09-01-preview/listtenantconfigurationviolations"
```


### Client Initialization

```go
client := listtenantconfigurationviolations.NewListTenantConfigurationViolationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ListTenantConfigurationViolationsClient.List`

```go
ctx := context.TODO()


// alternatively `client.List(ctx)` can be used to do batched pagination
items, err := client.ListComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
