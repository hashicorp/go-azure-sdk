
## `github.com/hashicorp/go-azure-sdk/resource-manager/frontdoor/2020-04-01/checkfrontdoornameavailabilitywithsubscription` Documentation

The `checkfrontdoornameavailabilitywithsubscription` SDK allows for interaction with the Azure Resource Manager Service `frontdoor` (API Version `2020-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/frontdoor/2020-04-01/checkfrontdoornameavailabilitywithsubscription"
```


### Client Initialization

```go
client := checkfrontdoornameavailabilitywithsubscription.NewCheckFrontDoorNameAvailabilityWithSubscriptionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CheckFrontDoorNameAvailabilityWithSubscriptionClient.CheckFrontDoorNameAvailabilityWithSubscription`

```go
ctx := context.TODO()
id := checkfrontdoornameavailabilitywithsubscription.NewSubscriptionID()

payload := checkfrontdoornameavailabilitywithsubscription.CheckNameAvailabilityInput{
	// ...
}


read, err := client.CheckFrontDoorNameAvailabilityWithSubscription(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
