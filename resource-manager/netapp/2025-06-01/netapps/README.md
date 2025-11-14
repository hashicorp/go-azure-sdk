
## `github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2025-06-01/netapps` Documentation

The `netapps` SDK allows for interaction with Azure Resource Manager `netapp` (API Version `2025-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2025-06-01/netapps"
```


### Client Initialization

```go
client := netapps.NewNetappsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `NetappsClient.NetAppResourceCheckFilePathAvailability`

```go
ctx := context.TODO()
id := netapps.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

payload := netapps.FilePathAvailabilityRequest{
	// ...
}


read, err := client.NetAppResourceCheckFilePathAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetappsClient.NetAppResourceCheckNameAvailability`

```go
ctx := context.TODO()
id := netapps.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

payload := netapps.ResourceNameAvailabilityRequest{
	// ...
}


read, err := client.NetAppResourceCheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetappsClient.NetAppResourceCheckQuotaAvailability`

```go
ctx := context.TODO()
id := netapps.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

payload := netapps.QuotaAvailabilityRequest{
	// ...
}


read, err := client.NetAppResourceCheckQuotaAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetappsClient.NetAppResourceQueryNetworkSiblingSet`

```go
ctx := context.TODO()
id := netapps.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

payload := netapps.QueryNetworkSiblingSetRequest{
	// ...
}


read, err := client.NetAppResourceQueryNetworkSiblingSet(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetappsClient.NetAppResourceQueryRegionInfo`

```go
ctx := context.TODO()
id := netapps.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

read, err := client.NetAppResourceQueryRegionInfo(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetappsClient.NetAppResourceUpdateNetworkSiblingSet`

```go
ctx := context.TODO()
id := netapps.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

payload := netapps.UpdateNetworkSiblingSetRequest{
	// ...
}


if err := client.NetAppResourceUpdateNetworkSiblingSetThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `NetappsClient.NetAppResourceUsagesGet`

```go
ctx := context.TODO()
id := netapps.NewUsageID("12345678-1234-9876-4563-123456789012", "locationName", "usageName")

read, err := client.NetAppResourceUsagesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetappsClient.NetAppResourceUsagesList`

```go
ctx := context.TODO()
id := netapps.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

// alternatively `client.NetAppResourceUsagesList(ctx, id)` can be used to do batched pagination
items, err := client.NetAppResourceUsagesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
