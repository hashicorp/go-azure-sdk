
## `github.com/hashicorp/go-azure-sdk/resource-manager/policies/beta/policyauthenticationstrengthpolicy` Documentation

The `policyauthenticationstrengthpolicy` SDK allows for interaction with the Azure Resource Manager Service `policies` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/policies/beta/policyauthenticationstrengthpolicy"
```


### Client Initialization

```go
client := policyauthenticationstrengthpolicy.NewPolicyAuthenticationStrengthPolicyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PolicyAuthenticationStrengthPolicyClient.CreatePolicyAuthenticationStrengthPolicy`

```go
ctx := context.TODO()

payload := policyauthenticationstrengthpolicy.AuthenticationStrengthPolicy{
	// ...
}


read, err := client.CreatePolicyAuthenticationStrengthPolicy(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyAuthenticationStrengthPolicyClient.DeletePolicyAuthenticationStrengthPolicyById`

```go
ctx := context.TODO()
id := policyauthenticationstrengthpolicy.NewPolicyAuthenticationStrengthPolicyID("authenticationStrengthPolicyIdValue")

read, err := client.DeletePolicyAuthenticationStrengthPolicyById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyAuthenticationStrengthPolicyClient.GetPolicyAuthenticationStrengthPolicyById`

```go
ctx := context.TODO()
id := policyauthenticationstrengthpolicy.NewPolicyAuthenticationStrengthPolicyID("authenticationStrengthPolicyIdValue")

read, err := client.GetPolicyAuthenticationStrengthPolicyById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyAuthenticationStrengthPolicyClient.GetPolicyAuthenticationStrengthPolicyCount`

```go
ctx := context.TODO()


read, err := client.GetPolicyAuthenticationStrengthPolicyCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyAuthenticationStrengthPolicyClient.ListPolicyAuthenticationStrengthPolicies`

```go
ctx := context.TODO()


// alternatively `client.ListPolicyAuthenticationStrengthPolicies(ctx)` can be used to do batched pagination
items, err := client.ListPolicyAuthenticationStrengthPoliciesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicyAuthenticationStrengthPolicyClient.UpdatePolicyAuthenticationStrengthPolicyById`

```go
ctx := context.TODO()
id := policyauthenticationstrengthpolicy.NewPolicyAuthenticationStrengthPolicyID("authenticationStrengthPolicyIdValue")

payload := policyauthenticationstrengthpolicy.AuthenticationStrengthPolicy{
	// ...
}


read, err := client.UpdatePolicyAuthenticationStrengthPolicyById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyAuthenticationStrengthPolicyClient.UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombination`

```go
ctx := context.TODO()
id := policyauthenticationstrengthpolicy.NewPolicyAuthenticationStrengthPolicyID("authenticationStrengthPolicyIdValue")

payload := policyauthenticationstrengthpolicy.UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombinationRequest{
	// ...
}


read, err := client.UpdatePolicyAuthenticationStrengthPolicyByIdAllowedCombination(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
