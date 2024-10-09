
## `github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2019-11-01/emailregistration` Documentation

The `emailregistration` SDK allows for interaction with Azure Resource Manager `datashare` (API Version `2019-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2019-11-01/emailregistration"
```


### Client Initialization

```go
client := emailregistration.NewEmailRegistrationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EmailRegistrationClient.ActivateEmail`

```go
ctx := context.TODO()
id := emailregistration.NewLocationID("locationName")

payload := emailregistration.EmailRegistration{
	// ...
}


read, err := client.ActivateEmail(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EmailRegistrationClient.RegisterEmail`

```go
ctx := context.TODO()
id := emailregistration.NewLocationID("locationName")

read, err := client.RegisterEmail(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
