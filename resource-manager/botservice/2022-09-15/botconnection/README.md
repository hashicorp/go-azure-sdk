
## `github.com/hashicorp/go-azure-sdk/resource-manager/botservice/2022-09-15/botconnection` Documentation

The `botconnection` SDK allows for interaction with Azure Resource Manager `botservice` (API Version `2022-09-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/botservice/2022-09-15/botconnection"
```


### Client Initialization

```go
client := botconnection.NewBotConnectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BotConnectionClient.Create`

```go
ctx := context.TODO()
id := botconnection.NewConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "botServiceName", "connectionName")

payload := botconnection.ConnectionSetting{
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


### Example Usage: `BotConnectionClient.Delete`

```go
ctx := context.TODO()
id := botconnection.NewConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "botServiceName", "connectionName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BotConnectionClient.Get`

```go
ctx := context.TODO()
id := botconnection.NewConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "botServiceName", "connectionName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BotConnectionClient.ListByBotService`

```go
ctx := context.TODO()
id := commonids.NewBotServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "botServiceName")

// alternatively `client.ListByBotService(ctx, id)` can be used to do batched pagination
items, err := client.ListByBotServiceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BotConnectionClient.ListWithSecrets`

```go
ctx := context.TODO()
id := botconnection.NewConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "botServiceName", "connectionName")

read, err := client.ListWithSecrets(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BotConnectionClient.Update`

```go
ctx := context.TODO()
id := botconnection.NewConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "botServiceName", "connectionName")

payload := botconnection.ConnectionSetting{
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
