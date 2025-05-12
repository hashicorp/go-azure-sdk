
## `github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-04-15/checknameavailability` Documentation

The `checknameavailability` SDK allows for interaction with Azure Resource Manager `cdn` (API Version `2025-04-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-04-15/checknameavailability"
```


### Client Initialization

```go
client := checknameavailability.NewCheckNameAvailabilityClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CheckNameAvailabilityClient.CheckNameAvailability`

```go
ctx := context.TODO()

payload := checknameavailability.CheckNameAvailabilityInput{
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
