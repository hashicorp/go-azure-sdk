
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/agreement` Documentation

The `agreement` SDK allows for interaction with the Azure Resource Manager Service `billing` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/agreement"
```


### Client Initialization

```go
client := agreement.NewAgreementClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AgreementClient.Get`

```go
ctx := context.TODO()
id := agreement.NewAgreementID("billingAccountValue", "agreementValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AgreementClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := agreement.NewBillingAccountID("billingAccountValue")

// alternatively `client.ListByBillingAccount(ctx, id, agreement.DefaultListByBillingAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id, agreement.DefaultListByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
