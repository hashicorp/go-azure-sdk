
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotectionsensitivitylabel` Documentation

The `siteinformationprotectionsensitivitylabel` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotectionsensitivitylabel"
```


### Client Initialization

```go
client := siteinformationprotectionsensitivitylabel.NewSiteInformationProtectionSensitivityLabelClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteInformationProtectionSensitivityLabelClient.CreateSiteInformationProtectionSensitivityLabel`

```go
ctx := context.TODO()
id := siteinformationprotectionsensitivitylabel.NewGroupIdSiteID("groupIdValue", "siteIdValue")

payload := siteinformationprotectionsensitivitylabel.SensitivityLabel{
	// ...
}


read, err := client.CreateSiteInformationProtectionSensitivityLabel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionSensitivityLabelClient.CreateSiteInformationProtectionSensitivityLabelEvaluate`

```go
ctx := context.TODO()
id := siteinformationprotectionsensitivitylabel.NewGroupIdSiteID("groupIdValue", "siteIdValue")

payload := siteinformationprotectionsensitivitylabel.CreateSiteInformationProtectionSensitivityLabelEvaluateRequest{
	// ...
}


read, err := client.CreateSiteInformationProtectionSensitivityLabelEvaluate(ctx, id, payload)
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
id := siteinformationprotectionsensitivitylabel.NewGroupIdSiteIdInformationProtectionSensitivityLabelID("groupIdValue", "siteIdValue", "sensitivityLabelIdValue")

read, err := client.DeleteSiteInformationProtectionSensitivityLabel(ctx, id, siteinformationprotectionsensitivitylabel.DefaultDeleteSiteInformationProtectionSensitivityLabelOperationOptions())
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
id := siteinformationprotectionsensitivitylabel.NewGroupIdSiteIdInformationProtectionSensitivityLabelID("groupIdValue", "siteIdValue", "sensitivityLabelIdValue")

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
id := siteinformationprotectionsensitivitylabel.NewGroupIdSiteID("groupIdValue", "siteIdValue")

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
id := siteinformationprotectionsensitivitylabel.NewGroupIdSiteID("groupIdValue", "siteIdValue")

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
id := siteinformationprotectionsensitivitylabel.NewGroupIdSiteIdInformationProtectionSensitivityLabelID("groupIdValue", "siteIdValue", "sensitivityLabelIdValue")

payload := siteinformationprotectionsensitivitylabel.SensitivityLabel{
	// ...
}


read, err := client.UpdateSiteInformationProtectionSensitivityLabel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
