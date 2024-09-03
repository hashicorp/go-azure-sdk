
## `github.com/hashicorp/go-azure-sdk/resource-manager/resources/2024-08-01/databoundaries` Documentation

The `databoundaries` SDK allows for interaction with the Azure Resource Manager Service `resources` (API Version `2024-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/resources/2024-08-01/databoundaries"
```


### Client Initialization

```go
client := databoundaries.NewDataBoundariesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DataBoundariesClient.GetScope`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

read, err := client.GetScope(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataBoundariesClient.GetTenant`

```go
ctx := context.TODO()


read, err := client.GetTenant(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataBoundariesClient.Put`

```go
ctx := context.TODO()

payload := databoundaries.DataBoundaryDefinition{
	// ...
}


read, err := client.Put(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
