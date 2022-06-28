
## `github.com/hashicorp/go-azure-sdk/resource-manager/datalakestore/2016-11-01/locations` Documentation

The `locations` SDK allows for interaction with the Azure Resource Manager Service `datalakestore` (API Version `2016-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datalakestore/2016-11-01/locations"
```


### Client Initialization

```go
client := locations.NewLocationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LocationsClient.GetCapability`

```go
ctx := context.TODO()
id := locations.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

read, err := client.GetCapability(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
