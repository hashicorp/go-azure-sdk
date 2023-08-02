
## `github.com/hashicorp/go-azure-sdk/resource-manager/managementgroups/2023-04-01/checknameavailability` Documentation

The `checknameavailability` SDK allows for interaction with the Azure Resource Manager Service `managementgroups` (API Version `2023-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/managementgroups/2023-04-01/checknameavailability"
```


### Client Initialization

```go
client := checknameavailability.NewCheckNameAvailabilityClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CheckNameAvailabilityClient.CheckNameAvailability`

```go
ctx := context.TODO()

payload := checknameavailability.CheckNameAvailabilityRequest{
	// ...
}


read, err := client.CheckNameAvailability(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
