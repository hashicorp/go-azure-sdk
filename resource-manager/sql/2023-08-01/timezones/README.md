
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01/timezones` Documentation

The `timezones` SDK allows for interaction with Azure Resource Manager `sql` (API Version `2023-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01/timezones"
```


### Client Initialization

```go
client := timezones.NewTimeZonesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TimeZonesClient.Get`

```go
ctx := context.TODO()
id := timezones.NewTimeZoneID("12345678-1234-9876-4563-123456789012", "locationName", "timeZoneId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TimeZonesClient.ListByLocation`

```go
ctx := context.TODO()
id := timezones.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

// alternatively `client.ListByLocation(ctx, id)` can be used to do batched pagination
items, err := client.ListByLocationComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
