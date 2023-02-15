
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/apirevision` Documentation

The `apirevision` SDK allows for interaction with the Azure Resource Manager Service `apimanagement` (API Version `2022-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/apirevision"
```


### Client Initialization

```go
client := apirevision.NewApiRevisionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ApiRevisionClient.ListByService`

```go
ctx := context.TODO()
id := apirevision.NewApiID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "apiIdValue")

// alternatively `client.ListByService(ctx, id, apirevision.DefaultListByServiceOperationOptions())` can be used to do batched pagination
items, err := client.ListByServiceComplete(ctx, id, apirevision.DefaultListByServiceOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
