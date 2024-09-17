
## `github.com/hashicorp/go-azure-sdk/resource-manager/trafficmanager/2022-04-01/realusermetrics` Documentation

The `realusermetrics` SDK allows for interaction with Azure Resource Manager `trafficmanager` (API Version `2022-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/trafficmanager/2022-04-01/realusermetrics"
```


### Client Initialization

```go
client := realusermetrics.NewRealUserMetricsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RealUserMetricsClient.TrafficManagerUserMetricsKeysCreateOrUpdate`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.TrafficManagerUserMetricsKeysCreateOrUpdate(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RealUserMetricsClient.TrafficManagerUserMetricsKeysDelete`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.TrafficManagerUserMetricsKeysDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RealUserMetricsClient.TrafficManagerUserMetricsKeysGet`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.TrafficManagerUserMetricsKeysGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
