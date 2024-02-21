
## `github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2024-02-01/extendedueinformation` Documentation

The `extendedueinformation` SDK allows for interaction with the Azure Resource Manager Service `mobilenetwork` (API Version `2024-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2024-02-01/extendedueinformation"
```


### Client Initialization

```go
client := extendedueinformation.NewExtendedUeInformationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ExtendedUeInformationClient.Get`

```go
ctx := context.TODO()
id := extendedueinformation.NewUeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "packetCoreControlPlaneValue", "ueIdValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
