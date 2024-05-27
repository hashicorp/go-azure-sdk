
## `github.com/hashicorp/go-azure-sdk/resource-manager/networkanalytics/2023-11-15/dataproducts` Documentation

The `dataproducts` SDK allows for interaction with the Azure Resource Manager Service `networkanalytics` (API Version `2023-11-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/networkanalytics/2023-11-15/dataproducts"
```


### Client Initialization

```go
client := dataproducts.NewDataProductsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DataProductsClient.AddUserRole`

```go
ctx := context.TODO()
id := dataproducts.NewDataProductID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataProductValue")

payload := dataproducts.RoleAssignmentCommonProperties{
	// ...
}


read, err := client.AddUserRole(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataProductsClient.Create`

```go
ctx := context.TODO()
id := dataproducts.NewDataProductID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataProductValue")

payload := dataproducts.DataProduct{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DataProductsClient.Delete`

```go
ctx := context.TODO()
id := dataproducts.NewDataProductID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataProductValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DataProductsClient.GenerateStorageAccountSasToken`

```go
ctx := context.TODO()
id := dataproducts.NewDataProductID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataProductValue")

payload := dataproducts.AccountSas{
	// ...
}


read, err := client.GenerateStorageAccountSasToken(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataProductsClient.Get`

```go
ctx := context.TODO()
id := dataproducts.NewDataProductID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataProductValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataProductsClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DataProductsClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DataProductsClient.ListRolesAssignments`

```go
ctx := context.TODO()
id := dataproducts.NewDataProductID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataProductValue")
var payload interface{}

read, err := client.ListRolesAssignments(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataProductsClient.RemoveUserRole`

```go
ctx := context.TODO()
id := dataproducts.NewDataProductID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataProductValue")

payload := dataproducts.RoleAssignmentDetail{
	// ...
}


read, err := client.RemoveUserRole(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataProductsClient.RotateKey`

```go
ctx := context.TODO()
id := dataproducts.NewDataProductID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataProductValue")

payload := dataproducts.KeyVaultInfo{
	// ...
}


read, err := client.RotateKey(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataProductsClient.Update`

```go
ctx := context.TODO()
id := dataproducts.NewDataProductID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataProductValue")

payload := dataproducts.DataProductUpdate{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
