
## `github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-12-01/feature` Documentation

The `feature` SDK allows for interaction with Azure Resource Manager `purview` (API Version `2021-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-12-01/feature"
```


### Client Initialization

```go
client := feature.NewFeatureClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `FeatureClient.AccountGet`

```go
ctx := context.TODO()
id := feature.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName")

payload := feature.BatchFeatureRequest{
	// ...
}


read, err := client.AccountGet(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FeatureClient.SubscriptionGet`

```go
ctx := context.TODO()
id := feature.NewLocationID("12345678-1234-9876-4563-123456789012", "location")

payload := feature.BatchFeatureRequest{
	// ...
}


read, err := client.SubscriptionGet(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
