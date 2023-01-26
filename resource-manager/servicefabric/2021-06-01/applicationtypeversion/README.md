
## `github.com/hashicorp/go-azure-sdk/resource-manager/servicefabric/2021-06-01/applicationtypeversion` Documentation

The `applicationtypeversion` SDK allows for interaction with the Azure Resource Manager Service `servicefabric` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/servicefabric/2021-06-01/applicationtypeversion"
```


### Client Initialization

```go
client := applicationtypeversion.NewApplicationTypeVersionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ApplicationTypeVersionClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := applicationtypeversion.NewVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue", "applicationTypeValue", "versionValue")

payload := applicationtypeversion.ApplicationTypeVersionResource{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ApplicationTypeVersionClient.Delete`

```go
ctx := context.TODO()
id := applicationtypeversion.NewVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue", "applicationTypeValue", "versionValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ApplicationTypeVersionClient.Get`

```go
ctx := context.TODO()
id := applicationtypeversion.NewVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue", "applicationTypeValue", "versionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationTypeVersionClient.List`

```go
ctx := context.TODO()
id := applicationtypeversion.NewApplicationTypeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue", "applicationTypeValue")

read, err := client.List(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
