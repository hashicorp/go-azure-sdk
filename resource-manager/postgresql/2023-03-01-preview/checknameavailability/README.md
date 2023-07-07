
## `github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-03-01-preview/checknameavailability` Documentation

The `checknameavailability` SDK allows for interaction with the Azure Resource Manager Service `postgresql` (API Version `2023-03-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-03-01-preview/checknameavailability"
```


### Client Initialization

```go
client := checknameavailability.NewCheckNameAvailabilityClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CheckNameAvailabilityClient.Execute`

```go
ctx := context.TODO()
id := checknameavailability.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

payload := checknameavailability.CheckNameAvailabilityRequest{
	// ...
}


read, err := client.Execute(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CheckNameAvailabilityClient.WithLocationExecute`

```go
ctx := context.TODO()
id := checknameavailability.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

payload := checknameavailability.CheckNameAvailabilityRequest{
	// ...
}


read, err := client.WithLocationExecute(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
