
## `github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2023-05-01/reservationdetails` Documentation

The `reservationdetails` SDK allows for interaction with Azure Resource Manager `consumption` (API Version `2023-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2023-05-01/reservationdetails"
```


### Client Initialization

```go
client := reservationdetails.NewReservationDetailsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReservationDetailsClient.ReservationsDetailsList`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.ReservationsDetailsList(ctx, id, reservationdetails.DefaultReservationsDetailsListOperationOptions())` can be used to do batched pagination
items, err := client.ReservationsDetailsListComplete(ctx, id, reservationdetails.DefaultReservationsDetailsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ReservationDetailsClient.ReservationsDetailsListByReservationOrder`

```go
ctx := context.TODO()
id := reservationdetails.NewReservationOrderID("reservationOrderId")

// alternatively `client.ReservationsDetailsListByReservationOrder(ctx, id, reservationdetails.DefaultReservationsDetailsListByReservationOrderOperationOptions())` can be used to do batched pagination
items, err := client.ReservationsDetailsListByReservationOrderComplete(ctx, id, reservationdetails.DefaultReservationsDetailsListByReservationOrderOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ReservationDetailsClient.ReservationsDetailsListByReservationOrderAndReservation`

```go
ctx := context.TODO()
id := reservationdetails.NewReservationID("reservationOrderId", "reservationId")

// alternatively `client.ReservationsDetailsListByReservationOrderAndReservation(ctx, id, reservationdetails.DefaultReservationsDetailsListByReservationOrderAndReservationOperationOptions())` can be used to do batched pagination
items, err := client.ReservationsDetailsListByReservationOrderAndReservationComplete(ctx, id, reservationdetails.DefaultReservationsDetailsListByReservationOrderAndReservationOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
