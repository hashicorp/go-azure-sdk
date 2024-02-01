
## `github.com/hashicorp/go-azure-sdk/resource-manager/botservice/2022-09-15/bothostsettings` Documentation

The `bothostsettings` SDK allows for interaction with the Azure Resource Manager Service `botservice` (API Version `2022-09-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/botservice/2022-09-15/bothostsettings"
```


### Client Initialization

```go
client := bothostsettings.NewBotHostSettingsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BotHostSettingsClient.HostSettingsGet`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.HostSettingsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
