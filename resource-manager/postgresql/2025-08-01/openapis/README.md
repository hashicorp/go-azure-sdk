
## `github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/openapis` Documentation

The `openapis` SDK allows for interaction with Azure Resource Manager `postgresql` (API Version `2025-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/openapis"
```


### Client Initialization

```go
client := openapis.NewOpenapisClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OpenapisClient.CapabilitiesByLocationList`

```go
ctx := context.TODO()
id := openapis.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

// alternatively `client.CapabilitiesByLocationList(ctx, id)` can be used to do batched pagination
items, err := client.CapabilitiesByLocationListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.NameAvailabilityCheckGlobally`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

payload := openapis.CheckNameAvailabilityRequest{
	// ...
}


read, err := client.NameAvailabilityCheckGlobally(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.NameAvailabilityCheckWithLocation`

```go
ctx := context.TODO()
id := openapis.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

payload := openapis.CheckNameAvailabilityRequest{
	// ...
}


read, err := client.NameAvailabilityCheckWithLocation(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.PrivateDnsZoneSuffixGet`

```go
ctx := context.TODO()


read, err := client.PrivateDnsZoneSuffixGet(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.QuotaUsagesList`

```go
ctx := context.TODO()
id := openapis.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

// alternatively `client.QuotaUsagesList(ctx, id)` can be used to do batched pagination
items, err := client.QuotaUsagesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.VirtualNetworkSubnetUsageList`

```go
ctx := context.TODO()
id := openapis.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

payload := openapis.VirtualNetworkSubnetUsageParameter{
	// ...
}


read, err := client.VirtualNetworkSubnetUsageList(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
