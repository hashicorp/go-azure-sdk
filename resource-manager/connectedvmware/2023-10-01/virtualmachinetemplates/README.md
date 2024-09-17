
## `github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2023-10-01/virtualmachinetemplates` Documentation

The `virtualmachinetemplates` SDK allows for interaction with Azure Resource Manager `connectedvmware` (API Version `2023-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2023-10-01/virtualmachinetemplates"
```


### Client Initialization

```go
client := virtualmachinetemplates.NewVirtualMachineTemplatesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VirtualMachineTemplatesClient.Create`

```go
ctx := context.TODO()
id := virtualmachinetemplates.NewVirtualMachineTemplateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineTemplateValue")

payload := virtualmachinetemplates.VirtualMachineTemplate{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualMachineTemplatesClient.Delete`

```go
ctx := context.TODO()
id := virtualmachinetemplates.NewVirtualMachineTemplateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineTemplateValue")

if err := client.DeleteThenPoll(ctx, id, virtualmachinetemplates.DefaultDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualMachineTemplatesClient.Get`

```go
ctx := context.TODO()
id := virtualmachinetemplates.NewVirtualMachineTemplateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineTemplateValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualMachineTemplatesClient.List`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VirtualMachineTemplatesClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VirtualMachineTemplatesClient.Update`

```go
ctx := context.TODO()
id := virtualmachinetemplates.NewVirtualMachineTemplateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineTemplateValue")

payload := virtualmachinetemplates.ResourcePatch{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
