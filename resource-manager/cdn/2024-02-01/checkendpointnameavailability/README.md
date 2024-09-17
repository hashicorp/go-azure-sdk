
## `github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2024-02-01/checkendpointnameavailability` Documentation

The `checkendpointnameavailability` SDK allows for interaction with Azure Resource Manager `cdn` (API Version `2024-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2024-02-01/checkendpointnameavailability"
```


### Client Initialization

```go
client := checkendpointnameavailability.NewCheckEndpointNameAvailabilityClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CheckEndpointNameAvailabilityClient.CheckEndpointNameAvailability`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

payload := checkendpointnameavailability.CheckEndpointNameAvailabilityInput{
	// ...
}


read, err := client.CheckEndpointNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
