
## `github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/raiexternalsafetyprovider` Documentation

The `raiexternalsafetyprovider` SDK allows for interaction with Azure Resource Manager `cognitive` (API Version `2025-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/raiexternalsafetyprovider"
```


### Client Initialization

```go
client := raiexternalsafetyprovider.NewRaiExternalSafetyProviderClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RaiExternalSafetyProviderClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := raiexternalsafetyprovider.NewRaiExternalSafetyProviderID("12345678-1234-9876-4563-123456789012", "raiExternalSafetyProviderName")

payload := raiexternalsafetyprovider.RaiExternalSafetyProviderSchema{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RaiExternalSafetyProviderClient.Delete`

```go
ctx := context.TODO()
id := raiexternalsafetyprovider.NewRaiExternalSafetyProviderID("12345678-1234-9876-4563-123456789012", "raiExternalSafetyProviderName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `RaiExternalSafetyProviderClient.Get`

```go
ctx := context.TODO()
id := raiexternalsafetyprovider.NewRaiExternalSafetyProviderID("12345678-1234-9876-4563-123456789012", "raiExternalSafetyProviderName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RaiExternalSafetyProviderClient.List`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
