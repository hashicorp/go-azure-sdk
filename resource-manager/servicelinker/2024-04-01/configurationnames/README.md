
## `github.com/hashicorp/go-azure-sdk/resource-manager/servicelinker/2024-04-01/configurationnames` Documentation

The `configurationnames` SDK allows for interaction with Azure Resource Manager `servicelinker` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/servicelinker/2024-04-01/configurationnames"
```


### Client Initialization

```go
client := configurationnames.NewConfigurationNamesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ConfigurationNamesClient.List`

```go
ctx := context.TODO()


// alternatively `client.List(ctx, configurationnames.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, configurationnames.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
