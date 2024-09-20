
## `github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2019-10-01/reservationsummaries` Documentation

The `reservationsummaries` SDK allows for interaction with Azure Resource Manager `consumption` (API Version `2019-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2019-10-01/reservationsummaries"
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
id := reservationsummaries.NewReservationOrderID("reservationOrderId")

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
id := reservationsummaries.NewReservationID("reservationOrderId", "reservationId")

// alternatively `client.ReservationsSummariesListByReservationOrderAndReservation(ctx, id, reservationsummaries.DefaultReservationsSummariesListByReservationOrderAndReservationOperationOptions())` can be used to do batched pagination
items, err := client.ReservationsSummariesListByReservationOrderAndReservationComplete(ctx, id, reservationsummaries.DefaultReservationsSummariesListByReservationOrderAndReservationOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
