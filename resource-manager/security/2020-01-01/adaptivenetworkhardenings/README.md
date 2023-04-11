
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/adaptivenetworkhardenings` Documentation

The `adaptivenetworkhardenings` SDK allows for interaction with the Azure Resource Manager Service `security` (API Version `2020-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/adaptivenetworkhardenings"
```


### Client Initialization

```go
client := adaptivenetworkhardenings.NewAdaptiveNetworkHardeningsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AdaptiveNetworkHardeningsClient.Enforce`

```go
ctx := context.TODO()
id := adaptivenetworkhardenings.NewAdaptiveNetworkHardeningID("12345678-1234-9876-4563-123456789012", "example-resource-group", "providerValue", "resourceTypeValue", "resourceValue", "adaptiveNetworkHardeningValue")

payload := adaptivenetworkhardenings.AdaptiveNetworkHardeningEnforceRequest{
	// ...
}


if err := client.EnforceThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AdaptiveNetworkHardeningsClient.Get`

```go
ctx := context.TODO()
id := adaptivenetworkhardenings.NewAdaptiveNetworkHardeningID("12345678-1234-9876-4563-123456789012", "example-resource-group", "providerValue", "resourceTypeValue", "resourceValue", "adaptiveNetworkHardeningValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdaptiveNetworkHardeningsClient.ListByExtendedResource`

```go
ctx := context.TODO()
id := adaptivenetworkhardenings.NewProviderID("12345678-1234-9876-4563-123456789012", "example-resource-group", "providerValue", "resourceTypeValue", "resourceValue")

// alternatively `client.ListByExtendedResource(ctx, id)` can be used to do batched pagination
items, err := client.ListByExtendedResourceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
