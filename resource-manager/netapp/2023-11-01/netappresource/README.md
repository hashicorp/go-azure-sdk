
## `github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2023-11-01/netappresource` Documentation

The `netappresource` SDK allows for interaction with Azure Resource Manager `netapp` (API Version `2023-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2023-11-01/netappresource"
```


### Client Initialization

```go
client := netappresource.NewNetAppResourceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `NetAppResourceClient.CheckFilePathAvailability`

```go
ctx := context.TODO()
id := netappresource.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

payload := netappresource.FilePathAvailabilityRequest{
	// ...
}


read, err := client.CheckFilePathAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetAppResourceClient.CheckNameAvailability`

```go
ctx := context.TODO()
id := netappresource.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

payload := netappresource.ResourceNameAvailabilityRequest{
	// ...
}


read, err := client.CheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetAppResourceClient.CheckQuotaAvailability`

```go
ctx := context.TODO()
id := netappresource.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

payload := netappresource.QuotaAvailabilityRequest{
	// ...
}


read, err := client.CheckQuotaAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetAppResourceClient.QueryNetworkSiblingSet`

```go
ctx := context.TODO()
id := netappresource.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

payload := netappresource.QueryNetworkSiblingSetRequest{
	// ...
}


read, err := client.QueryNetworkSiblingSet(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetAppResourceClient.QueryRegionInfo`

```go
ctx := context.TODO()
id := netappresource.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

read, err := client.QueryRegionInfo(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetAppResourceClient.QuotaLimitsGet`

```go
ctx := context.TODO()
id := netappresource.NewQuotaLimitID("12345678-1234-9876-4563-123456789012", "locationName", "quotaLimitName")

read, err := client.QuotaLimitsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetAppResourceClient.QuotaLimitsList`

```go
ctx := context.TODO()
id := netappresource.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

read, err := client.QuotaLimitsList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetAppResourceClient.RegionInfosGet`

```go
ctx := context.TODO()
id := netappresource.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

read, err := client.RegionInfosGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetAppResourceClient.RegionInfosList`

```go
ctx := context.TODO()
id := netappresource.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

// alternatively `client.RegionInfosList(ctx, id)` can be used to do batched pagination
items, err := client.RegionInfosListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetAppResourceClient.UpdateNetworkSiblingSet`

```go
ctx := context.TODO()
id := netappresource.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

payload := netappresource.UpdateNetworkSiblingSetRequest{
	// ...
}


if err := client.UpdateNetworkSiblingSetThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
