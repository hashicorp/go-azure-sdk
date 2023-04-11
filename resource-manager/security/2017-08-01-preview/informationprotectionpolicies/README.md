
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2017-08-01-preview/informationprotectionpolicies` Documentation

The `informationprotectionpolicies` SDK allows for interaction with the Azure Resource Manager Service `security` (API Version `2017-08-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2017-08-01-preview/informationprotectionpolicies"
```


### Client Initialization

```go
client := informationprotectionpolicies.NewInformationProtectionPoliciesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InformationProtectionPoliciesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := informationprotectionpolicies.NewScopedInformationProtectionPolicyID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "custom")

payload := informationprotectionpolicies.InformationProtectionPolicy{
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


### Example Usage: `InformationProtectionPoliciesClient.Get`

```go
ctx := context.TODO()
id := informationprotectionpolicies.NewScopedInformationProtectionPolicyID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "custom")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionPoliciesClient.List`

```go
ctx := context.TODO()
id := informationprotectionpolicies.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
