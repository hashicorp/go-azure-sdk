
## `github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-06-01-preview/flexibleservercapabilities` Documentation

The `flexibleservercapabilities` SDK allows for interaction with Azure Resource Manager `postgresql` (API Version `2023-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-06-01-preview/flexibleservercapabilities"
```


### Client Initialization

```go
client := flexibleservercapabilities.NewFlexibleServerCapabilitiesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `FlexibleServerCapabilitiesClient.ServerCapabilitiesList`

```go
ctx := context.TODO()
id := flexibleservercapabilities.NewFlexibleServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerName")

// alternatively `client.ServerCapabilitiesList(ctx, id)` can be used to do batched pagination
items, err := client.ServerCapabilitiesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
