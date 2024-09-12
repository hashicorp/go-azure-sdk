
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotectionsensitivitylabelsublabel` Documentation

The `siteinformationprotectionsensitivitylabelsublabel` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotectionsensitivitylabelsublabel"
```


### Client Initialization

```go
client := siteinformationprotectionsensitivitylabelsublabel.NewSiteInformationProtectionSensitivityLabelSublabelClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteInformationProtectionSensitivityLabelSublabelClient.CreateSiteInformationProtectionSensitivityLabelSublabel`

```go
ctx := context.TODO()
id := siteinformationprotectionsensitivitylabelsublabel.NewGroupIdSiteIdInformationProtectionSensitivityLabelID("groupIdValue", "siteIdValue", "sensitivityLabelIdValue")

payload := siteinformationprotectionsensitivitylabelsublabel.SensitivityLabel{
	// ...
}


read, err := client.CreateSiteInformationProtectionSensitivityLabelSublabel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionSensitivityLabelSublabelClient.CreateSiteInformationProtectionSensitivityLabelSublabelEvaluate`

```go
ctx := context.TODO()
id := siteinformationprotectionsensitivitylabelsublabel.NewGroupIdSiteIdInformationProtectionSensitivityLabelID("groupIdValue", "siteIdValue", "sensitivityLabelIdValue")

payload := siteinformationprotectionsensitivitylabelsublabel.CreateSiteInformationProtectionSensitivityLabelSublabelEvaluateRequest{
	// ...
}


read, err := client.CreateSiteInformationProtectionSensitivityLabelSublabelEvaluate(ctx, id, payload)
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
id := siteinformationprotectionsensitivitylabelsublabel.NewGroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelID("groupIdValue", "siteIdValue", "sensitivityLabelIdValue", "sensitivityLabelId1Value")

read, err := client.DeleteSiteInformationProtectionSensitivityLabelSublabel(ctx, id, siteinformationprotectionsensitivitylabelsublabel.DefaultDeleteSiteInformationProtectionSensitivityLabelSublabelOperationOptions())
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
id := siteinformationprotectionsensitivitylabelsublabel.NewGroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelID("groupIdValue", "siteIdValue", "sensitivityLabelIdValue", "sensitivityLabelId1Value")

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
id := siteinformationprotectionsensitivitylabelsublabel.NewGroupIdSiteIdInformationProtectionSensitivityLabelID("groupIdValue", "siteIdValue", "sensitivityLabelIdValue")

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
id := siteinformationprotectionsensitivitylabelsublabel.NewGroupIdSiteIdInformationProtectionSensitivityLabelID("groupIdValue", "siteIdValue", "sensitivityLabelIdValue")

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
id := siteinformationprotectionsensitivitylabelsublabel.NewGroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelID("groupIdValue", "siteIdValue", "sensitivityLabelIdValue", "sensitivityLabelId1Value")

payload := siteinformationprotectionsensitivitylabelsublabel.SensitivityLabel{
	// ...
}


read, err := client.UpdateSiteInformationProtectionSensitivityLabelSublabel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
