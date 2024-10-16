
## `github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2024-11-01/billingcontainers` Documentation

The `billingcontainers` SDK allows for interaction with Azure Resource Manager `deviceregistry` (API Version `2024-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2024-11-01/billingcontainers"
```


### Client Initialization

```go
client := billingcontainers.NewBillingContainersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BillingContainersClient.Get`

```go
ctx := context.TODO()
id := billingcontainers.NewBillingContainerID("12345678-1234-9876-4563-123456789012", "billingContainerName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingContainersClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
