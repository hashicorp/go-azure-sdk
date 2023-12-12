
## `github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2023-11-01/benefitutilizationsummariesasync` Documentation

The `benefitutilizationsummariesasync` SDK allows for interaction with the Azure Resource Manager Service `costmanagement` (API Version `2023-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2023-11-01/benefitutilizationsummariesasync"
```


### Client Initialization

```go
client := benefitutilizationsummariesasync.NewBenefitUtilizationSummariesAsyncClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BenefitUtilizationSummariesAsyncClient.GenerateBenefitUtilizationSummariesReportGenerateByBillingAccount`

```go
ctx := context.TODO()
id := benefitutilizationsummariesasync.NewBillingAccountID("billingAccountIdValue")

payload := benefitutilizationsummariesasync.BenefitUtilizationSummariesRequest{
	// ...
}


if err := client.GenerateBenefitUtilizationSummariesReportGenerateByBillingAccountThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BenefitUtilizationSummariesAsyncClient.GenerateBenefitUtilizationSummariesReportGenerateByBillingProfile`

```go
ctx := context.TODO()
id := benefitutilizationsummariesasync.NewBillingProfileID("billingAccountIdValue", "billingProfileIdValue")

payload := benefitutilizationsummariesasync.BenefitUtilizationSummariesRequest{
	// ...
}


if err := client.GenerateBenefitUtilizationSummariesReportGenerateByBillingProfileThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BenefitUtilizationSummariesAsyncClient.GenerateBenefitUtilizationSummariesReportGenerateByReservationId`

```go
ctx := context.TODO()
id := benefitutilizationsummariesasync.NewReservationID("reservationOrderIdValue", "reservationIdValue")

payload := benefitutilizationsummariesasync.BenefitUtilizationSummariesRequest{
	// ...
}


if err := client.GenerateBenefitUtilizationSummariesReportGenerateByReservationIdThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BenefitUtilizationSummariesAsyncClient.GenerateBenefitUtilizationSummariesReportGenerateByReservationOrderId`

```go
ctx := context.TODO()
id := benefitutilizationsummariesasync.NewReservationOrderID("reservationOrderIdValue")

payload := benefitutilizationsummariesasync.BenefitUtilizationSummariesRequest{
	// ...
}


if err := client.GenerateBenefitUtilizationSummariesReportGenerateByReservationOrderIdThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BenefitUtilizationSummariesAsyncClient.GenerateBenefitUtilizationSummariesReportGenerateBySavingsPlanId`

```go
ctx := context.TODO()
id := benefitutilizationsummariesasync.NewSavingsPlanID("savingsPlanOrderIdValue", "savingsPlanIdValue")

payload := benefitutilizationsummariesasync.BenefitUtilizationSummariesRequest{
	// ...
}


if err := client.GenerateBenefitUtilizationSummariesReportGenerateBySavingsPlanIdThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BenefitUtilizationSummariesAsyncClient.GenerateBenefitUtilizationSummariesReportGenerateBySavingsPlanOrderId`

```go
ctx := context.TODO()
id := benefitutilizationsummariesasync.NewSavingsPlanOrderID("savingsPlanOrderIdValue")

payload := benefitutilizationsummariesasync.BenefitUtilizationSummariesRequest{
	// ...
}


if err := client.GenerateBenefitUtilizationSummariesReportGenerateBySavingsPlanOrderIdThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
