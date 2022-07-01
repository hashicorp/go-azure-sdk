
## `github.com/hashicorp/go-azure-sdk/resource-manager/aad/2021-03-01/oucontainer` Documentation

The `oucontainer` SDK allows for interaction with the Azure Resource Manager Service `aad` (API Version `2021-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/aad/2021-03-01/oucontainer"
```


### Client Initialization

```go
client := oucontainer.NewOuContainerClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OuContainerClient.Create`

```go
ctx := context.TODO()
id := oucontainer.NewOuContainerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "domainServiceValue", "ouContainerValue")

payload := oucontainer.ContainerAccount{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `OuContainerClient.Delete`

```go
ctx := context.TODO()
id := oucontainer.NewOuContainerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "domainServiceValue", "ouContainerValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `OuContainerClient.Get`

```go
ctx := context.TODO()
id := oucontainer.NewOuContainerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "domainServiceValue", "ouContainerValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OuContainerClient.List`

```go
ctx := context.TODO()
id := oucontainer.NewDomainServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "domainServiceValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OuContainerClient.Update`

```go
ctx := context.TODO()
id := oucontainer.NewOuContainerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "domainServiceValue", "ouContainerValue")

payload := oucontainer.ContainerAccount{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
