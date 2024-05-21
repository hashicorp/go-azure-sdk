
## `github.com/hashicorp/go-azure-sdk/resource-manager/resources/2015-11-01/tenants` Documentation

The `tenants` SDK allows for interaction with the Azure Resource Manager Service `resources` (API Version `2015-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/resources/2015-11-01/tenants"
```


### Client Initialization

```go
client := tenants.NewTenantsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TenantsClient.CheckResourceName`

```go
ctx := context.TODO()

payload := tenants.ResourceName{
	// ...
}


read, err := client.CheckResourceName(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TenantsClient.List`

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
