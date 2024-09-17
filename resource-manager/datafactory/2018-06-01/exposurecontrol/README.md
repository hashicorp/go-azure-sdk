
## `github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/exposurecontrol` Documentation

The `exposurecontrol` SDK allows for interaction with Azure Resource Manager `datafactory` (API Version `2018-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/exposurecontrol"
```


### Client Initialization

```go
client := exposurecontrol.NewExposureControlClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ExposureControlClient.GetFeatureValue`

```go
ctx := context.TODO()
id := exposurecontrol.NewLocationID("12345678-1234-9876-4563-123456789012", "locationIdValue")

payload := exposurecontrol.ExposureControlRequest{
	// ...
}


read, err := client.GetFeatureValue(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExposureControlClient.GetFeatureValueByFactory`

```go
ctx := context.TODO()
id := exposurecontrol.NewFactoryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue")

payload := exposurecontrol.ExposureControlRequest{
	// ...
}


read, err := client.GetFeatureValueByFactory(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExposureControlClient.QueryFeatureValuesByFactory`

```go
ctx := context.TODO()
id := exposurecontrol.NewFactoryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue")

payload := exposurecontrol.ExposureControlBatchRequest{
	// ...
}


read, err := client.QueryFeatureValuesByFactory(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
