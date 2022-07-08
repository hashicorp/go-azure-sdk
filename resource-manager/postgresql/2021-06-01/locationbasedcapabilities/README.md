
## `github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2021-06-01/locationbasedcapabilities` Documentation

The `locationbasedcapabilities` SDK allows for interaction with the Azure Resource Manager Service `postgresql` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2021-06-01/locationbasedcapabilities"
```


### Client Initialization

```go
client := locationbasedcapabilities.NewLocationBasedCapabilitiesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LocationBasedCapabilitiesClient.Execute`

```go
ctx := context.TODO()
id := locationbasedcapabilities.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

// alternatively `client.Execute(ctx, id)` can be used to do batched pagination
items, err := client.ExecuteComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
