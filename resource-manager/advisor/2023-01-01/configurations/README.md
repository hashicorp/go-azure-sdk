
## `github.com/hashicorp/go-azure-sdk/resource-manager/advisor/2023-01-01/configurations` Documentation

The `configurations` SDK allows for interaction with the Azure Resource Manager Service `advisor` (API Version `2023-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/advisor/2023-01-01/configurations"
```


### Client Initialization

```go
client := configurations.NewConfigurationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ConfigurationsClient.CreateInResourceGroup`

```go
ctx := context.TODO()
id := configurations.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

payload := configurations.ConfigData{
	// ...
}


read, err := client.CreateInResourceGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConfigurationsClient.CreateInSubscription`

```go
ctx := context.TODO()
id := configurations.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

payload := configurations.ConfigData{
	// ...
}


read, err := client.CreateInSubscription(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConfigurationsClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := configurations.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

read, err := client.ListByResourceGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConfigurationsClient.ListBySubscription`

```go
ctx := context.TODO()
id := configurations.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
