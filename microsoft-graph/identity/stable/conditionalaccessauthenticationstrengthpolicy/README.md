
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccessauthenticationstrengthpolicy` Documentation

The `conditionalaccessauthenticationstrengthpolicy` SDK allows for interaction with Microsoft Graph `identity` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccessauthenticationstrengthpolicy"
```


### Client Initialization

```go
client := conditionalaccessauthenticationstrengthpolicy.NewConditionalAccessAuthenticationStrengthPolicyClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ConditionalAccessAuthenticationStrengthPolicyClient.CreateConditionalAccessAuthenticationStrengthPolicy`

```go
ctx := context.TODO()

payload := conditionalaccessauthenticationstrengthpolicy.AuthenticationStrengthPolicy{
	// ...
}


read, err := client.CreateConditionalAccessAuthenticationStrengthPolicy(ctx, payload, conditionalaccessauthenticationstrengthpolicy.DefaultCreateConditionalAccessAuthenticationStrengthPolicyOperationOptions())
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
id := conditionalaccessauthenticationstrengthpolicy.NewIdentityConditionalAccessAuthenticationStrengthPolicyID("authenticationStrengthPolicyId")

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
id := conditionalaccessauthenticationstrengthpolicy.NewIdentityConditionalAccessAuthenticationStrengthPolicyID("authenticationStrengthPolicyId")

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
id := conditionalaccessauthenticationstrengthpolicy.NewIdentityConditionalAccessAuthenticationStrengthPolicyID("authenticationStrengthPolicyId")

payload := conditionalaccessauthenticationstrengthpolicy.AuthenticationStrengthPolicy{
	// ...
}


read, err := client.UpdateConditionalAccessAuthenticationStrengthPolicy(ctx, id, payload, conditionalaccessauthenticationstrengthpolicy.DefaultUpdateConditionalAccessAuthenticationStrengthPolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConditionalAccessAuthenticationStrengthPolicyClient.UpdateConditionalAccessAuthenticationStrengthPolicyAllowedCombinations`

```go
ctx := context.TODO()
id := conditionalaccessauthenticationstrengthpolicy.NewIdentityConditionalAccessAuthenticationStrengthPolicyID("authenticationStrengthPolicyId")

payload := conditionalaccessauthenticationstrengthpolicy.UpdateConditionalAccessAuthenticationStrengthPolicyAllowedCombinationsRequest{
	// ...
}


read, err := client.UpdateConditionalAccessAuthenticationStrengthPolicyAllowedCombinations(ctx, id, payload, conditionalaccessauthenticationstrengthpolicy.DefaultUpdateConditionalAccessAuthenticationStrengthPolicyAllowedCombinationsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
