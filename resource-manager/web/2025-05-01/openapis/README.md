
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/openapis` Documentation

The `openapis` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/openapis"
```


### Client Initialization

```go
client := openapis.NewOpenapisClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OpenapisClient.CheckNameAvailability`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

payload := openapis.ResourceNameAvailabilityRequest{
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


### Example Usage: `OpenapisClient.GetPublishingUser`

```go
ctx := context.TODO()


read, err := client.GetPublishingUser(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.GetSourceControl`

```go
ctx := context.TODO()
id := openapis.NewSourceControlID("sourceControlName")

read, err := client.GetSourceControl(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.GetSubscriptionDeploymentLocations`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.GetSubscriptionDeploymentLocations(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.GetUsagesInLocationlist`

```go
ctx := context.TODO()
id := openapis.NewProviderLocationID("12345678-1234-9876-4563-123456789012", "locationName")

// alternatively `client.GetUsagesInLocationlist(ctx, id)` can be used to do batched pagination
items, err := client.GetUsagesInLocationlistComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.ListAseRegions`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListAseRegions(ctx, id)` can be used to do batched pagination
items, err := client.ListAseRegionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.ListBillingMeters`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBillingMeters(ctx, id, openapis.DefaultListBillingMetersOperationOptions())` can be used to do batched pagination
items, err := client.ListBillingMetersComplete(ctx, id, openapis.DefaultListBillingMetersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.ListCustomHostNameSites`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListCustomHostNameSites(ctx, id, openapis.DefaultListCustomHostNameSitesOperationOptions())` can be used to do batched pagination
items, err := client.ListCustomHostNameSitesComplete(ctx, id, openapis.DefaultListCustomHostNameSitesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.ListGeoRegions`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListGeoRegions(ctx, id, openapis.DefaultListGeoRegionsOperationOptions())` can be used to do batched pagination
items, err := client.ListGeoRegionsComplete(ctx, id, openapis.DefaultListGeoRegionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.ListPremierAddOnOffers`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListPremierAddOnOffers(ctx, id)` can be used to do batched pagination
items, err := client.ListPremierAddOnOffersComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.ListSiteIdentifiersAssignedToHostName`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

payload := openapis.NameIdentifier{
	// ...
}


// alternatively `client.ListSiteIdentifiersAssignedToHostName(ctx, id, payload)` can be used to do batched pagination
items, err := client.ListSiteIdentifiersAssignedToHostNameComplete(ctx, id, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.ListSkus`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.ListSkus(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.ListSourceControls`

```go
ctx := context.TODO()


// alternatively `client.ListSourceControls(ctx)` can be used to do batched pagination
items, err := client.ListSourceControlsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.Move`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

payload := openapis.CsmMoveResourceEnvelope{
	// ...
}


read, err := client.Move(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.RegionalCheckNameAvailability`

```go
ctx := context.TODO()
id := openapis.NewProviderLocationID("12345678-1234-9876-4563-123456789012", "locationName")

payload := openapis.DnlResourceNameAvailabilityRequest{
	// ...
}


read, err := client.RegionalCheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.UpdatePublishingUser`

```go
ctx := context.TODO()

payload := openapis.User{
	// ...
}


read, err := client.UpdatePublishingUser(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.UpdateSourceControl`

```go
ctx := context.TODO()
id := openapis.NewSourceControlID("sourceControlName")

payload := openapis.SourceControl{
	// ...
}


read, err := client.UpdateSourceControl(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.Validate`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

payload := openapis.ValidateRequest{
	// ...
}


read, err := client.Validate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.ValidateMove`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

payload := openapis.CsmMoveResourceEnvelope{
	// ...
}


read, err := client.ValidateMove(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.VerifyHostingEnvironmentVnet`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

payload := openapis.VnetParameters{
	// ...
}


read, err := client.VerifyHostingEnvironmentVnet(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
