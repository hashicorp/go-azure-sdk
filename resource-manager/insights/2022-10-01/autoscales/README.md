
## `github.com/hashicorp/go-azure-sdk/resource-manager/insights/2022-10-01/autoscales` Documentation

The `autoscales` SDK allows for interaction with Azure Resource Manager `insights` (API Version `2022-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/insights/2022-10-01/autoscales"
```


### Client Initialization

```go
client := autoscales.NewAutoScalesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AutoScalesClient.ettingsUpdate`

```go
ctx := context.TODO()
id := autoscales.NewAutoScaleSettingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "autoScaleSettingName")

payload := autoscales.AutoscaleSettingResourcePatch{
	// ...
}


read, err := client.ettingsUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
