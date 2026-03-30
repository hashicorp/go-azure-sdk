
## `github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/baremetalmachines` Documentation

The `baremetalmachines` SDK allows for interaction with Azure Resource Manager `networkcloud` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/baremetalmachines"
```


### Client Initialization

```go
client := baremetalmachines.NewBareMetalMachinesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BareMetalMachinesClient.Cordon`

```go
ctx := context.TODO()
id := baremetalmachines.NewBareMetalMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bareMetalMachineName")

payload := baremetalmachines.BareMetalMachineCordonParameters{
	// ...
}


if err := client.CordonThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BareMetalMachinesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := baremetalmachines.NewBareMetalMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bareMetalMachineName")

payload := baremetalmachines.BareMetalMachine{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload, baremetalmachines.DefaultCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `BareMetalMachinesClient.Delete`

```go
ctx := context.TODO()
id := baremetalmachines.NewBareMetalMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bareMetalMachineName")

if err := client.DeleteThenPoll(ctx, id, baremetalmachines.DefaultDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `BareMetalMachinesClient.Get`

```go
ctx := context.TODO()
id := baremetalmachines.NewBareMetalMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bareMetalMachineName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BareMetalMachinesClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id, baremetalmachines.DefaultListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id, baremetalmachines.DefaultListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BareMetalMachinesClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id, baremetalmachines.DefaultListBySubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id, baremetalmachines.DefaultListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BareMetalMachinesClient.PowerOff`

```go
ctx := context.TODO()
id := baremetalmachines.NewBareMetalMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bareMetalMachineName")

payload := baremetalmachines.BareMetalMachinePowerOffParameters{
	// ...
}


if err := client.PowerOffThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BareMetalMachinesClient.Reimage`

```go
ctx := context.TODO()
id := baremetalmachines.NewBareMetalMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bareMetalMachineName")

if err := client.ReimageThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `BareMetalMachinesClient.Replace`

```go
ctx := context.TODO()
id := baremetalmachines.NewBareMetalMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bareMetalMachineName")

payload := baremetalmachines.BareMetalMachineReplaceParameters{
	// ...
}


if err := client.ReplaceThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BareMetalMachinesClient.Restart`

```go
ctx := context.TODO()
id := baremetalmachines.NewBareMetalMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bareMetalMachineName")

if err := client.RestartThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `BareMetalMachinesClient.RunCommand`

```go
ctx := context.TODO()
id := baremetalmachines.NewBareMetalMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bareMetalMachineName")

payload := baremetalmachines.BareMetalMachineRunCommandParameters{
	// ...
}


if err := client.RunCommandThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BareMetalMachinesClient.RunDataExtracts`

```go
ctx := context.TODO()
id := baremetalmachines.NewBareMetalMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bareMetalMachineName")

payload := baremetalmachines.BareMetalMachineRunDataExtractsParameters{
	// ...
}


if err := client.RunDataExtractsThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BareMetalMachinesClient.RunDataExtractsRestricted`

```go
ctx := context.TODO()
id := baremetalmachines.NewBareMetalMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bareMetalMachineName")

payload := baremetalmachines.BareMetalMachineRunDataExtractsParameters{
	// ...
}


if err := client.RunDataExtractsRestrictedThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BareMetalMachinesClient.RunReadCommands`

```go
ctx := context.TODO()
id := baremetalmachines.NewBareMetalMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bareMetalMachineName")

payload := baremetalmachines.BareMetalMachineRunReadCommandsParameters{
	// ...
}


if err := client.RunReadCommandsThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BareMetalMachinesClient.Start`

```go
ctx := context.TODO()
id := baremetalmachines.NewBareMetalMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bareMetalMachineName")

if err := client.StartThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `BareMetalMachinesClient.Uncordon`

```go
ctx := context.TODO()
id := baremetalmachines.NewBareMetalMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bareMetalMachineName")

if err := client.UncordonThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `BareMetalMachinesClient.Update`

```go
ctx := context.TODO()
id := baremetalmachines.NewBareMetalMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bareMetalMachineName")

payload := baremetalmachines.BareMetalMachinePatchParameters{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload, baremetalmachines.DefaultUpdateOperationOptions()); err != nil {
	// handle the error
}
```
