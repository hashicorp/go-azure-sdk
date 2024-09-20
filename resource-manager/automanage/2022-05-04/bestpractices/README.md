
## `github.com/hashicorp/go-azure-sdk/resource-manager/automanage/2022-05-04/bestpractices` Documentation

The `bestpractices` SDK allows for interaction with Azure Resource Manager `automanage` (API Version `2022-05-04`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/automanage/2022-05-04/bestpractices"
```


### Client Initialization

```go
client := bestpractices.NewBestPracticesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BestPracticesClient.Get`

```go
ctx := context.TODO()
id := bestpractices.NewBestPracticeID("bestPracticeName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BestPracticesClient.ListByTenant`

```go
ctx := context.TODO()


read, err := client.ListByTenant(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
