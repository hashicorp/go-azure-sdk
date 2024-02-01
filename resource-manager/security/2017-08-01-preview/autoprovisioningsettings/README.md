
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2017-08-01-preview/autoprovisioningsettings` Documentation

The `autoprovisioningsettings` SDK allows for interaction with the Azure Resource Manager Service `security` (API Version `2017-08-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2017-08-01-preview/autoprovisioningsettings"
```


### Client Initialization

```go
client := autoprovisioningsettings.NewAutoProvisioningSettingsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AutoProvisioningSettingsClient.Create`

```go
ctx := context.TODO()
id := autoprovisioningsettings.NewAutoProvisioningSettingID("12345678-1234-9876-4563-123456789012", "autoProvisioningSettingValue")

payload := autoprovisioningsettings.AutoProvisioningSetting{
	// ...
}


read, err := client.Create(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AutoProvisioningSettingsClient.Get`

```go
ctx := context.TODO()
id := autoprovisioningsettings.NewAutoProvisioningSettingID("12345678-1234-9876-4563-123456789012", "autoProvisioningSettingValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AutoProvisioningSettingsClient.List`

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
