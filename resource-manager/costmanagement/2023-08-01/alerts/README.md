
## `github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2023-08-01/alerts` Documentation

The `alerts` SDK allows for interaction with the Azure Resource Manager Service `costmanagement` (API Version `2023-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2023-08-01/alerts"
```


### Client Initialization

```go
client := alerts.NewAlertsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AlertsClient.Dismiss`

```go
ctx := context.TODO()
id := alerts.NewScopedAlertID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "alertIdValue")

payload := alerts.DismissAlertPayload{
	// ...
}


read, err := client.Dismiss(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AlertsClient.Get`

```go
ctx := context.TODO()
id := alerts.NewScopedAlertID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "alertIdValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AlertsClient.List`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

read, err := client.List(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AlertsClient.ListExternal`

```go
ctx := context.TODO()
id := alerts.NewExternalCloudProviderTypeID("example", "externalCloudProviderIdValue")

read, err := client.ListExternal(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
