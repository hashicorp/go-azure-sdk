
## `github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-05-01-preview/insights` Documentation

The `insights` SDK allows for interaction with the Azure Resource Manager Service `insights` (API Version `2021-05-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-05-01-preview/insights"
```


### Client Initialization

```go
client := insights.NewInsightsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InsightsClient.AutoscaleSettingsUpdate`

```go
ctx := context.TODO()
id := insights.NewAutoScaleSettingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "autoscaleSettingValue")

payload := insights.AutoscaleSettingResourcePatch{
	// ...
}


read, err := client.AutoscaleSettingsUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
