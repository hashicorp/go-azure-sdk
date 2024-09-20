
## `github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2023-08-01/reservedinstances` Documentation

The `reservedinstances` SDK allows for interaction with Azure Resource Manager `costmanagement` (API Version `2023-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2023-08-01/reservedinstances"
```


### Client Initialization

```go
client := reservedinstances.NewReservedInstancesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReservedInstancesClient.GenerateReservationDetailsReportByBillingAccountId`

```go
ctx := context.TODO()
id := reservedinstances.NewBillingAccountID("billingAccountId")

if err := client.GenerateReservationDetailsReportByBillingAccountIdThenPoll(ctx, id, reservedinstances.DefaultGenerateReservationDetailsReportByBillingAccountIdOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `ReservedInstancesClient.GenerateReservationDetailsReportByBillingProfileId`

```go
ctx := context.TODO()
id := reservedinstances.NewBillingProfileID("billingAccountId", "billingProfileId")

if err := client.GenerateReservationDetailsReportByBillingProfileIdThenPoll(ctx, id, reservedinstances.DefaultGenerateReservationDetailsReportByBillingProfileIdOperationOptions()); err != nil {
	// handle the error
}
```
