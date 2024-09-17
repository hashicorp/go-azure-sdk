
## `github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/virtualmachineextensionimages` Documentation

The `virtualmachineextensionimages` SDK allows for interaction with Azure Resource Manager `compute` (API Version `2021-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/virtualmachineextensionimages"
```


### Client Initialization

```go
client := virtualmachineextensionimages.NewVirtualMachineExtensionImagesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VirtualMachineExtensionImagesClient.Get`

```go
ctx := context.TODO()
id := virtualmachineextensionimages.NewVersionID("12345678-1234-9876-4563-123456789012", "locationValue", "publisherValue", "typeValue", "versionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualMachineExtensionImagesClient.ListTypes`

```go
ctx := context.TODO()
id := virtualmachineextensionimages.NewPublisherID("12345678-1234-9876-4563-123456789012", "locationValue", "publisherValue")

read, err := client.ListTypes(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualMachineExtensionImagesClient.ListVersions`

```go
ctx := context.TODO()
id := virtualmachineextensionimages.NewTypeID("12345678-1234-9876-4563-123456789012", "locationValue", "publisherValue", "typeValue")

read, err := client.ListVersions(ctx, id, virtualmachineextensionimages.DefaultListVersionsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
