
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-10-01/datareference` Documentation

The `datareference` SDK allows for interaction with Azure Resource Manager `machinelearningservices` (API Version `2023-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-10-01/datareference"
```


### Client Initialization

```go
client := datareference.NewDataReferenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DataReferenceClient.RegistryDataReferencesGetBlobReferenceSAS`

```go
ctx := context.TODO()
id := datareference.NewDataReferenceVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "registryName", "name", "version")

payload := datareference.GetBlobReferenceSASRequestDto{
	// ...
}


read, err := client.RegistryDataReferencesGetBlobReferenceSAS(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
