
## `github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2024-02-01/checknameavailabilitywithsubscription` Documentation

The `checknameavailabilitywithsubscription` SDK allows for interaction with Azure Resource Manager `cdn` (API Version `2024-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2024-02-01/checknameavailabilitywithsubscription"
```


### Client Initialization

```go
client := checknameavailabilitywithsubscription.NewCheckNameAvailabilityWithSubscriptionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CheckNameAvailabilityWithSubscriptionClient.CheckNameAvailabilityWithSubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

payload := checknameavailabilitywithsubscription.CheckNameAvailabilityInput{
	// ...
}


read, err := client.CheckNameAvailabilityWithSubscription(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
