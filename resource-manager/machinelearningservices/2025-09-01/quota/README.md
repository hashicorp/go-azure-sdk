
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2025-09-01/quota` Documentation

The `quota` SDK allows for interaction with Azure Resource Manager `machinelearningservices` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2025-09-01/quota"
```


### Client Initialization

```go
client := quota.NewQuotaClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `QuotaClient.List`

```go
ctx := context.TODO()
id := quota.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `QuotaClient.Update`

```go
ctx := context.TODO()
id := quota.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

payload := quota.QuotaUpdateParameters{
	// ...
}


// alternatively `client.Update(ctx, id, payload)` can be used to do batched pagination
items, err := client.UpdateComplete(ctx, id, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
