
## `github.com/hashicorp/go-azure-sdk/resource-manager/insights/2023-10-01/metrics` Documentation

The `metrics` SDK allows for interaction with the Azure Resource Manager Service `insights` (API Version `2023-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/insights/2023-10-01/metrics"
```


### Client Initialization

```go
client := metrics.NewMetricsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MetricsClient.List`

```go
ctx := context.TODO()
id := metrics.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

read, err := client.List(ctx, id, metrics.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MetricsClient.ListAtSubscriptionScope`

```go
ctx := context.TODO()
id := metrics.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.ListAtSubscriptionScope(ctx, id, metrics.DefaultListAtSubscriptionScopeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MetricsClient.ListAtSubscriptionScopePost`

```go
ctx := context.TODO()
id := metrics.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

payload := metrics.SubscriptionScopeMetricsRequestBodyParameters{
	// ...
}


read, err := client.ListAtSubscriptionScopePost(ctx, id, payload, metrics.DefaultListAtSubscriptionScopePostOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
