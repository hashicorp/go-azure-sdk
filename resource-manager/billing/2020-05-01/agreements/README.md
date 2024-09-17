
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/agreements` Documentation

The `agreements` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2020-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/agreements"
```


### Client Initialization

```go
client := agreements.NewAgreementsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AgreementsClient.Get`

```go
ctx := context.TODO()
id := agreements.NewAgreementID("billingAccountValue", "agreementValue")

read, err := client.Get(ctx, id, agreements.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AgreementsClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := agreements.NewBillingAccountID("billingAccountValue")

// alternatively `client.ListByBillingAccount(ctx, id, agreements.DefaultListByBillingAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id, agreements.DefaultListByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
