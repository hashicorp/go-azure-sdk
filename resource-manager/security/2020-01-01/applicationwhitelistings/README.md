
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/applicationwhitelistings` Documentation

The `applicationwhitelistings` SDK allows for interaction with the Azure Resource Manager Service `security` (API Version `2020-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/applicationwhitelistings"
```


### Client Initialization

```go
client := applicationwhitelistings.NewApplicationWhitelistingsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ApplicationWhitelistingsClient.AdaptiveApplicationControlsDelete`

```go
ctx := context.TODO()
id := applicationwhitelistings.NewApplicationWhitelistingID("12345678-1234-9876-4563-123456789012", "locationValue", "applicationWhitelistingValue")

read, err := client.AdaptiveApplicationControlsDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationWhitelistingsClient.AdaptiveApplicationControlsGet`

```go
ctx := context.TODO()
id := applicationwhitelistings.NewApplicationWhitelistingID("12345678-1234-9876-4563-123456789012", "locationValue", "applicationWhitelistingValue")

read, err := client.AdaptiveApplicationControlsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationWhitelistingsClient.AdaptiveApplicationControlsList`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.AdaptiveApplicationControlsList(ctx, id, applicationwhitelistings.DefaultAdaptiveApplicationControlsListOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationWhitelistingsClient.AdaptiveApplicationControlsPut`

```go
ctx := context.TODO()
id := applicationwhitelistings.NewApplicationWhitelistingID("12345678-1234-9876-4563-123456789012", "locationValue", "applicationWhitelistingValue")

payload := applicationwhitelistings.AdaptiveApplicationControlGroup{
	// ...
}


read, err := client.AdaptiveApplicationControlsPut(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
