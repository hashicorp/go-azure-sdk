
## `github.com/hashicorp/go-azure-sdk/resource-manager/operationsmanagement/2015-11-01-preview/managementassociation` Documentation

The `managementassociation` SDK allows for interaction with the Azure Resource Manager Service `operationsmanagement` (API Version `2015-11-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/operationsmanagement/2015-11-01-preview/managementassociation"
```


### Client Initialization

```go
client := managementassociation.NewManagementAssociationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagementAssociationClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := managementassociation.NewScopedManagementAssociationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "managementAssociationValue")

payload := managementassociation.ManagementAssociation{
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


### Example Usage: `ManagementAssociationClient.Delete`

```go
ctx := context.TODO()
id := managementassociation.NewScopedManagementAssociationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "managementAssociationValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagementAssociationClient.Get`

```go
ctx := context.TODO()
id := managementassociation.NewScopedManagementAssociationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "managementAssociationValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagementAssociationClient.ListBySubscription`

```go
ctx := context.TODO()
id := managementassociation.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.ListBySubscription(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
