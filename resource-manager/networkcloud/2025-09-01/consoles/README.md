
## `github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/consoles` Documentation

The `consoles` SDK allows for interaction with Azure Resource Manager `networkcloud` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/consoles"
```


### Client Initialization

```go
client := consoles.NewConsolesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ConsolesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := consoles.NewConsoleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName", "consoleName")

payload := consoles.Console{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload, consoles.DefaultCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `ConsolesClient.Delete`

```go
ctx := context.TODO()
id := consoles.NewConsoleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName", "consoleName")

if err := client.DeleteThenPoll(ctx, id, consoles.DefaultDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `ConsolesClient.Get`

```go
ctx := context.TODO()
id := consoles.NewConsoleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName", "consoleName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConsolesClient.ListByVirtualMachine`

```go
ctx := context.TODO()
id := consoles.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName")

// alternatively `client.ListByVirtualMachine(ctx, id, consoles.DefaultListByVirtualMachineOperationOptions())` can be used to do batched pagination
items, err := client.ListByVirtualMachineComplete(ctx, id, consoles.DefaultListByVirtualMachineOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ConsolesClient.Update`

```go
ctx := context.TODO()
id := consoles.NewConsoleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName", "consoleName")

payload := consoles.ConsolePatchParameters{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload, consoles.DefaultUpdateOperationOptions()); err != nil {
	// handle the error
}
```
