
## `github.com/hashicorp/go-azure-sdk/resource-manager/resources/2023-11-01/bicepclient` Documentation

The `bicepclient` SDK allows for interaction with the Azure Resource Manager Service `resources` (API Version `2023-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/resources/2023-11-01/bicepclient"
```


### Client Initialization

```go
client := bicepclient.NewBicepClientClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BicepClientClient.DecompileBicep`

```go
ctx := context.TODO()
id := bicepclient.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

payload := bicepclient.DecompileOperationRequest{
	// ...
}


read, err := client.DecompileBicep(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
