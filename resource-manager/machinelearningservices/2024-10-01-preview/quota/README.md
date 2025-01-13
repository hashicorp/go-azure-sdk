
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/quota` Documentation

The `quota` SDK allows for interaction with Azure Resource Manager `machinelearningservices` (API Version `2024-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/quota"
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


### Example Usage: `QuotaClient.PTUQuotaGetAvailable`

```go
ctx := context.TODO()
id := quota.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

read, err := client.PTUQuotaGetAvailable(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `QuotaClient.PTUQuotaList`

```go
ctx := context.TODO()
id := quota.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

// alternatively `client.PTUQuotaList(ctx, id, quota.DefaultPTUQuotaListOperationOptions())` can be used to do batched pagination
items, err := client.PTUQuotaListComplete(ctx, id, quota.DefaultPTUQuotaListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `QuotaClient.PTUQuotaListAvailable`

```go
ctx := context.TODO()
id := quota.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

// alternatively `client.PTUQuotaListAvailable(ctx, id, quota.DefaultPTUQuotaListAvailableOperationOptions())` can be used to do batched pagination
items, err := client.PTUQuotaListAvailableComplete(ctx, id, quota.DefaultPTUQuotaListAvailableOperationOptions())
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
