
## `github.com/hashicorp/go-azure-sdk/resource-manager/serviceprincipals/v1.0/serviceprincipalsynchronization` Documentation

The `serviceprincipalsynchronization` SDK allows for interaction with the Azure Resource Manager Service `serviceprincipals` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/serviceprincipals/v1.0/serviceprincipalsynchronization"
```


### Client Initialization

```go
client := serviceprincipalsynchronization.NewServicePrincipalSynchronizationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ServicePrincipalSynchronizationClient.AcquireServicePrincipalByIdSynchronizationAccessToken`

```go
ctx := context.TODO()
id := serviceprincipalsynchronization.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipalsynchronization.AcquireServicePrincipalByIdSynchronizationAccessTokenRequest{
	// ...
}


read, err := client.AcquireServicePrincipalByIdSynchronizationAccessToken(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalSynchronizationClient.CreateUpdateServicePrincipalByIdSynchronization`

```go
ctx := context.TODO()
id := serviceprincipalsynchronization.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipalsynchronization.Synchronization{
	// ...
}


read, err := client.CreateUpdateServicePrincipalByIdSynchronization(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalSynchronizationClient.DeleteServicePrincipalByIdSynchronization`

```go
ctx := context.TODO()
id := serviceprincipalsynchronization.NewServicePrincipalID("servicePrincipalIdValue")

read, err := client.DeleteServicePrincipalByIdSynchronization(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalSynchronizationClient.GetServicePrincipalByIdSynchronization`

```go
ctx := context.TODO()
id := serviceprincipalsynchronization.NewServicePrincipalID("servicePrincipalIdValue")

read, err := client.GetServicePrincipalByIdSynchronization(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
