
## `github.com/hashicorp/go-azure-sdk/resource-manager/confidentialledger/2022-05-13/nameavailability` Documentation

The `nameavailability` SDK allows for interaction with the Azure Resource Manager Service `confidentialledger` (API Version `2022-05-13`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/confidentialledger/2022-05-13/nameavailability"
```


### Client Initialization

```go
client := nameavailability.NewNameAvailabilityClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `NameAvailabilityClient.CheckNameAvailability`

```go
ctx := context.TODO()
id := nameavailability.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

payload := nameavailability.CheckNameAvailabilityRequest{
	// ...
}


read, err := client.CheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
