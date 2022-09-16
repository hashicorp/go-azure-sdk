
## `github.com/hashicorp/go-azure-sdk/resource-manager/automanage/2022-05-04/serviceprincipals` Documentation

The `serviceprincipals` SDK allows for interaction with the Azure Resource Manager Service `automanage` (API Version `2022-05-04`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/automanage/2022-05-04/serviceprincipals"
```


### Client Initialization

```go
client := serviceprincipals.NewServicePrincipalsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ServicePrincipalsClient.ServicePrincipalsGet`

```go
ctx := context.TODO()
id := serviceprincipals.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.ServicePrincipalsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalsClient.ServicePrincipalsListBySubscription`

```go
ctx := context.TODO()
id := serviceprincipals.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.ServicePrincipalsListBySubscription(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
