
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotectionpolicylabel` Documentation

The `siteinformationprotectionpolicylabel` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotectionpolicylabel"
```


### Client Initialization

```go
client := siteinformationprotectionpolicylabel.NewSiteInformationProtectionPolicyLabelClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteInformationProtectionPolicyLabelClient.CreateSiteInformationProtectionPolicyLabel`

```go
ctx := context.TODO()
id := siteinformationprotectionpolicylabel.NewGroupIdSiteID("groupId", "siteId")

payload := siteinformationprotectionpolicylabel.InformationProtectionLabel{
	// ...
}


read, err := client.CreateSiteInformationProtectionPolicyLabel(ctx, id, payload, siteinformationprotectionpolicylabel.DefaultCreateSiteInformationProtectionPolicyLabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionPolicyLabelClient.DeleteSiteInformationProtectionPolicyLabel`

```go
ctx := context.TODO()
id := siteinformationprotectionpolicylabel.NewGroupIdSiteIdInformationProtectionPolicyLabelID("groupId", "siteId", "informationProtectionLabelId")

read, err := client.DeleteSiteInformationProtectionPolicyLabel(ctx, id, siteinformationprotectionpolicylabel.DefaultDeleteSiteInformationProtectionPolicyLabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionPolicyLabelClient.EvaluateSiteInformationProtectionPolicyLabelsApplications`

```go
ctx := context.TODO()
id := siteinformationprotectionpolicylabel.NewGroupIdSiteID("groupId", "siteId")

payload := siteinformationprotectionpolicylabel.EvaluateSiteInformationProtectionPolicyLabelsApplicationsRequest{
	// ...
}


// alternatively `client.EvaluateSiteInformationProtectionPolicyLabelsApplications(ctx, id, payload, siteinformationprotectionpolicylabel.DefaultEvaluateSiteInformationProtectionPolicyLabelsApplicationsOperationOptions())` can be used to do batched pagination
items, err := client.EvaluateSiteInformationProtectionPolicyLabelsApplicationsComplete(ctx, id, payload, siteinformationprotectionpolicylabel.DefaultEvaluateSiteInformationProtectionPolicyLabelsApplicationsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteInformationProtectionPolicyLabelClient.EvaluateSiteInformationProtectionPolicyLabelsClassificationResults`

```go
ctx := context.TODO()
id := siteinformationprotectionpolicylabel.NewGroupIdSiteID("groupId", "siteId")

payload := siteinformationprotectionpolicylabel.EvaluateSiteInformationProtectionPolicyLabelsClassificationResultsRequest{
	// ...
}


// alternatively `client.EvaluateSiteInformationProtectionPolicyLabelsClassificationResults(ctx, id, payload, siteinformationprotectionpolicylabel.DefaultEvaluateSiteInformationProtectionPolicyLabelsClassificationResultsOperationOptions())` can be used to do batched pagination
items, err := client.EvaluateSiteInformationProtectionPolicyLabelsClassificationResultsComplete(ctx, id, payload, siteinformationprotectionpolicylabel.DefaultEvaluateSiteInformationProtectionPolicyLabelsClassificationResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteInformationProtectionPolicyLabelClient.EvaluateSiteInformationProtectionPolicyLabelsRemovals`

```go
ctx := context.TODO()
id := siteinformationprotectionpolicylabel.NewGroupIdSiteID("groupId", "siteId")

payload := siteinformationprotectionpolicylabel.EvaluateSiteInformationProtectionPolicyLabelsRemovalsRequest{
	// ...
}


// alternatively `client.EvaluateSiteInformationProtectionPolicyLabelsRemovals(ctx, id, payload, siteinformationprotectionpolicylabel.DefaultEvaluateSiteInformationProtectionPolicyLabelsRemovalsOperationOptions())` can be used to do batched pagination
items, err := client.EvaluateSiteInformationProtectionPolicyLabelsRemovalsComplete(ctx, id, payload, siteinformationprotectionpolicylabel.DefaultEvaluateSiteInformationProtectionPolicyLabelsRemovalsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteInformationProtectionPolicyLabelClient.ExtractSiteInformationProtectionPolicyLabelsLabel`

```go
ctx := context.TODO()
id := siteinformationprotectionpolicylabel.NewGroupIdSiteID("groupId", "siteId")

payload := siteinformationprotectionpolicylabel.ExtractSiteInformationProtectionPolicyLabelsLabelRequest{
	// ...
}


read, err := client.ExtractSiteInformationProtectionPolicyLabelsLabel(ctx, id, payload, siteinformationprotectionpolicylabel.DefaultExtractSiteInformationProtectionPolicyLabelsLabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionPolicyLabelClient.GetSiteInformationProtectionPolicyLabel`

```go
ctx := context.TODO()
id := siteinformationprotectionpolicylabel.NewGroupIdSiteIdInformationProtectionPolicyLabelID("groupId", "siteId", "informationProtectionLabelId")

read, err := client.GetSiteInformationProtectionPolicyLabel(ctx, id, siteinformationprotectionpolicylabel.DefaultGetSiteInformationProtectionPolicyLabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionPolicyLabelClient.GetSiteInformationProtectionPolicyLabelsCount`

```go
ctx := context.TODO()
id := siteinformationprotectionpolicylabel.NewGroupIdSiteID("groupId", "siteId")

read, err := client.GetSiteInformationProtectionPolicyLabelsCount(ctx, id, siteinformationprotectionpolicylabel.DefaultGetSiteInformationProtectionPolicyLabelsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionPolicyLabelClient.ListSiteInformationProtectionPolicyLabels`

```go
ctx := context.TODO()
id := siteinformationprotectionpolicylabel.NewGroupIdSiteID("groupId", "siteId")

// alternatively `client.ListSiteInformationProtectionPolicyLabels(ctx, id, siteinformationprotectionpolicylabel.DefaultListSiteInformationProtectionPolicyLabelsOperationOptions())` can be used to do batched pagination
items, err := client.ListSiteInformationProtectionPolicyLabelsComplete(ctx, id, siteinformationprotectionpolicylabel.DefaultListSiteInformationProtectionPolicyLabelsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteInformationProtectionPolicyLabelClient.UpdateSiteInformationProtectionPolicyLabel`

```go
ctx := context.TODO()
id := siteinformationprotectionpolicylabel.NewGroupIdSiteIdInformationProtectionPolicyLabelID("groupId", "siteId", "informationProtectionLabelId")

payload := siteinformationprotectionpolicylabel.InformationProtectionLabel{
	// ...
}


read, err := client.UpdateSiteInformationProtectionPolicyLabel(ctx, id, payload, siteinformationprotectionpolicylabel.DefaultUpdateSiteInformationProtectionPolicyLabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
