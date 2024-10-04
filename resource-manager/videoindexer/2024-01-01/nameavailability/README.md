
## `github.com/hashicorp/go-azure-sdk/resource-manager/videoindexer/2024-01-01/nameavailability` Documentation

The `nameavailability` SDK allows for interaction with Azure Resource Manager `videoindexer` (API Version `2024-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/videoindexer/2024-01-01/nameavailability"
```


### Client Initialization

```go
client := nameavailability.NewNameAvailabilityClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `NameAvailabilityClient.AccountsCheckNameAvailability`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

payload := nameavailability.AccountCheckNameAvailabilityParameters{
	// ...
}


read, err := client.AccountsCheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
