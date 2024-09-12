
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/authenticationstrengthpolicy` Documentation

The `authenticationstrengthpolicy` SDK allows for interaction with the Azure Resource Manager Service `policies` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/authenticationstrengthpolicy"
```


### Client Initialization

```go
client := authenticationstrengthpolicy.NewAuthenticationStrengthPolicyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AuthenticationStrengthPolicyClient.CreateAuthenticationStrengthPolicy`

```go
ctx := context.TODO()

payload := authenticationstrengthpolicy.AuthenticationStrengthPolicy{
	// ...
}


read, err := client.CreateAuthenticationStrengthPolicy(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthenticationStrengthPolicyClient.DeleteAuthenticationStrengthPolicy`

```go
ctx := context.TODO()
id := authenticationstrengthpolicy.NewPolicyAuthenticationStrengthPolicyID("authenticationStrengthPolicyIdValue")

read, err := client.DeleteAuthenticationStrengthPolicy(ctx, id, authenticationstrengthpolicy.DefaultDeleteAuthenticationStrengthPolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthenticationStrengthPolicyClient.GetAuthenticationStrengthPoliciesCount`

```go
ctx := context.TODO()


read, err := client.GetAuthenticationStrengthPoliciesCount(ctx, authenticationstrengthpolicy.DefaultGetAuthenticationStrengthPoliciesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthenticationStrengthPolicyClient.GetAuthenticationStrengthPolicy`

```go
ctx := context.TODO()
id := authenticationstrengthpolicy.NewPolicyAuthenticationStrengthPolicyID("authenticationStrengthPolicyIdValue")

read, err := client.GetAuthenticationStrengthPolicy(ctx, id, authenticationstrengthpolicy.DefaultGetAuthenticationStrengthPolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthenticationStrengthPolicyClient.ListAuthenticationStrengthPolicies`

```go
ctx := context.TODO()


// alternatively `client.ListAuthenticationStrengthPolicies(ctx, authenticationstrengthpolicy.DefaultListAuthenticationStrengthPoliciesOperationOptions())` can be used to do batched pagination
items, err := client.ListAuthenticationStrengthPoliciesComplete(ctx, authenticationstrengthpolicy.DefaultListAuthenticationStrengthPoliciesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AuthenticationStrengthPolicyClient.UpdateAuthenticationStrengthPolicy`

```go
ctx := context.TODO()
id := authenticationstrengthpolicy.NewPolicyAuthenticationStrengthPolicyID("authenticationStrengthPolicyIdValue")

payload := authenticationstrengthpolicy.AuthenticationStrengthPolicy{
	// ...
}


read, err := client.UpdateAuthenticationStrengthPolicy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthenticationStrengthPolicyClient.UpdateAuthenticationStrengthPolicyAllowedCombination`

```go
ctx := context.TODO()
id := authenticationstrengthpolicy.NewPolicyAuthenticationStrengthPolicyID("authenticationStrengthPolicyIdValue")

payload := authenticationstrengthpolicy.UpdateAuthenticationStrengthPolicyAllowedCombinationRequest{
	// ...
}


read, err := client.UpdateAuthenticationStrengthPolicyAllowedCombination(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
