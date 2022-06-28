
## `github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-01-01-preview/checknameavailabilitynamespaces` Documentation

The `checknameavailabilitynamespaces` SDK allows for interaction with the Azure Resource Manager Service `eventhub` (API Version `2021-01-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-01-01-preview/checknameavailabilitynamespaces"
```


### Client Initialization

```go
client := checknameavailabilitynamespaces.NewCheckNameAvailabilityNamespacesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CheckNameAvailabilityNamespacesClient.NamespacesCheckNameAvailability`

```go
ctx := context.TODO()
id := checknameavailabilitynamespaces.NewSubscriptionID()

payload := checknameavailabilitynamespaces.CheckNameAvailabilityParameter{
	// ...
}


read, err := client.NamespacesCheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
