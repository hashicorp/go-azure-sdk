
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotectionsensitivitylabelsublabel` Documentation

The `siteinformationprotectionsensitivitylabelsublabel` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotectionsensitivitylabelsublabel"
```


### Client Initialization

```go
client := siteinformationprotectionsensitivitylabelsublabel.NewSiteInformationProtectionSensitivityLabelSublabelClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteInformationProtectionSensitivityLabelSublabelClient.CreateSiteInformationProtectionSensitivityLabelSublabel`

```go
ctx := context.TODO()
id := siteinformationprotectionsensitivitylabelsublabel.NewGroupIdSiteIdInformationProtectionSensitivityLabelID("groupId", "siteId", "sensitivityLabelId")

payload := siteinformationprotectionsensitivitylabelsublabel.SensitivityLabel{
	// ...
}


read, err := client.CreateSiteInformationProtectionSensitivityLabelSublabel(ctx, id, payload, siteinformationprotectionsensitivitylabelsublabel.DefaultCreateSiteInformationProtectionSensitivityLabelSublabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionSensitivityLabelSublabelClient.DeleteSiteInformationProtectionSensitivityLabelSublabel`

```go
ctx := context.TODO()
id := siteinformationprotectionsensitivitylabelsublabel.NewGroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelID("groupId", "siteId", "sensitivityLabelId", "sensitivityLabelId1")

read, err := client.DeleteSiteInformationProtectionSensitivityLabelSublabel(ctx, id, siteinformationprotectionsensitivitylabelsublabel.DefaultDeleteSiteInformationProtectionSensitivityLabelSublabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionSensitivityLabelSublabelClient.EvaluateSiteInformationProtectionSensitivityLabelSublabels`

```go
ctx := context.TODO()
id := siteinformationprotectionsensitivitylabelsublabel.NewGroupIdSiteIdInformationProtectionSensitivityLabelID("groupId", "siteId", "sensitivityLabelId")

payload := siteinformationprotectionsensitivitylabelsublabel.EvaluateSiteInformationProtectionSensitivityLabelSublabelsRequest{
	// ...
}


read, err := client.EvaluateSiteInformationProtectionSensitivityLabelSublabels(ctx, id, payload, siteinformationprotectionsensitivitylabelsublabel.DefaultEvaluateSiteInformationProtectionSensitivityLabelSublabelsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionSensitivityLabelSublabelClient.GetSiteInformationProtectionSensitivityLabelSublabel`

```go
ctx := context.TODO()
id := siteinformationprotectionsensitivitylabelsublabel.NewGroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelID("groupId", "siteId", "sensitivityLabelId", "sensitivityLabelId1")

read, err := client.GetSiteInformationProtectionSensitivityLabelSublabel(ctx, id, siteinformationprotectionsensitivitylabelsublabel.DefaultGetSiteInformationProtectionSensitivityLabelSublabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionSensitivityLabelSublabelClient.GetSiteInformationProtectionSensitivityLabelSublabelsCount`

```go
ctx := context.TODO()
id := siteinformationprotectionsensitivitylabelsublabel.NewGroupIdSiteIdInformationProtectionSensitivityLabelID("groupId", "siteId", "sensitivityLabelId")

read, err := client.GetSiteInformationProtectionSensitivityLabelSublabelsCount(ctx, id, siteinformationprotectionsensitivitylabelsublabel.DefaultGetSiteInformationProtectionSensitivityLabelSublabelsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionSensitivityLabelSublabelClient.ListSiteInformationProtectionSensitivityLabelSublabels`

```go
ctx := context.TODO()
id := siteinformationprotectionsensitivitylabelsublabel.NewGroupIdSiteIdInformationProtectionSensitivityLabelID("groupId", "siteId", "sensitivityLabelId")

// alternatively `client.ListSiteInformationProtectionSensitivityLabelSublabels(ctx, id, siteinformationprotectionsensitivitylabelsublabel.DefaultListSiteInformationProtectionSensitivityLabelSublabelsOperationOptions())` can be used to do batched pagination
items, err := client.ListSiteInformationProtectionSensitivityLabelSublabelsComplete(ctx, id, siteinformationprotectionsensitivitylabelsublabel.DefaultListSiteInformationProtectionSensitivityLabelSublabelsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteInformationProtectionSensitivityLabelSublabelClient.UpdateSiteInformationProtectionSensitivityLabelSublabel`

```go
ctx := context.TODO()
id := siteinformationprotectionsensitivitylabelsublabel.NewGroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelID("groupId", "siteId", "sensitivityLabelId", "sensitivityLabelId1")

payload := siteinformationprotectionsensitivitylabelsublabel.SensitivityLabel{
	// ...
}


read, err := client.UpdateSiteInformationProtectionSensitivityLabelSublabel(ctx, id, payload, siteinformationprotectionsensitivitylabelsublabel.DefaultUpdateSiteInformationProtectionSensitivityLabelSublabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
