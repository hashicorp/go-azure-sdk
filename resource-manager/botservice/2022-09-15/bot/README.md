
## `github.com/hashicorp/go-azure-sdk/resource-manager/botservice/2022-09-15/bot` Documentation

The `bot` SDK allows for interaction with the Azure Resource Manager Service `botservice` (API Version `2022-09-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/botservice/2022-09-15/bot"
```


### Client Initialization

```go
client := bot.NewBotClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BotClient.Create`

```go
ctx := context.TODO()
id := bot.NewBotServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "botServiceValue")

payload := bot.Bot{
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


### Example Usage: `BotClient.Delete`

```go
ctx := context.TODO()
id := bot.NewBotServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "botServiceValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BotClient.Get`

```go
ctx := context.TODO()
id := bot.NewBotServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "botServiceValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BotClient.GetCheckNameAvailability`

```go
ctx := context.TODO()

payload := bot.CheckNameAvailabilityRequestBody{
	// ...
}


read, err := client.GetCheckNameAvailability(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BotClient.List`

```go
ctx := context.TODO()
id := bot.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BotClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := bot.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BotClient.Update`

```go
ctx := context.TODO()
id := bot.NewBotServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "botServiceValue")

payload := bot.Bot{
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
