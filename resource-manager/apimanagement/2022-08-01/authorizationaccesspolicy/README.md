
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/authorizationaccesspolicy` Documentation

The `authorizationaccesspolicy` SDK allows for interaction with the Azure Resource Manager Service `apimanagement` (API Version `2022-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/authorizationaccesspolicy"
```


### Client Initialization

```go
client := authorizationaccesspolicy.NewAuthorizationAccessPolicyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AuthorizationAccessPolicyClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := authorizationaccesspolicy.NewAccessPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "authorizationProviderIdValue", "authorizationIdValue", "authorizationAccessPolicyIdValue")

payload := authorizationaccesspolicy.AuthorizationAccessPolicyContract{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload, authorizationaccesspolicy.DefaultCreateOrUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthorizationAccessPolicyClient.Delete`

```go
ctx := context.TODO()
id := authorizationaccesspolicy.NewAccessPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "authorizationProviderIdValue", "authorizationIdValue", "authorizationAccessPolicyIdValue")

read, err := client.Delete(ctx, id, authorizationaccesspolicy.DefaultDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthorizationAccessPolicyClient.Get`

```go
ctx := context.TODO()
id := authorizationaccesspolicy.NewAccessPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "authorizationProviderIdValue", "authorizationIdValue", "authorizationAccessPolicyIdValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthorizationAccessPolicyClient.ListByAuthorization`

```go
ctx := context.TODO()
id := authorizationaccesspolicy.NewAuthorizationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "authorizationProviderIdValue", "authorizationIdValue")

// alternatively `client.ListByAuthorization(ctx, id, authorizationaccesspolicy.DefaultListByAuthorizationOperationOptions())` can be used to do batched pagination
items, err := client.ListByAuthorizationComplete(ctx, id, authorizationaccesspolicy.DefaultListByAuthorizationOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
