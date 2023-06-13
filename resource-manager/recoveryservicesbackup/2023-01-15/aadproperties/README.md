
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-01-15/aadproperties` Documentation

The `aadproperties` SDK allows for interaction with the Azure Resource Manager Service `recoveryservicesbackup` (API Version `2023-01-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-01-15/aadproperties"
```


### Client Initialization

```go
client := aadproperties.NewAadPropertiesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AadPropertiesClient.Get`

```go
ctx := context.TODO()
id := aadproperties.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

read, err := client.Get(ctx, id, aadproperties.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
