
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotectionsensitivitylabel` Documentation

The `siteinformationprotectionsensitivitylabel` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotectionsensitivitylabel"
```


### Client Initialization

```go
client := siteinformationprotectionsensitivitylabel.NewSiteInformationProtectionSensitivityLabelClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteInformationProtectionSensitivityLabelClient.CreateSiteInformationProtectionSensitivityLabel`

```go
ctx := context.TODO()
id := siteinformationprotectionsensitivitylabel.NewGroupIdSiteID("groupId", "siteId")

payload := siteinformationprotectionsensitivitylabel.SensitivityLabel{
	// ...
}


read, err := client.CreateSiteInformationProtectionSensitivityLabel(ctx, id, payload, siteinformationprotectionsensitivitylabel.DefaultCreateSiteInformationProtectionSensitivityLabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionSensitivityLabelClient.CreateSiteInformationProtectionSensitivityLabelComputeRightsAndInheritance`

```go
ctx := context.TODO()
id := siteinformationprotectionsensitivitylabel.NewGroupIdSiteID("groupId", "siteId")

payload := siteinformationprotectionsensitivitylabel.CreateSiteInformationProtectionSensitivityLabelComputeRightsAndInheritanceRequest{
	// ...
}


read, err := client.CreateSiteInformationProtectionSensitivityLabelComputeRightsAndInheritance(ctx, id, payload, siteinformationprotectionsensitivitylabel.DefaultCreateSiteInformationProtectionSensitivityLabelComputeRightsAndInheritanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionSensitivityLabelClient.DeleteSiteInformationProtectionSensitivityLabel`

```go
ctx := context.TODO()
id := siteinformationprotectionsensitivitylabel.NewGroupIdSiteIdInformationProtectionSensitivityLabelID("groupId", "siteId", "sensitivityLabelId")

read, err := client.DeleteSiteInformationProtectionSensitivityLabel(ctx, id, siteinformationprotectionsensitivitylabel.DefaultDeleteSiteInformationProtectionSensitivityLabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionSensitivityLabelClient.EvaluateSiteInformationProtectionSensitivityLabels`

```go
ctx := context.TODO()
id := siteinformationprotectionsensitivitylabel.NewGroupIdSiteID("groupId", "siteId")

payload := siteinformationprotectionsensitivitylabel.EvaluateSiteInformationProtectionSensitivityLabelsRequest{
	// ...
}


read, err := client.EvaluateSiteInformationProtectionSensitivityLabels(ctx, id, payload, siteinformationprotectionsensitivitylabel.DefaultEvaluateSiteInformationProtectionSensitivityLabelsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionSensitivityLabelClient.GetSiteInformationProtectionSensitivityLabel`

```go
ctx := context.TODO()
id := siteinformationprotectionsensitivitylabel.NewGroupIdSiteIdInformationProtectionSensitivityLabelID("groupId", "siteId", "sensitivityLabelId")

read, err := client.GetSiteInformationProtectionSensitivityLabel(ctx, id, siteinformationprotectionsensitivitylabel.DefaultGetSiteInformationProtectionSensitivityLabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionSensitivityLabelClient.GetSiteInformationProtectionSensitivityLabelsCount`

```go
ctx := context.TODO()
id := siteinformationprotectionsensitivitylabel.NewGroupIdSiteID("groupId", "siteId")

read, err := client.GetSiteInformationProtectionSensitivityLabelsCount(ctx, id, siteinformationprotectionsensitivitylabel.DefaultGetSiteInformationProtectionSensitivityLabelsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionSensitivityLabelClient.ListSiteInformationProtectionSensitivityLabels`

```go
ctx := context.TODO()
id := siteinformationprotectionsensitivitylabel.NewGroupIdSiteID("groupId", "siteId")

// alternatively `client.ListSiteInformationProtectionSensitivityLabels(ctx, id, siteinformationprotectionsensitivitylabel.DefaultListSiteInformationProtectionSensitivityLabelsOperationOptions())` can be used to do batched pagination
items, err := client.ListSiteInformationProtectionSensitivityLabelsComplete(ctx, id, siteinformationprotectionsensitivitylabel.DefaultListSiteInformationProtectionSensitivityLabelsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteInformationProtectionSensitivityLabelClient.UpdateSiteInformationProtectionSensitivityLabel`

```go
ctx := context.TODO()
id := siteinformationprotectionsensitivitylabel.NewGroupIdSiteIdInformationProtectionSensitivityLabelID("groupId", "siteId", "sensitivityLabelId")

payload := siteinformationprotectionsensitivitylabel.SensitivityLabel{
	// ...
}


read, err := client.UpdateSiteInformationProtectionSensitivityLabel(ctx, id, payload, siteinformationprotectionsensitivitylabel.DefaultUpdateSiteInformationProtectionSensitivityLabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
