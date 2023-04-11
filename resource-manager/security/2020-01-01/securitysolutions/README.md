
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/securitysolutions` Documentation

The `securitysolutions` SDK allows for interaction with the Azure Resource Manager Service `security` (API Version `2020-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/securitysolutions"
```


### Client Initialization

```go
client := securitysolutions.NewSecuritySolutionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SecuritySolutionsClient.Get`

```go
ctx := context.TODO()
id := securitysolutions.NewSecuritySolutionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationValue", "securitySolutionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SecuritySolutionsClient.List`

```go
ctx := context.TODO()
id := securitysolutions.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
