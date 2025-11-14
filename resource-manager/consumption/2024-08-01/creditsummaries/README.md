
## `github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2024-08-01/creditsummaries` Documentation

The `creditsummaries` SDK allows for interaction with Azure Resource Manager `consumption` (API Version `2024-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2024-08-01/creditsummaries"
```


### Client Initialization

```go
client := creditsummaries.NewCreditSummariesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CreditSummariesClient.CreditsGet`

```go
ctx := context.TODO()
id := creditsummaries.NewBillingProfileID("billingAccountId", "billingProfileId")

read, err := client.CreditsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
