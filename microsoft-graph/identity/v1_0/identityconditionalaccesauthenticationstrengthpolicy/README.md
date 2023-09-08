
## `github.com/hashicorp/go-azure-sdk/resource-manager/identity/v1.0/identityconditionalaccesauthenticationstrengthpolicy` Documentation

The `identityconditionalaccesauthenticationstrengthpolicy` SDK allows for interaction with the Azure Resource Manager Service `identity` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/identity/v1.0/identityconditionalaccesauthenticationstrengthpolicy"
```


### Client Initialization

```go
client := identityconditionalaccesauthenticationstrengthpolicy.NewIdentityConditionalAccesAuthenticationStrengthPolicyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IdentityConditionalAccesAuthenticationStrengthPolicyClient.CreateIdentityConditionalAccesAuthenticationStrengthPolicy`

```go
ctx := context.TODO()

payload := identityconditionalaccesauthenticationstrengthpolicy.AuthenticationStrengthPolicy{
	// ...
}


read, err := client.CreateIdentityConditionalAccesAuthenticationStrengthPolicy(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentityConditionalAccesAuthenticationStrengthPolicyClient.DeleteIdentityConditionalAccesAuthenticationStrengthPolicyById`

```go
ctx := context.TODO()
id := identityconditionalaccesauthenticationstrengthpolicy.NewIdentityConditionalAccesAuthenticationStrengthPolicyID("authenticationStrengthPolicyIdValue")

read, err := client.DeleteIdentityConditionalAccesAuthenticationStrengthPolicyById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentityConditionalAccesAuthenticationStrengthPolicyClient.GetIdentityConditionalAccesAuthenticationStrengthPolicyById`

```go
ctx := context.TODO()
id := identityconditionalaccesauthenticationstrengthpolicy.NewIdentityConditionalAccesAuthenticationStrengthPolicyID("authenticationStrengthPolicyIdValue")

read, err := client.GetIdentityConditionalAccesAuthenticationStrengthPolicyById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentityConditionalAccesAuthenticationStrengthPolicyClient.GetIdentityConditionalAccesAuthenticationStrengthPolicyCount`

```go
ctx := context.TODO()


read, err := client.GetIdentityConditionalAccesAuthenticationStrengthPolicyCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentityConditionalAccesAuthenticationStrengthPolicyClient.ListIdentityConditionalAccesAuthenticationStrengthPolicies`

```go
ctx := context.TODO()


// alternatively `client.ListIdentityConditionalAccesAuthenticationStrengthPolicies(ctx)` can be used to do batched pagination
items, err := client.ListIdentityConditionalAccesAuthenticationStrengthPoliciesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IdentityConditionalAccesAuthenticationStrengthPolicyClient.UpdateIdentityConditionalAccesAuthenticationStrengthPolicyById`

```go
ctx := context.TODO()
id := identityconditionalaccesauthenticationstrengthpolicy.NewIdentityConditionalAccesAuthenticationStrengthPolicyID("authenticationStrengthPolicyIdValue")

payload := identityconditionalaccesauthenticationstrengthpolicy.AuthenticationStrengthPolicy{
	// ...
}


read, err := client.UpdateIdentityConditionalAccesAuthenticationStrengthPolicyById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentityConditionalAccesAuthenticationStrengthPolicyClient.UpdateIdentityConditionalAccesAuthenticationStrengthPolicyByIdAllowedCombination`

```go
ctx := context.TODO()
id := identityconditionalaccesauthenticationstrengthpolicy.NewIdentityConditionalAccesAuthenticationStrengthPolicyID("authenticationStrengthPolicyIdValue")

payload := identityconditionalaccesauthenticationstrengthpolicy.UpdateIdentityConditionalAccesAuthenticationStrengthPolicyByIdAllowedCombinationRequest{
	// ...
}


read, err := client.UpdateIdentityConditionalAccesAuthenticationStrengthPolicyByIdAllowedCombination(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
