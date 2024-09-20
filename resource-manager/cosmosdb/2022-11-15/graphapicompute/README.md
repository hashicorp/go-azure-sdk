
## `github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2022-11-15/graphapicompute` Documentation

The `graphapicompute` SDK allows for interaction with Azure Resource Manager `cosmosdb` (API Version `2022-11-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2022-11-15/graphapicompute"
```


### Client Initialization

```go
client := graphapicompute.NewGraphAPIComputeClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GraphAPIComputeClient.ServiceCreate`

```go
ctx := context.TODO()
id := graphapicompute.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "serviceName")

payload := graphapicompute.ServiceResourceCreateUpdateParameters{
	// ...
}


if err := client.ServiceCreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `GraphAPIComputeClient.ServiceDelete`

```go
ctx := context.TODO()
id := graphapicompute.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "serviceName")

if err := client.ServiceDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `GraphAPIComputeClient.ServiceGet`

```go
ctx := context.TODO()
id := graphapicompute.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "serviceName")

read, err := client.ServiceGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
