
## `github.com/hashicorp/go-azure-sdk/resource-manager/trafficmanager/2018-08-01/realusermetrics` Documentation

The `realusermetrics` SDK allows for interaction with the Azure Resource Manager Service `trafficmanager` (API Version `2018-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/trafficmanager/2018-08-01/realusermetrics"
```


### Client Initialization

```go
client := realusermetrics.NewRealUserMetricsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `RealUserMetricsClient.TrafficManagerUserMetricsKeysCreateOrUpdate`

```go
ctx := context.TODO()
id := realusermetrics.NewSubscriptionID()
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
id := realusermetrics.NewSubscriptionID()
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
id := realusermetrics.NewSubscriptionID()
read, err := client.TrafficManagerUserMetricsKeysGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
