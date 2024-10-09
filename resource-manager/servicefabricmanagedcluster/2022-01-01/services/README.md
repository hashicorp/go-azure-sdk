
## `github.com/hashicorp/go-azure-sdk/resource-manager/servicefabricmanagedcluster/2022-01-01/services` Documentation

The `services` SDK allows for interaction with Azure Resource Manager `servicefabricmanagedcluster` (API Version `2022-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/servicefabricmanagedcluster/2022-01-01/services"
```


### Client Initialization

```go
client := services.NewServicesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ServicesClient.Update`

```go
ctx := context.TODO()
id := services.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterName", "applicationName", "serviceName")

payload := services.ServiceUpdateParameters{
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
