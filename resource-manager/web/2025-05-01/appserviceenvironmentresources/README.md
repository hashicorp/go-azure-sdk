
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/appserviceenvironmentresources` Documentation

The `appserviceenvironmentresources` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/appserviceenvironmentresources"
```


### Client Initialization

```go
client := appserviceenvironmentresources.NewAppServiceEnvironmentResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AppServiceEnvironmentResourcesClient.AppServiceEnvironmentsChangeVnet`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

payload := appserviceenvironmentresources.VirtualNetworkProfile{
	// ...
}


// alternatively `client.AppServiceEnvironmentsChangeVnet(ctx, id, payload)` can be used to do batched pagination
items, err := client.AppServiceEnvironmentsChangeVnetComplete(ctx, id, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AppServiceEnvironmentResourcesClient.AppServiceEnvironmentsCreateOrUpdate`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

payload := appserviceenvironmentresources.AppServiceEnvironmentResource{
	// ...
}


if err := client.AppServiceEnvironmentsCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AppServiceEnvironmentResourcesClient.AppServiceEnvironmentsDelete`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

if err := client.AppServiceEnvironmentsDeleteThenPoll(ctx, id, appserviceenvironmentresources.DefaultAppServiceEnvironmentsDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `AppServiceEnvironmentResourcesClient.AppServiceEnvironmentsGet`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

read, err := client.AppServiceEnvironmentsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AppServiceEnvironmentResourcesClient.AppServiceEnvironmentsGetDiagnosticsItem`

```go
ctx := context.TODO()
id := appserviceenvironmentresources.NewHostingEnvironmentDiagnosticID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName", "diagnosticName")

read, err := client.AppServiceEnvironmentsGetDiagnosticsItem(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AppServiceEnvironmentResourcesClient.AppServiceEnvironmentsGetInboundNetworkDependenciesEndpoints`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

// alternatively `client.AppServiceEnvironmentsGetInboundNetworkDependenciesEndpoints(ctx, id)` can be used to do batched pagination
items, err := client.AppServiceEnvironmentsGetInboundNetworkDependenciesEndpointsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AppServiceEnvironmentResourcesClient.AppServiceEnvironmentsGetOutboundNetworkDependenciesEndpoints`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

// alternatively `client.AppServiceEnvironmentsGetOutboundNetworkDependenciesEndpoints(ctx, id)` can be used to do batched pagination
items, err := client.AppServiceEnvironmentsGetOutboundNetworkDependenciesEndpointsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AppServiceEnvironmentResourcesClient.AppServiceEnvironmentsGetPrivateLinkResources`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

read, err := client.AppServiceEnvironmentsGetPrivateLinkResources(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AppServiceEnvironmentResourcesClient.AppServiceEnvironmentsList`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.AppServiceEnvironmentsList(ctx, id)` can be used to do batched pagination
items, err := client.AppServiceEnvironmentsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AppServiceEnvironmentResourcesClient.AppServiceEnvironmentsListAppServicePlans`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

// alternatively `client.AppServiceEnvironmentsListAppServicePlans(ctx, id)` can be used to do batched pagination
items, err := client.AppServiceEnvironmentsListAppServicePlansComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AppServiceEnvironmentResourcesClient.AppServiceEnvironmentsListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.AppServiceEnvironmentsListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.AppServiceEnvironmentsListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AppServiceEnvironmentResourcesClient.AppServiceEnvironmentsListCapacities`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

// alternatively `client.AppServiceEnvironmentsListCapacities(ctx, id)` can be used to do batched pagination
items, err := client.AppServiceEnvironmentsListCapacitiesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AppServiceEnvironmentResourcesClient.AppServiceEnvironmentsListDiagnostics`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

read, err := client.AppServiceEnvironmentsListDiagnostics(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AppServiceEnvironmentResourcesClient.AppServiceEnvironmentsListOperations`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

read, err := client.AppServiceEnvironmentsListOperations(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AppServiceEnvironmentResourcesClient.AppServiceEnvironmentsListUsages`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

// alternatively `client.AppServiceEnvironmentsListUsages(ctx, id, appserviceenvironmentresources.DefaultAppServiceEnvironmentsListUsagesOperationOptions())` can be used to do batched pagination
items, err := client.AppServiceEnvironmentsListUsagesComplete(ctx, id, appserviceenvironmentresources.DefaultAppServiceEnvironmentsListUsagesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AppServiceEnvironmentResourcesClient.AppServiceEnvironmentsListWebApps`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

// alternatively `client.AppServiceEnvironmentsListWebApps(ctx, id, appserviceenvironmentresources.DefaultAppServiceEnvironmentsListWebAppsOperationOptions())` can be used to do batched pagination
items, err := client.AppServiceEnvironmentsListWebAppsComplete(ctx, id, appserviceenvironmentresources.DefaultAppServiceEnvironmentsListWebAppsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AppServiceEnvironmentResourcesClient.AppServiceEnvironmentsReboot`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

read, err := client.AppServiceEnvironmentsReboot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AppServiceEnvironmentResourcesClient.AppServiceEnvironmentsResume`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

// alternatively `client.AppServiceEnvironmentsResume(ctx, id)` can be used to do batched pagination
items, err := client.AppServiceEnvironmentsResumeComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AppServiceEnvironmentResourcesClient.AppServiceEnvironmentsSuspend`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

// alternatively `client.AppServiceEnvironmentsSuspend(ctx, id)` can be used to do batched pagination
items, err := client.AppServiceEnvironmentsSuspendComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AppServiceEnvironmentResourcesClient.AppServiceEnvironmentsTestUpgradeAvailableNotification`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

read, err := client.AppServiceEnvironmentsTestUpgradeAvailableNotification(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AppServiceEnvironmentResourcesClient.AppServiceEnvironmentsUpdate`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

payload := appserviceenvironmentresources.AppServiceEnvironmentPatchResource{
	// ...
}


read, err := client.AppServiceEnvironmentsUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AppServiceEnvironmentResourcesClient.AppServiceEnvironmentsUpgrade`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

if err := client.AppServiceEnvironmentsUpgradeThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `AppServiceEnvironmentResourcesClient.RecommendationsDisableAllForHostingEnvironment`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

read, err := client.RecommendationsDisableAllForHostingEnvironment(ctx, id, appserviceenvironmentresources.DefaultRecommendationsDisableAllForHostingEnvironmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AppServiceEnvironmentResourcesClient.RecommendationsListHistoryForHostingEnvironment`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

// alternatively `client.RecommendationsListHistoryForHostingEnvironment(ctx, id, appserviceenvironmentresources.DefaultRecommendationsListHistoryForHostingEnvironmentOperationOptions())` can be used to do batched pagination
items, err := client.RecommendationsListHistoryForHostingEnvironmentComplete(ctx, id, appserviceenvironmentresources.DefaultRecommendationsListHistoryForHostingEnvironmentOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AppServiceEnvironmentResourcesClient.RecommendationsListRecommendedRulesForHostingEnvironment`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

// alternatively `client.RecommendationsListRecommendedRulesForHostingEnvironment(ctx, id, appserviceenvironmentresources.DefaultRecommendationsListRecommendedRulesForHostingEnvironmentOperationOptions())` can be used to do batched pagination
items, err := client.RecommendationsListRecommendedRulesForHostingEnvironmentComplete(ctx, id, appserviceenvironmentresources.DefaultRecommendationsListRecommendedRulesForHostingEnvironmentOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AppServiceEnvironmentResourcesClient.RecommendationsResetAllFiltersForHostingEnvironment`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

read, err := client.RecommendationsResetAllFiltersForHostingEnvironment(ctx, id, appserviceenvironmentresources.DefaultRecommendationsResetAllFiltersForHostingEnvironmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
