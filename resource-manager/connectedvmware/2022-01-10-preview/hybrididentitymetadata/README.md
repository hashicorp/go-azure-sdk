
## `github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2022-01-10-preview/hybrididentitymetadata` Documentation

The `hybrididentitymetadata` SDK allows for interaction with Azure Resource Manager `connectedvmware` (API Version `2022-01-10-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2022-01-10-preview/hybrididentitymetadata"
```


### Client Initialization

```go
client := hybrididentitymetadata.NewHybridIdentityMetadataClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `HybridIdentityMetadataClient.Create`

```go
ctx := context.TODO()
id := hybrididentitymetadata.NewHybridIdentityMetadataID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName", "hybridIdentityMetadataName")

payload := hybrididentitymetadata.HybridIdentityMetadata{
	// ...
}


read, err := client.Create(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HybridIdentityMetadataClient.Delete`

```go
ctx := context.TODO()
id := hybrididentitymetadata.NewHybridIdentityMetadataID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName", "hybridIdentityMetadataName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HybridIdentityMetadataClient.Get`

```go
ctx := context.TODO()
id := hybrididentitymetadata.NewHybridIdentityMetadataID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName", "hybridIdentityMetadataName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HybridIdentityMetadataClient.ListByVM`

```go
ctx := context.TODO()
id := hybrididentitymetadata.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName")

// alternatively `client.ListByVM(ctx, id)` can be used to do batched pagination
items, err := client.ListByVMComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
