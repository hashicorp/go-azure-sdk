
## `github.com/hashicorp/go-azure-sdk/resource-manager/policyinsights/2024-10-01/checkpolicyrestrictions` Documentation

The `checkpolicyrestrictions` SDK allows for interaction with Azure Resource Manager `policyinsights` (API Version `2024-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/policyinsights/2024-10-01/checkpolicyrestrictions"
```


### Client Initialization

```go
client := checkpolicyrestrictions.NewCheckPolicyRestrictionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CheckPolicyRestrictionsClient.PolicyRestrictionsCheckAtManagementGroupScope`

```go
ctx := context.TODO()
id := checkpolicyrestrictions.NewManagementGroupID("managementGroupId")

payload := checkpolicyrestrictions.CheckManagementGroupRestrictionsRequest{
	// ...
}


read, err := client.PolicyRestrictionsCheckAtManagementGroupScope(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CheckPolicyRestrictionsClient.PolicyRestrictionsCheckAtResourceGroupScope`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

payload := checkpolicyrestrictions.CheckRestrictionsRequest{
	// ...
}


read, err := client.PolicyRestrictionsCheckAtResourceGroupScope(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CheckPolicyRestrictionsClient.PolicyRestrictionsCheckAtSubscriptionScope`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

payload := checkpolicyrestrictions.CheckRestrictionsRequest{
	// ...
}


read, err := client.PolicyRestrictionsCheckAtSubscriptionScope(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
