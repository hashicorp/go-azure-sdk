
## `github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2022-05-01/locations` Documentation

The `locations` SDK allows for interaction with the Azure Resource Manager Service `vmware` (API Version `2022-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2022-05-01/locations"
```


### Client Initialization

```go
client := locations.NewLocationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LocationsClient.CheckQuotaAvailability`

```go
ctx := context.TODO()
id := locations.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

read, err := client.CheckQuotaAvailability(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LocationsClient.CheckTrialAvailability`

```go
ctx := context.TODO()
id := locations.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

payload := locations.Sku{
	// ...
}


read, err := client.CheckTrialAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
