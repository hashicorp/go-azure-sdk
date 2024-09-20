
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2023-05-01/connectedenvironmentsstorages` Documentation

The `connectedenvironmentsstorages` SDK allows for interaction with Azure Resource Manager `containerapps` (API Version `2023-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2023-05-01/connectedenvironmentsstorages"
```


### Client Initialization

```go
client := connectedenvironmentsstorages.NewConnectedEnvironmentsStoragesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ConnectedEnvironmentsStoragesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := connectedenvironmentsstorages.NewConnectedEnvironmentStorageID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectedEnvironmentName", "storageName")

payload := connectedenvironmentsstorages.ConnectedEnvironmentStorage{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConnectedEnvironmentsStoragesClient.Delete`

```go
ctx := context.TODO()
id := connectedenvironmentsstorages.NewConnectedEnvironmentStorageID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectedEnvironmentName", "storageName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConnectedEnvironmentsStoragesClient.Get`

```go
ctx := context.TODO()
id := connectedenvironmentsstorages.NewConnectedEnvironmentStorageID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectedEnvironmentName", "storageName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConnectedEnvironmentsStoragesClient.List`

```go
ctx := context.TODO()
id := connectedenvironmentsstorages.NewConnectedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectedEnvironmentName")

read, err := client.List(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
