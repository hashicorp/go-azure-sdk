
## `github.com/hashicorp/go-azure-sdk/resource-manager/frontdoor/2020-05-01/checkfrontdoornameavailability` Documentation

The `checkfrontdoornameavailability` SDK allows for interaction with Azure Resource Manager `frontdoor` (API Version `2020-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/frontdoor/2020-05-01/checkfrontdoornameavailability"
```


### Client Initialization

```go
client := checkfrontdoornameavailability.NewCheckFrontDoorNameAvailabilityClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CheckFrontDoorNameAvailabilityClient.FrontDoorNameAvailabilityCheck`

```go
ctx := context.TODO()

payload := checkfrontdoornameavailability.CheckNameAvailabilityInput{
	// ...
}


read, err := client.FrontDoorNameAvailabilityCheck(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
