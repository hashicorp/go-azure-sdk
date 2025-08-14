
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2025-02-01/openapis` Documentation

The `openapis` SDK allows for interaction with Azure Resource Manager `recoveryservices` (API Version `2025-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2025-02-01/openapis"
```


### Client Initialization

```go
client := openapis.NewOpenapisClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OpenapisClient.RecoveryServicesCapabilities`

```go
ctx := context.TODO()
id := openapis.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

payload := openapis.ResourceCapabilities{
	// ...
}


read, err := client.RecoveryServicesCapabilities(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.RecoveryServicesCheckNameAvailability`

```go
ctx := context.TODO()
id := openapis.NewProviderLocationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationName")

payload := openapis.CheckNameAvailabilityParameters{
	// ...
}


read, err := client.RecoveryServicesCheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
