
## `github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/openapis` Documentation

The `openapis` SDK allows for interaction with Azure Resource Manager `costmanagement` (API Version `2025-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/openapis"
```


### Client Initialization

```go
client := openapis.NewOpenapisClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OpenapisClient.AlertsListExternal`

```go
ctx := context.TODO()
id := openapis.NewExternalCloudProviderTypeID("externalBillingAccounts", "externalCloudProviderId")

// alternatively `client.AlertsListExternal(ctx, id)` can be used to do batched pagination
items, err := client.AlertsListExternalComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.BenefitRecommendationsList`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.BenefitRecommendationsList(ctx, id, openapis.DefaultBenefitRecommendationsListOperationOptions())` can be used to do batched pagination
items, err := client.BenefitRecommendationsListComplete(ctx, id, openapis.DefaultBenefitRecommendationsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.BenefitUtilizationSummariesListByBillingAccountId`

```go
ctx := context.TODO()
id := openapis.NewBillingAccountID("billingAccountId")

// alternatively `client.BenefitUtilizationSummariesListByBillingAccountId(ctx, id, openapis.DefaultBenefitUtilizationSummariesListByBillingAccountIdOperationOptions())` can be used to do batched pagination
items, err := client.BenefitUtilizationSummariesListByBillingAccountIdComplete(ctx, id, openapis.DefaultBenefitUtilizationSummariesListByBillingAccountIdOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.BenefitUtilizationSummariesListByBillingProfileId`

```go
ctx := context.TODO()
id := openapis.NewBillingProfileID("billingAccountName", "billingProfileName")

// alternatively `client.BenefitUtilizationSummariesListByBillingProfileId(ctx, id, openapis.DefaultBenefitUtilizationSummariesListByBillingProfileIdOperationOptions())` can be used to do batched pagination
items, err := client.BenefitUtilizationSummariesListByBillingProfileIdComplete(ctx, id, openapis.DefaultBenefitUtilizationSummariesListByBillingProfileIdOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.BenefitUtilizationSummariesListBySavingsPlanId`

```go
ctx := context.TODO()
id := openapis.NewSavingsPlanID("savingsPlanOrderId", "savingsPlanId")

// alternatively `client.BenefitUtilizationSummariesListBySavingsPlanId(ctx, id, openapis.DefaultBenefitUtilizationSummariesListBySavingsPlanIdOperationOptions())` can be used to do batched pagination
items, err := client.BenefitUtilizationSummariesListBySavingsPlanIdComplete(ctx, id, openapis.DefaultBenefitUtilizationSummariesListBySavingsPlanIdOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.BenefitUtilizationSummariesListBySavingsPlanOrder`

```go
ctx := context.TODO()
id := openapis.NewSavingsPlanOrderID("savingsPlanOrderId")

// alternatively `client.BenefitUtilizationSummariesListBySavingsPlanOrder(ctx, id, openapis.DefaultBenefitUtilizationSummariesListBySavingsPlanOrderOperationOptions())` can be used to do batched pagination
items, err := client.BenefitUtilizationSummariesListBySavingsPlanOrderComplete(ctx, id, openapis.DefaultBenefitUtilizationSummariesListBySavingsPlanOrderOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.CostAllocationRulesCheckNameAvailability`

```go
ctx := context.TODO()
id := openapis.NewBillingAccountID("billingAccountId")

payload := openapis.CostAllocationRuleCheckNameAvailabilityRequest{
	// ...
}


read, err := client.CostAllocationRulesCheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.DimensionsByExternalCloudProviderType`

```go
ctx := context.TODO()
id := openapis.NewExternalCloudProviderTypeID("externalBillingAccounts", "externalCloudProviderId")

// alternatively `client.DimensionsByExternalCloudProviderType(ctx, id, openapis.DefaultDimensionsByExternalCloudProviderTypeOperationOptions())` can be used to do batched pagination
items, err := client.DimensionsByExternalCloudProviderTypeComplete(ctx, id, openapis.DefaultDimensionsByExternalCloudProviderTypeOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.DimensionsList`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.DimensionsList(ctx, id, openapis.DefaultDimensionsListOperationOptions())` can be used to do batched pagination
items, err := client.DimensionsListComplete(ctx, id, openapis.DefaultDimensionsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.ForecastExternalCloudProviderUsage`

```go
ctx := context.TODO()
id := openapis.NewExternalCloudProviderTypeID("externalBillingAccounts", "externalCloudProviderId")

payload := openapis.ForecastDefinition{
	// ...
}


read, err := client.ForecastExternalCloudProviderUsage(ctx, id, payload, openapis.DefaultForecastExternalCloudProviderUsageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.ForecastUsage`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

payload := openapis.ForecastDefinition{
	// ...
}


read, err := client.ForecastUsage(ctx, id, payload, openapis.DefaultForecastUsageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.GenerateBenefitUtilizationSummariesReportGenerateByBillingAccount`

```go
ctx := context.TODO()
id := openapis.NewBillingAccountID("billingAccountId")

payload := openapis.BenefitUtilizationSummariesRequest{
	// ...
}


if err := client.GenerateBenefitUtilizationSummariesReportGenerateByBillingAccountThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `OpenapisClient.GenerateBenefitUtilizationSummariesReportGenerateByBillingProfile`

```go
ctx := context.TODO()
id := openapis.NewBillingProfileID("billingAccountName", "billingProfileName")

payload := openapis.BenefitUtilizationSummariesRequest{
	// ...
}


if err := client.GenerateBenefitUtilizationSummariesReportGenerateByBillingProfileThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `OpenapisClient.GenerateBenefitUtilizationSummariesReportGenerateByReservationId`

```go
ctx := context.TODO()
id := openapis.NewReservationID("reservationOrderId", "reservationId")

payload := openapis.BenefitUtilizationSummariesRequest{
	// ...
}


if err := client.GenerateBenefitUtilizationSummariesReportGenerateByReservationIdThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `OpenapisClient.GenerateBenefitUtilizationSummariesReportGenerateByReservationOrderId`

```go
ctx := context.TODO()
id := openapis.NewReservationOrderID("reservationOrderId")

payload := openapis.BenefitUtilizationSummariesRequest{
	// ...
}


if err := client.GenerateBenefitUtilizationSummariesReportGenerateByReservationOrderIdThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `OpenapisClient.GenerateBenefitUtilizationSummariesReportGenerateBySavingsPlanId`

```go
ctx := context.TODO()
id := openapis.NewSavingsPlanID("savingsPlanOrderId", "savingsPlanId")

payload := openapis.BenefitUtilizationSummariesRequest{
	// ...
}


if err := client.GenerateBenefitUtilizationSummariesReportGenerateBySavingsPlanIdThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `OpenapisClient.GenerateBenefitUtilizationSummariesReportGenerateBySavingsPlanOrderId`

```go
ctx := context.TODO()
id := openapis.NewSavingsPlanOrderID("savingsPlanOrderId")

payload := openapis.BenefitUtilizationSummariesRequest{
	// ...
}


if err := client.GenerateBenefitUtilizationSummariesReportGenerateBySavingsPlanOrderIdThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `OpenapisClient.GenerateCostDetailsReportCreateOperation`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

payload := openapis.GenerateCostDetailsReportRequestDefinition{
	// ...
}


if err := client.GenerateCostDetailsReportCreateOperationThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `OpenapisClient.GenerateDetailedCostReportCreateOperation`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

payload := openapis.GenerateDetailedCostReportDefinition{
	// ...
}


if err := client.GenerateDetailedCostReportCreateOperationThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `OpenapisClient.GenerateReservationDetailsReportByBillingAccountId`

```go
ctx := context.TODO()
id := openapis.NewBillingAccountID("billingAccountId")

if err := client.GenerateReservationDetailsReportByBillingAccountIdThenPoll(ctx, id, openapis.DefaultGenerateReservationDetailsReportByBillingAccountIdOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `OpenapisClient.GenerateReservationDetailsReportByBillingProfileId`

```go
ctx := context.TODO()
id := openapis.NewBillingProfileID("billingAccountName", "billingProfileName")

if err := client.GenerateReservationDetailsReportByBillingProfileIdThenPoll(ctx, id, openapis.DefaultGenerateReservationDetailsReportByBillingProfileIdOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `OpenapisClient.PriceSheetDownloadByBillingAccount`

```go
ctx := context.TODO()
id := openapis.NewBillingPeriodID("billingAccountId", "billingPeriodName")

if err := client.PriceSheetDownloadByBillingAccountThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `OpenapisClient.PriceSheetDownloadByBillingProfile`

```go
ctx := context.TODO()
id := openapis.NewBillingProfileID("billingAccountName", "billingProfileName")

if err := client.PriceSheetDownloadByBillingProfileThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `OpenapisClient.PriceSheetDownloadByInvoice`

```go
ctx := context.TODO()
id := openapis.NewInvoiceID("billingAccountName", "billingProfileName", "invoiceName")

if err := client.PriceSheetDownloadByInvoiceThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `OpenapisClient.QueryUsage`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

payload := openapis.QueryDefinition{
	// ...
}


read, err := client.QueryUsage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.QueryUsageByExternalCloudProviderType`

```go
ctx := context.TODO()
id := openapis.NewExternalCloudProviderTypeID("externalBillingAccounts", "externalCloudProviderId")

payload := openapis.QueryDefinition{
	// ...
}


read, err := client.QueryUsageByExternalCloudProviderType(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.ScheduledActionsCheckNameAvailability`

```go
ctx := context.TODO()

payload := openapis.CheckNameAvailabilityRequest{
	// ...
}


read, err := client.ScheduledActionsCheckNameAvailability(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.ScheduledActionsCheckNameAvailabilityByScope`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

payload := openapis.CheckNameAvailabilityRequest{
	// ...
}


read, err := client.ScheduledActionsCheckNameAvailabilityByScope(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
