
## `github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2022-10-01/reservedinstances` Documentation

The `reservedinstances` SDK allows for interaction with the Azure Resource Manager Service `costmanagement` (API Version `2022-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2022-10-01/reservedinstances"
```


### Client Initialization

```go
client := reservedinstances.NewReservedInstancesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReservedInstancesClient.GenerateReservationDetailsReportByBillingAccountId`

```go
ctx := context.TODO()
id := reservedinstances.NewBillingAccountID("billingAccountIdValue")

if err := client.GenerateReservationDetailsReportByBillingAccountIdThenPoll(ctx, id, reservedinstances.DefaultGenerateReservationDetailsReportByBillingAccountIdOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `ReservedInstancesClient.GenerateReservationDetailsReportByBillingProfileId`

```go
ctx := context.TODO()
id := reservedinstances.NewBillingProfileID("billingAccountIdValue", "billingProfileIdValue")

if err := client.GenerateReservationDetailsReportByBillingProfileIdThenPoll(ctx, id, reservedinstances.DefaultGenerateReservationDetailsReportByBillingProfileIdOperationOptions()); err != nil {
	// handle the error
}
```
