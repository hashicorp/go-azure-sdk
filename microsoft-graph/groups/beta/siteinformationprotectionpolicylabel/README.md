
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotectionpolicylabel` Documentation

The `siteinformationprotectionpolicylabel` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotectionpolicylabel"
```


### Client Initialization

```go
client := siteinformationprotectionpolicylabel.NewSiteInformationProtectionPolicyLabelClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteInformationProtectionPolicyLabelClient.CreateSiteInformationProtectionPolicyLabel`

```go
ctx := context.TODO()
id := siteinformationprotectionpolicylabel.NewGroupIdSiteID("groupIdValue", "siteIdValue")

payload := siteinformationprotectionpolicylabel.InformationProtectionLabel{
	// ...
}


read, err := client.CreateSiteInformationProtectionPolicyLabel(ctx, id, payload)
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
id := siteinformationprotectionpolicylabel.NewGroupIdSiteIdInformationProtectionPolicyLabelID("groupIdValue", "siteIdValue", "informationProtectionLabelIdValue")

read, err := client.DeleteSiteInformationProtectionPolicyLabel(ctx, id, siteinformationprotectionpolicylabel.DefaultDeleteSiteInformationProtectionPolicyLabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionPolicyLabelClient.ExtractSiteInformationProtectionPolicyLabel`

```go
ctx := context.TODO()
id := siteinformationprotectionpolicylabel.NewGroupIdSiteID("groupIdValue", "siteIdValue")

payload := siteinformationprotectionpolicylabel.ExtractSiteInformationProtectionPolicyLabelRequest{
	// ...
}


read, err := client.ExtractSiteInformationProtectionPolicyLabel(ctx, id, payload)
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
id := siteinformationprotectionpolicylabel.NewGroupIdSiteIdInformationProtectionPolicyLabelID("groupIdValue", "siteIdValue", "informationProtectionLabelIdValue")

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
id := siteinformationprotectionpolicylabel.NewGroupIdSiteID("groupIdValue", "siteIdValue")

read, err := client.GetSiteInformationProtectionPolicyLabelsCount(ctx, id, siteinformationprotectionpolicylabel.DefaultGetSiteInformationProtectionPolicyLabelsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionPolicyLabelClient.ListSiteInformationProtectionPolicyLabelEvaluateApplications`

```go
ctx := context.TODO()
id := siteinformationprotectionpolicylabel.NewGroupIdSiteID("groupIdValue", "siteIdValue")

payload := siteinformationprotectionpolicylabel.ListSiteInformationProtectionPolicyLabelEvaluateApplicationsRequest{
	// ...
}


// alternatively `client.ListSiteInformationProtectionPolicyLabelEvaluateApplications(ctx, id, payload, siteinformationprotectionpolicylabel.DefaultListSiteInformationProtectionPolicyLabelEvaluateApplicationsOperationOptions())` can be used to do batched pagination
items, err := client.ListSiteInformationProtectionPolicyLabelEvaluateApplicationsComplete(ctx, id, payload, siteinformationprotectionpolicylabel.DefaultListSiteInformationProtectionPolicyLabelEvaluateApplicationsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteInformationProtectionPolicyLabelClient.ListSiteInformationProtectionPolicyLabelEvaluateClassificationResults`

```go
ctx := context.TODO()
id := siteinformationprotectionpolicylabel.NewGroupIdSiteID("groupIdValue", "siteIdValue")

payload := siteinformationprotectionpolicylabel.ListSiteInformationProtectionPolicyLabelEvaluateClassificationResultsRequest{
	// ...
}


// alternatively `client.ListSiteInformationProtectionPolicyLabelEvaluateClassificationResults(ctx, id, payload, siteinformationprotectionpolicylabel.DefaultListSiteInformationProtectionPolicyLabelEvaluateClassificationResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListSiteInformationProtectionPolicyLabelEvaluateClassificationResultsComplete(ctx, id, payload, siteinformationprotectionpolicylabel.DefaultListSiteInformationProtectionPolicyLabelEvaluateClassificationResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteInformationProtectionPolicyLabelClient.ListSiteInformationProtectionPolicyLabelEvaluateRemovals`

```go
ctx := context.TODO()
id := siteinformationprotectionpolicylabel.NewGroupIdSiteID("groupIdValue", "siteIdValue")

payload := siteinformationprotectionpolicylabel.ListSiteInformationProtectionPolicyLabelEvaluateRemovalsRequest{
	// ...
}


// alternatively `client.ListSiteInformationProtectionPolicyLabelEvaluateRemovals(ctx, id, payload, siteinformationprotectionpolicylabel.DefaultListSiteInformationProtectionPolicyLabelEvaluateRemovalsOperationOptions())` can be used to do batched pagination
items, err := client.ListSiteInformationProtectionPolicyLabelEvaluateRemovalsComplete(ctx, id, payload, siteinformationprotectionpolicylabel.DefaultListSiteInformationProtectionPolicyLabelEvaluateRemovalsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteInformationProtectionPolicyLabelClient.ListSiteInformationProtectionPolicyLabels`

```go
ctx := context.TODO()
id := siteinformationprotectionpolicylabel.NewGroupIdSiteID("groupIdValue", "siteIdValue")

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
id := siteinformationprotectionpolicylabel.NewGroupIdSiteIdInformationProtectionPolicyLabelID("groupIdValue", "siteIdValue", "informationProtectionLabelIdValue")

payload := siteinformationprotectionpolicylabel.InformationProtectionLabel{
	// ...
}


read, err := client.UpdateSiteInformationProtectionPolicyLabel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
