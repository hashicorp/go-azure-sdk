
## `github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/testraiexternalsafetyprovider` Documentation

The `testraiexternalsafetyprovider` SDK allows for interaction with Azure Resource Manager `cognitive` (API Version `2025-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/testraiexternalsafetyprovider"
```


### Client Initialization

```go
client := testraiexternalsafetyprovider.NewTestRaiExternalSafetyProviderClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TestRaiExternalSafetyProviderClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := testraiexternalsafetyprovider.NewTestRaiExternalSafetyProviderID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "testRaiExternalSafetyProviderName")

payload := testraiexternalsafetyprovider.RaiExternalSafetyProviderSchema{
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
