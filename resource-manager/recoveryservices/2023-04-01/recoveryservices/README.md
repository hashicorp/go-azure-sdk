
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2023-04-01/recoveryservices` Documentation

The `recoveryservices` SDK allows for interaction with the Azure Resource Manager Service `recoveryservices` (API Version `2023-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2023-04-01/recoveryservices"
```


### Client Initialization

```go
client := recoveryservices.NewRecoveryServicesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RecoveryServicesClient.Capabilities`

```go
ctx := context.TODO()
id := recoveryservices.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

payload := recoveryservices.ResourceCapabilities{
	// ...
}


read, err := client.Capabilities(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecoveryServicesClient.CheckNameAvailability`

```go
ctx := context.TODO()
id := recoveryservices.NewProviderLocationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationValue")

payload := recoveryservices.CheckNameAvailabilityParameters{
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
