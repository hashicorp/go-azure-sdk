
## `github.com/hashicorp/go-azure-sdk/resource-manager/resources/2022-06-01/resourcevalidator` Documentation

The `resourcevalidator` SDK allows for interaction with Azure Resource Manager `resources` (API Version `2022-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/resources/2022-06-01/resourcevalidator"
```


### Client Initialization

```go
client := resourcevalidator.NewResourceValidatorClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ResourceValidatorClient.ValidateResources`

```go
ctx := context.TODO()

payload := resourcevalidator.ResourceValidationRequest{
	// ...
}


read, err := client.ValidateResources(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
