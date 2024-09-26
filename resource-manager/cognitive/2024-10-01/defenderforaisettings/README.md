
## `github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2024-10-01/defenderforaisettings` Documentation

The `defenderforaisettings` SDK allows for interaction with Azure Resource Manager `cognitive` (API Version `2024-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2024-10-01/defenderforaisettings"
```


### Client Initialization

```go
client := defenderforaisettings.NewDefenderForAISettingsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DefenderForAISettingsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := defenderforaisettings.NewDefenderForAISettingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "defenderForAISettingName")

payload := defenderforaisettings.DefenderForAISetting{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DefenderForAISettingsClient.Get`

```go
ctx := context.TODO()
id := defenderforaisettings.NewDefenderForAISettingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "defenderForAISettingName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DefenderForAISettingsClient.List`

```go
ctx := context.TODO()
id := defenderforaisettings.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DefenderForAISettingsClient.Update`

```go
ctx := context.TODO()
id := defenderforaisettings.NewDefenderForAISettingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "defenderForAISettingName")

payload := defenderforaisettings.DefenderForAISetting{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
