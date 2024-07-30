
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/authenticationstrengthpolicy` Documentation

The `authenticationstrengthpolicy` SDK allows for interaction with the Azure Resource Manager Service `policies` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/authenticationstrengthpolicy"
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

read, err := client.DeleteAuthenticationStrengthPolicy(ctx, id)
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

read, err := client.GetAuthenticationStrengthPolicy(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthenticationStrengthPolicyClient.GetAuthenticationStrengthPolicyCount`

```go
ctx := context.TODO()


read, err := client.GetAuthenticationStrengthPolicyCount(ctx)
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


// alternatively `client.ListAuthenticationStrengthPolicies(ctx)` can be used to do batched pagination
items, err := client.ListAuthenticationStrengthPoliciesComplete(ctx)
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


### Example Usage: `AuthenticationStrengthPolicyClient.UpdatePolicyAuthenticationStrengthPolicyAllowedCombination`

```go
ctx := context.TODO()
id := authenticationstrengthpolicy.NewPolicyAuthenticationStrengthPolicyID("authenticationStrengthPolicyIdValue")

payload := authenticationstrengthpolicy.UpdatePolicyAuthenticationStrengthPolicyAllowedCombinationRequest{
	// ...
}


read, err := client.UpdatePolicyAuthenticationStrengthPolicyAllowedCombination(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
