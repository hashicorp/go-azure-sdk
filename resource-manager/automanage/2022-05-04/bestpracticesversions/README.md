
## `github.com/hashicorp/go-azure-sdk/resource-manager/automanage/2022-05-04/bestpracticesversions` Documentation

The `bestpracticesversions` SDK allows for interaction with the Azure Resource Manager Service `automanage` (API Version `2022-05-04`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/automanage/2022-05-04/bestpracticesversions"
```


### Client Initialization

```go
client := bestpracticesversions.NewBestPracticesVersionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BestPracticesVersionsClient.Get`

```go
ctx := context.TODO()
id := bestpracticesversions.NewVersionID("bestPracticeValue", "versionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BestPracticesVersionsClient.ListByTenant`

```go
ctx := context.TODO()
id := bestpracticesversions.NewBestPracticeID("bestPracticeValue")

read, err := client.ListByTenant(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
