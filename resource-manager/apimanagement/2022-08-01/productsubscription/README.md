
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/productsubscription` Documentation

The `productsubscription` SDK allows for interaction with the Azure Resource Manager Service `apimanagement` (API Version `2022-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/productsubscription"
```


### Client Initialization

```go
client := productsubscription.NewProductSubscriptionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ProductSubscriptionClient.List`

```go
ctx := context.TODO()
id := productsubscription.NewProductID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "productIdValue")

// alternatively `client.List(ctx, id, productsubscription.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, productsubscription.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
