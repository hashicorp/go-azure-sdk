
## `github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/environments` Documentation

The `environments` SDK allows for interaction with Azure Resource Manager `devtestlab` (API Version `2018-09-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/environments"
```


### Client Initialization

```go
client := environments.NewEnvironmentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EnvironmentsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := environments.NewEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "userName", "name")

payload := environments.DtlEnvironment{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `EnvironmentsClient.Delete`

```go
ctx := context.TODO()
id := environments.NewEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "userName", "name")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `EnvironmentsClient.Get`

```go
ctx := context.TODO()
id := environments.NewEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "userName", "name")

read, err := client.Get(ctx, id, environments.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EnvironmentsClient.List`

```go
ctx := context.TODO()
id := environments.NewUserID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "userName")

// alternatively `client.List(ctx, id, environments.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, environments.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EnvironmentsClient.Update`

```go
ctx := context.TODO()
id := environments.NewEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "userName", "name")

payload := environments.UpdateResource{
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
