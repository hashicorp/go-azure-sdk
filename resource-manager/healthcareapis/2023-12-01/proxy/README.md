
## `github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2023-12-01/proxy` Documentation

The `proxy` SDK allows for interaction with the Azure Resource Manager Service `healthcareapis` (API Version `2023-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2023-12-01/proxy"
```


### Client Initialization

```go
client := proxy.NewProxyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ProxyClient.ServicesCheckNameAvailability`

```go
ctx := context.TODO()
id := proxy.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

payload := proxy.CheckNameAvailabilityParameters{
	// ...
}


read, err := client.ServicesCheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
