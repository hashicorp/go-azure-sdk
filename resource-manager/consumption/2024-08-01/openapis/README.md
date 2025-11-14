
## `github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2024-08-01/openapis` Documentation

The `openapis` SDK allows for interaction with Azure Resource Manager `consumption` (API Version `2024-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2024-08-01/openapis"
```


### Client Initialization

```go
client := openapis.NewOpenapisClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OpenapisClient.AggregatedCostGetByManagementGroup`

```go
ctx := context.TODO()
id := commonids.NewManagementGroupID("groupId")

read, err := client.AggregatedCostGetByManagementGroup(ctx, id, openapis.DefaultAggregatedCostGetByManagementGroupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.AggregatedCostGetForBillingPeriodByManagementGroup`

```go
ctx := context.TODO()
id := openapis.NewProviders2BillingPeriodID("managementGroupId", "billingPeriodName")

read, err := client.AggregatedCostGetForBillingPeriodByManagementGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.BalancesGetByBillingAccount`

```go
ctx := context.TODO()
id := openapis.NewBillingAccountID("billingAccountId")

read, err := client.BalancesGetByBillingAccount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.BalancesGetForBillingPeriodByBillingAccount`

```go
ctx := context.TODO()
id := openapis.NewBillingAccountBillingPeriodID("billingAccountId", "billingPeriodName")

read, err := client.BalancesGetForBillingPeriodByBillingAccount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.ChargesList`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

read, err := client.ChargesList(ctx, id, openapis.DefaultChargesListOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.EventsListByBillingAccount`

```go
ctx := context.TODO()
id := openapis.NewBillingAccountID("billingAccountId")

// alternatively `client.EventsListByBillingAccount(ctx, id, openapis.DefaultEventsListByBillingAccountOperationOptions())` can be used to do batched pagination
items, err := client.EventsListByBillingAccountComplete(ctx, id, openapis.DefaultEventsListByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.EventsListByBillingProfile`

```go
ctx := context.TODO()
id := openapis.NewBillingProfileID("billingAccountId", "billingProfileId")

// alternatively `client.EventsListByBillingProfile(ctx, id, openapis.DefaultEventsListByBillingProfileOperationOptions())` can be used to do batched pagination
items, err := client.EventsListByBillingProfileComplete(ctx, id, openapis.DefaultEventsListByBillingProfileOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.LotsListByBillingAccount`

```go
ctx := context.TODO()
id := openapis.NewBillingAccountID("billingAccountId")

// alternatively `client.LotsListByBillingAccount(ctx, id, openapis.DefaultLotsListByBillingAccountOperationOptions())` can be used to do batched pagination
items, err := client.LotsListByBillingAccountComplete(ctx, id, openapis.DefaultLotsListByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.LotsListByBillingProfile`

```go
ctx := context.TODO()
id := openapis.NewBillingProfileID("billingAccountId", "billingProfileId")

// alternatively `client.LotsListByBillingProfile(ctx, id)` can be used to do batched pagination
items, err := client.LotsListByBillingProfileComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.LotsListByCustomer`

```go
ctx := context.TODO()
id := openapis.NewCustomerID("billingAccountId", "customerId")

// alternatively `client.LotsListByCustomer(ctx, id, openapis.DefaultLotsListByCustomerOperationOptions())` can be used to do batched pagination
items, err := client.LotsListByCustomerComplete(ctx, id, openapis.DefaultLotsListByCustomerOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.MarketplacesList`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.MarketplacesList(ctx, id, openapis.DefaultMarketplacesListOperationOptions())` can be used to do batched pagination
items, err := client.MarketplacesListComplete(ctx, id, openapis.DefaultMarketplacesListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.PriceSheetDownloadByBillingAccountPeriod`

```go
ctx := context.TODO()
id := openapis.NewBillingAccountBillingPeriodID("billingAccountId", "billingPeriodName")

if err := client.PriceSheetDownloadByBillingAccountPeriodThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `OpenapisClient.ReservationRecommendationDetailsGet`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

read, err := client.ReservationRecommendationDetailsGet(ctx, id, openapis.DefaultReservationRecommendationDetailsGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.ReservationRecommendationsList`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.ReservationRecommendationsList(ctx, id, openapis.DefaultReservationRecommendationsListOperationOptions())` can be used to do batched pagination
items, err := client.ReservationRecommendationsListComplete(ctx, id, openapis.DefaultReservationRecommendationsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.ReservationTransactionsList`

```go
ctx := context.TODO()
id := openapis.NewBillingAccountID("billingAccountId")

// alternatively `client.ReservationTransactionsList(ctx, id, openapis.DefaultReservationTransactionsListOperationOptions())` can be used to do batched pagination
items, err := client.ReservationTransactionsListComplete(ctx, id, openapis.DefaultReservationTransactionsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.ReservationTransactionsListByBillingProfile`

```go
ctx := context.TODO()
id := openapis.NewBillingProfileID("billingAccountId", "billingProfileId")

// alternatively `client.ReservationTransactionsListByBillingProfile(ctx, id, openapis.DefaultReservationTransactionsListByBillingProfileOperationOptions())` can be used to do batched pagination
items, err := client.ReservationTransactionsListByBillingProfileComplete(ctx, id, openapis.DefaultReservationTransactionsListByBillingProfileOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.ReservationsDetailsList`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.ReservationsDetailsList(ctx, id, openapis.DefaultReservationsDetailsListOperationOptions())` can be used to do batched pagination
items, err := client.ReservationsDetailsListComplete(ctx, id, openapis.DefaultReservationsDetailsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.ReservationsDetailsListByReservationOrder`

```go
ctx := context.TODO()
id := openapis.NewReservationOrderID("reservationOrderId")

// alternatively `client.ReservationsDetailsListByReservationOrder(ctx, id, openapis.DefaultReservationsDetailsListByReservationOrderOperationOptions())` can be used to do batched pagination
items, err := client.ReservationsDetailsListByReservationOrderComplete(ctx, id, openapis.DefaultReservationsDetailsListByReservationOrderOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.ReservationsDetailsListByReservationOrderAndReservation`

```go
ctx := context.TODO()
id := openapis.NewReservationID("reservationOrderId", "reservationId")

// alternatively `client.ReservationsDetailsListByReservationOrderAndReservation(ctx, id, openapis.DefaultReservationsDetailsListByReservationOrderAndReservationOperationOptions())` can be used to do batched pagination
items, err := client.ReservationsDetailsListByReservationOrderAndReservationComplete(ctx, id, openapis.DefaultReservationsDetailsListByReservationOrderAndReservationOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.ReservationsSummariesList`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.ReservationsSummariesList(ctx, id, openapis.DefaultReservationsSummariesListOperationOptions())` can be used to do batched pagination
items, err := client.ReservationsSummariesListComplete(ctx, id, openapis.DefaultReservationsSummariesListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.ReservationsSummariesListByReservationOrder`

```go
ctx := context.TODO()
id := openapis.NewReservationOrderID("reservationOrderId")

// alternatively `client.ReservationsSummariesListByReservationOrder(ctx, id, openapis.DefaultReservationsSummariesListByReservationOrderOperationOptions())` can be used to do batched pagination
items, err := client.ReservationsSummariesListByReservationOrderComplete(ctx, id, openapis.DefaultReservationsSummariesListByReservationOrderOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.ReservationsSummariesListByReservationOrderAndReservation`

```go
ctx := context.TODO()
id := openapis.NewReservationID("reservationOrderId", "reservationId")

// alternatively `client.ReservationsSummariesListByReservationOrderAndReservation(ctx, id, openapis.DefaultReservationsSummariesListByReservationOrderAndReservationOperationOptions())` can be used to do batched pagination
items, err := client.ReservationsSummariesListByReservationOrderAndReservationComplete(ctx, id, openapis.DefaultReservationsSummariesListByReservationOrderAndReservationOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
