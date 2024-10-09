
## `github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/servicerunners` Documentation

The `servicerunners` SDK allows for interaction with Azure Resource Manager `devtestlab` (API Version `2018-09-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/servicerunners"
```


### Client Initialization

```go
client := servicerunners.NewServiceRunnersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ServiceRunnersClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := servicerunners.NewServiceRunnerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "serviceRunnerName")

payload := servicerunners.ServiceRunner{
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


### Example Usage: `ServiceRunnersClient.Delete`

```go
ctx := context.TODO()
id := servicerunners.NewServiceRunnerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "serviceRunnerName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServiceRunnersClient.Get`

```go
ctx := context.TODO()
id := servicerunners.NewServiceRunnerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "serviceRunnerName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
