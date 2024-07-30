
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/synchronizationsecret` Documentation

The `synchronizationsecret` SDK allows for interaction with the Azure Resource Manager Service `serviceprincipals` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/synchronizationsecret"
```


### Client Initialization

```go
client := synchronizationsecret.NewSynchronizationSecretClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SynchronizationSecretClient.CreateUpdateSynchronizationSecret`

```go
ctx := context.TODO()
id := synchronizationsecret.NewServicePrincipalID("servicePrincipalIdValue")

payload := synchronizationsecret.CreateUpdateSynchronizationSecretRequest{
	// ...
}


read, err := client.CreateUpdateSynchronizationSecret(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationSecretClient.GetSynchronizationSecretCount`

```go
ctx := context.TODO()
id := synchronizationsecret.NewServicePrincipalID("servicePrincipalIdValue")

read, err := client.GetSynchronizationSecretCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
