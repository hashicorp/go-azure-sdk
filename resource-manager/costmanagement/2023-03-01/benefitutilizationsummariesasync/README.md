
## `github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2023-03-01/benefitutilizationsummariesasync` Documentation

The `benefitutilizationsummariesasync` SDK allows for interaction with the Azure Resource Manager Service `costmanagement` (API Version `2023-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2023-03-01/benefitutilizationsummariesasync"
```


### Client Initialization

```go
client := benefitutilizationsummariesasync.NewBenefitUtilizationSummariesAsyncClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BenefitUtilizationSummariesAsyncClient.BillingAccountScopeGenerateBenefitUtilizationSummariesReport`

```go
ctx := context.TODO()
id := benefitutilizationsummariesasync.NewBillingAccountID("billingAccountIdValue")

payload := benefitutilizationsummariesasync.BenefitUtilizationSummariesRequest{
	// ...
}


if err := client.BillingAccountScopeGenerateBenefitUtilizationSummariesReportThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BenefitUtilizationSummariesAsyncClient.BillingProfileScopeGenerateBenefitUtilizationSummariesReport`

```go
ctx := context.TODO()
id := benefitutilizationsummariesasync.NewBillingProfileID("billingAccountIdValue", "billingProfileIdValue")

payload := benefitutilizationsummariesasync.BenefitUtilizationSummariesRequest{
	// ...
}


if err := client.BillingProfileScopeGenerateBenefitUtilizationSummariesReportThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BenefitUtilizationSummariesAsyncClient.ReservationOrderScopeGenerateBenefitUtilizationSummariesReport`

```go
ctx := context.TODO()
id := benefitutilizationsummariesasync.NewReservationOrderID("reservationOrderIdValue")

payload := benefitutilizationsummariesasync.BenefitUtilizationSummariesRequest{
	// ...
}


if err := client.ReservationOrderScopeGenerateBenefitUtilizationSummariesReportThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BenefitUtilizationSummariesAsyncClient.ReservationScopeGenerateBenefitUtilizationSummariesReport`

```go
ctx := context.TODO()
id := benefitutilizationsummariesasync.NewReservationID("reservationOrderIdValue", "reservationIdValue")

payload := benefitutilizationsummariesasync.BenefitUtilizationSummariesRequest{
	// ...
}


if err := client.ReservationScopeGenerateBenefitUtilizationSummariesReportThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BenefitUtilizationSummariesAsyncClient.SavingsPlanOrderScopeGenerateBenefitUtilizationSummariesReport`

```go
ctx := context.TODO()
id := benefitutilizationsummariesasync.NewSavingsPlanOrderID("savingsPlanOrderIdValue")

payload := benefitutilizationsummariesasync.BenefitUtilizationSummariesRequest{
	// ...
}


if err := client.SavingsPlanOrderScopeGenerateBenefitUtilizationSummariesReportThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BenefitUtilizationSummariesAsyncClient.SavingsPlanScopeGenerateBenefitUtilizationSummariesReportAsync`

```go
ctx := context.TODO()
id := benefitutilizationsummariesasync.NewSavingsPlanID("savingsPlanOrderIdValue", "savingsPlanIdValue")

payload := benefitutilizationsummariesasync.BenefitUtilizationSummariesRequest{
	// ...
}


if err := client.SavingsPlanScopeGenerateBenefitUtilizationSummariesReportAsyncThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
