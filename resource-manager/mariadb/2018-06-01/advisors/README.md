
## `github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/advisors` Documentation

The `advisors` SDK allows for interaction with the Azure Resource Manager Service `mariadb` (API Version `2018-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/advisors"
```


### Client Initialization

```go
client := advisors.NewAdvisorsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AdvisorsClient.Get`

```go
ctx := context.TODO()
id := advisors.NewAdvisorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "advisorValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdvisorsClient.ListByServer`

```go
ctx := context.TODO()
id := advisors.NewServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue")

// alternatively `client.ListByServer(ctx, id)` can be used to do batched pagination
items, err := client.ListByServerComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
