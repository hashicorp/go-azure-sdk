
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccesauthenticationstrengthpolicy` Documentation

The `conditionalaccesauthenticationstrengthpolicy` SDK allows for interaction with the Azure Resource Manager Service `identity` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccesauthenticationstrengthpolicy"
```


### Client Initialization

```go
client := conditionalaccesauthenticationstrengthpolicy.NewConditionalAccesAuthenticationStrengthPolicyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ConditionalAccesAuthenticationStrengthPolicyClient.CreateConditionalAccesAuthenticationStrengthPolicy`

```go
ctx := context.TODO()

payload := conditionalaccesauthenticationstrengthpolicy.AuthenticationStrengthPolicy{
	// ...
}


read, err := client.CreateConditionalAccesAuthenticationStrengthPolicy(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConditionalAccesAuthenticationStrengthPolicyClient.DeleteConditionalAccesAuthenticationStrengthPolicy`

```go
ctx := context.TODO()
id := conditionalaccesauthenticationstrengthpolicy.NewIdentityConditionalAccesAuthenticationStrengthPolicyID("authenticationStrengthPolicyIdValue")

read, err := client.DeleteConditionalAccesAuthenticationStrengthPolicy(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConditionalAccesAuthenticationStrengthPolicyClient.GetConditionalAccesAuthenticationStrengthPolicy`

```go
ctx := context.TODO()
id := conditionalaccesauthenticationstrengthpolicy.NewIdentityConditionalAccesAuthenticationStrengthPolicyID("authenticationStrengthPolicyIdValue")

read, err := client.GetConditionalAccesAuthenticationStrengthPolicy(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConditionalAccesAuthenticationStrengthPolicyClient.GetConditionalAccesAuthenticationStrengthPolicyCount`

```go
ctx := context.TODO()


read, err := client.GetConditionalAccesAuthenticationStrengthPolicyCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConditionalAccesAuthenticationStrengthPolicyClient.ListConditionalAccesAuthenticationStrengthPolicies`

```go
ctx := context.TODO()


// alternatively `client.ListConditionalAccesAuthenticationStrengthPolicies(ctx)` can be used to do batched pagination
items, err := client.ListConditionalAccesAuthenticationStrengthPoliciesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ConditionalAccesAuthenticationStrengthPolicyClient.UpdateConditionalAccesAuthenticationStrengthPolicy`

```go
ctx := context.TODO()
id := conditionalaccesauthenticationstrengthpolicy.NewIdentityConditionalAccesAuthenticationStrengthPolicyID("authenticationStrengthPolicyIdValue")

payload := conditionalaccesauthenticationstrengthpolicy.AuthenticationStrengthPolicy{
	// ...
}


read, err := client.UpdateConditionalAccesAuthenticationStrengthPolicy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConditionalAccesAuthenticationStrengthPolicyClient.UpdateIdentityConditionalAccesAuthenticationStrengthPolicyAllowedCombination`

```go
ctx := context.TODO()
id := conditionalaccesauthenticationstrengthpolicy.NewIdentityConditionalAccesAuthenticationStrengthPolicyID("authenticationStrengthPolicyIdValue")

payload := conditionalaccesauthenticationstrengthpolicy.UpdateIdentityConditionalAccesAuthenticationStrengthPolicyAllowedCombinationRequest{
	// ...
}


read, err := client.UpdateIdentityConditionalAccesAuthenticationStrengthPolicyAllowedCombination(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
