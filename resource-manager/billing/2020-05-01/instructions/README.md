
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/instructions` Documentation

The `instructions` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2020-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/instructions"
```


### Client Initialization

```go
client := instructions.NewInstructionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InstructionsClient.Get`

```go
ctx := context.TODO()
id := instructions.NewInstructionID("billingAccountValue", "billingProfileValue", "instructionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InstructionsClient.ListByBillingProfile`

```go
ctx := context.TODO()
id := instructions.NewBillingProfileID("billingAccountValue", "billingProfileValue")

// alternatively `client.ListByBillingProfile(ctx, id)` can be used to do batched pagination
items, err := client.ListByBillingProfileComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `InstructionsClient.Put`

```go
ctx := context.TODO()
id := instructions.NewInstructionID("billingAccountValue", "billingProfileValue", "instructionValue")

payload := instructions.Instruction{
	// ...
}


read, err := client.Put(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
