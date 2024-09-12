
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccessauthenticationstrengthpolicy` Documentation

The `conditionalaccessauthenticationstrengthpolicy` SDK allows for interaction with the Azure Resource Manager Service `identity` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccessauthenticationstrengthpolicy"
```


### Client Initialization

```go
client := conditionalaccessauthenticationstrengthpolicy.NewConditionalAccessAuthenticationStrengthPolicyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ConditionalAccessAuthenticationStrengthPolicyClient.CreateConditionalAccessAuthenticationStrengthPolicy`

```go
ctx := context.TODO()

payload := conditionalaccessauthenticationstrengthpolicy.AuthenticationStrengthPolicy{
	// ...
}


read, err := client.CreateConditionalAccessAuthenticationStrengthPolicy(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConditionalAccessAuthenticationStrengthPolicyClient.DeleteConditionalAccessAuthenticationStrengthPolicy`

```go
ctx := context.TODO()
id := conditionalaccessauthenticationstrengthpolicy.NewIdentityConditionalAccessAuthenticationStrengthPolicyID("authenticationStrengthPolicyIdValue")

read, err := client.DeleteConditionalAccessAuthenticationStrengthPolicy(ctx, id, conditionalaccessauthenticationstrengthpolicy.DefaultDeleteConditionalAccessAuthenticationStrengthPolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConditionalAccessAuthenticationStrengthPolicyClient.GetConditionalAccessAuthenticationStrengthPoliciesCount`

```go
ctx := context.TODO()


read, err := client.GetConditionalAccessAuthenticationStrengthPoliciesCount(ctx, conditionalaccessauthenticationstrengthpolicy.DefaultGetConditionalAccessAuthenticationStrengthPoliciesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConditionalAccessAuthenticationStrengthPolicyClient.GetConditionalAccessAuthenticationStrengthPolicy`

```go
ctx := context.TODO()
id := conditionalaccessauthenticationstrengthpolicy.NewIdentityConditionalAccessAuthenticationStrengthPolicyID("authenticationStrengthPolicyIdValue")

read, err := client.GetConditionalAccessAuthenticationStrengthPolicy(ctx, id, conditionalaccessauthenticationstrengthpolicy.DefaultGetConditionalAccessAuthenticationStrengthPolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConditionalAccessAuthenticationStrengthPolicyClient.ListConditionalAccessAuthenticationStrengthPolicies`

```go
ctx := context.TODO()


// alternatively `client.ListConditionalAccessAuthenticationStrengthPolicies(ctx, conditionalaccessauthenticationstrengthpolicy.DefaultListConditionalAccessAuthenticationStrengthPoliciesOperationOptions())` can be used to do batched pagination
items, err := client.ListConditionalAccessAuthenticationStrengthPoliciesComplete(ctx, conditionalaccessauthenticationstrengthpolicy.DefaultListConditionalAccessAuthenticationStrengthPoliciesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ConditionalAccessAuthenticationStrengthPolicyClient.UpdateConditionalAccessAuthenticationStrengthPolicy`

```go
ctx := context.TODO()
id := conditionalaccessauthenticationstrengthpolicy.NewIdentityConditionalAccessAuthenticationStrengthPolicyID("authenticationStrengthPolicyIdValue")

payload := conditionalaccessauthenticationstrengthpolicy.AuthenticationStrengthPolicy{
	// ...
}


read, err := client.UpdateConditionalAccessAuthenticationStrengthPolicy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConditionalAccessAuthenticationStrengthPolicyClient.UpdateConditionalAccessAuthenticationStrengthPolicyAllowedCombination`

```go
ctx := context.TODO()
id := conditionalaccessauthenticationstrengthpolicy.NewIdentityConditionalAccessAuthenticationStrengthPolicyID("authenticationStrengthPolicyIdValue")

payload := conditionalaccessauthenticationstrengthpolicy.UpdateConditionalAccessAuthenticationStrengthPolicyAllowedCombinationRequest{
	// ...
}


read, err := client.UpdateConditionalAccessAuthenticationStrengthPolicyAllowedCombination(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
