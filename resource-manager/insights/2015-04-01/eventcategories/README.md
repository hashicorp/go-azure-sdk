
## `github.com/hashicorp/go-azure-sdk/resource-manager/insights/2015-04-01/eventcategories` Documentation

The `eventcategories` SDK allows for interaction with the Azure Resource Manager Service `insights` (API Version `2015-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/insights/2015-04-01/eventcategories"
```


### Client Initialization

```go
client := eventcategories.NewEventCategoriesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EventCategoriesClient.List`

```go
ctx := context.TODO()


read, err := client.List(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
