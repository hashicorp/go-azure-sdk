
## `github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2024-08-01/reservationsummaries` Documentation

The `reservationsummaries` SDK allows for interaction with the Azure Resource Manager Service `consumption` (API Version `2024-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2024-08-01/reservationsummaries"
```


### Client Initialization

```go
client := reservationsummaries.NewReservationSummariesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReservationSummariesClient.ReservationsSummariesList`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.ReservationsSummariesList(ctx, id, reservationsummaries.DefaultReservationsSummariesListOperationOptions())` can be used to do batched pagination
items, err := client.ReservationsSummariesListComplete(ctx, id, reservationsummaries.DefaultReservationsSummariesListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ReservationSummariesClient.ReservationsSummariesListByReservationOrder`

```go
ctx := context.TODO()
id := reservationsummaries.NewReservationOrderID("reservationOrderIdValue")

// alternatively `client.ReservationsSummariesListByReservationOrder(ctx, id, reservationsummaries.DefaultReservationsSummariesListByReservationOrderOperationOptions())` can be used to do batched pagination
items, err := client.ReservationsSummariesListByReservationOrderComplete(ctx, id, reservationsummaries.DefaultReservationsSummariesListByReservationOrderOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ReservationSummariesClient.ReservationsSummariesListByReservationOrderAndReservation`

```go
ctx := context.TODO()
id := reservationsummaries.NewReservationID("reservationOrderIdValue", "reservationIdValue")

// alternatively `client.ReservationsSummariesListByReservationOrderAndReservation(ctx, id, reservationsummaries.DefaultReservationsSummariesListByReservationOrderAndReservationOperationOptions())` can be used to do batched pagination
items, err := client.ReservationsSummariesListByReservationOrderAndReservationComplete(ctx, id, reservationsummaries.DefaultReservationsSummariesListByReservationOrderAndReservationOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
