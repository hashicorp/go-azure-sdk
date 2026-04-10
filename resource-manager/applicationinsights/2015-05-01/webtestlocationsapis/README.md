
## `github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2015-05-01/webtestlocationsapis` Documentation

The `webtestlocationsapis` SDK allows for interaction with Azure Resource Manager `applicationinsights` (API Version `2015-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2015-05-01/webtestlocationsapis"
```


### Client Initialization

```go
client := webtestlocationsapis.NewWebTestLocationsAPIsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WebTestLocationsAPIsClient.WebTestLocationsList`

```go
ctx := context.TODO()
id := webtestlocationsapis.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName")

// alternatively `client.WebTestLocationsList(ctx, id)` can be used to do batched pagination
items, err := client.WebTestLocationsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
