
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-06-01/featuresupport` Documentation

The `featuresupport` SDK allows for interaction with the Azure Resource Manager Service `recoveryservicesbackup` (API Version `2023-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-06-01/featuresupport"
```


### Client Initialization

```go
client := featuresupport.NewFeatureSupportClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `FeatureSupportClient.Validate`

```go
ctx := context.TODO()
id := featuresupport.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

payload := featuresupport.FeatureSupportRequest{
	// ...
}


read, err := client.Validate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
