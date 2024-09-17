
## `github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2022-08-01/virtualmachine` Documentation

The `virtualmachine` SDK allows for interaction with Azure Resource Manager `labservices` (API Version `2022-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2022-08-01/virtualmachine"
```


### Client Initialization

```go
client := virtualmachine.NewVirtualMachineClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VirtualMachineClient.Get`

```go
ctx := context.TODO()
id := virtualmachine.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "virtualMachineValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualMachineClient.LabPlansSaveImage`

```go
ctx := context.TODO()
id := virtualmachine.NewLabPlanID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labPlanValue")

payload := virtualmachine.SaveImageBody{
	// ...
}


if err := client.LabPlansSaveImageThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualMachineClient.ListByLab`

```go
ctx := context.TODO()
id := virtualmachine.NewLabID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue")

// alternatively `client.ListByLab(ctx, id)` can be used to do batched pagination
items, err := client.ListByLabComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VirtualMachineClient.Redeploy`

```go
ctx := context.TODO()
id := virtualmachine.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "virtualMachineValue")

if err := client.RedeployThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualMachineClient.Reimage`

```go
ctx := context.TODO()
id := virtualmachine.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "virtualMachineValue")

if err := client.ReimageThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualMachineClient.ResetPassword`

```go
ctx := context.TODO()
id := virtualmachine.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "virtualMachineValue")

payload := virtualmachine.ResetPasswordBody{
	// ...
}


if err := client.ResetPasswordThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualMachineClient.Start`

```go
ctx := context.TODO()
id := virtualmachine.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "virtualMachineValue")

if err := client.StartThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualMachineClient.Stop`

```go
ctx := context.TODO()
id := virtualmachine.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "virtualMachineValue")

if err := client.StopThenPoll(ctx, id); err != nil {
	// handle the error
}
```
