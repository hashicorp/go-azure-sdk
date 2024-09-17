
## `github.com/hashicorp/go-azure-sdk/resource-manager/informatica/2024-05-08/serverlessruntimes` Documentation

The `serverlessruntimes` SDK allows for interaction with Azure Resource Manager `informatica` (API Version `2024-05-08`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/informatica/2024-05-08/serverlessruntimes"
```


### Client Initialization

```go
client := serverlessruntimes.NewServerlessRuntimesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ServerlessRuntimesClient.CheckDependencies`

```go
ctx := context.TODO()
id := serverlessruntimes.NewServerlessRuntimeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "organizationValue", "serverlessRuntimeValue")

read, err := client.CheckDependencies(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServerlessRuntimesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := serverlessruntimes.NewServerlessRuntimeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "organizationValue", "serverlessRuntimeValue")

payload := serverlessruntimes.InformaticaServerlessRuntimeResource{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ServerlessRuntimesClient.Delete`

```go
ctx := context.TODO()
id := serverlessruntimes.NewServerlessRuntimeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "organizationValue", "serverlessRuntimeValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ServerlessRuntimesClient.Get`

```go
ctx := context.TODO()
id := serverlessruntimes.NewServerlessRuntimeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "organizationValue", "serverlessRuntimeValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServerlessRuntimesClient.ListByInformaticaOrganizationResource`

```go
ctx := context.TODO()
id := serverlessruntimes.NewOrganizationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "organizationValue")

// alternatively `client.ListByInformaticaOrganizationResource(ctx, id)` can be used to do batched pagination
items, err := client.ListByInformaticaOrganizationResourceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ServerlessRuntimesClient.ServerlessResourceById`

```go
ctx := context.TODO()
id := serverlessruntimes.NewServerlessRuntimeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "organizationValue", "serverlessRuntimeValue")

read, err := client.ServerlessResourceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServerlessRuntimesClient.StartFailedServerlessRuntime`

```go
ctx := context.TODO()
id := serverlessruntimes.NewServerlessRuntimeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "organizationValue", "serverlessRuntimeValue")

read, err := client.StartFailedServerlessRuntime(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServerlessRuntimesClient.Update`

```go
ctx := context.TODO()
id := serverlessruntimes.NewServerlessRuntimeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "organizationValue", "serverlessRuntimeValue")

payload := serverlessruntimes.InformaticaServerlessRuntimeResourceUpdate{
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
