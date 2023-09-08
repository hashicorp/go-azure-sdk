
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupsiteinformationprotectionsensitivitylabelsublabel` Documentation

The `groupsiteinformationprotectionsensitivitylabelsublabel` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupsiteinformationprotectionsensitivitylabelsublabel"
```


### Client Initialization

```go
client := groupsiteinformationprotectionsensitivitylabelsublabel.NewGroupSiteInformationProtectionSensitivityLabelSublabelClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupSiteInformationProtectionSensitivityLabelSublabelClient.CreateGroupByIdSiteByIdInformationProtectionSensitivityLabelByIdSublabel`

```go
ctx := context.TODO()
id := groupsiteinformationprotectionsensitivitylabelsublabel.NewGroupSiteInformationProtectionSensitivityLabelID("groupIdValue", "siteIdValue", "sensitivityLabelIdValue")

payload := groupsiteinformationprotectionsensitivitylabelsublabel.SensitivityLabel{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdInformationProtectionSensitivityLabelByIdSublabel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteInformationProtectionSensitivityLabelSublabelClient.CreateGroupByIdSiteByIdInformationProtectionSensitivityLabelByIdSublabelEvaluate`

```go
ctx := context.TODO()
id := groupsiteinformationprotectionsensitivitylabelsublabel.NewGroupSiteInformationProtectionSensitivityLabelID("groupIdValue", "siteIdValue", "sensitivityLabelIdValue")

payload := groupsiteinformationprotectionsensitivitylabelsublabel.CreateGroupByIdSiteByIdInformationProtectionSensitivityLabelByIdSublabelEvaluateRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdInformationProtectionSensitivityLabelByIdSublabelEvaluate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteInformationProtectionSensitivityLabelSublabelClient.DeleteGroupByIdSiteByIdInformationProtectionSensitivityLabelByIdSublabelById`

```go
ctx := context.TODO()
id := groupsiteinformationprotectionsensitivitylabelsublabel.NewGroupSiteInformationProtectionSensitivityLabelSublabelID("groupIdValue", "siteIdValue", "sensitivityLabelIdValue", "sensitivityLabelId1Value")

read, err := client.DeleteGroupByIdSiteByIdInformationProtectionSensitivityLabelByIdSublabelById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteInformationProtectionSensitivityLabelSublabelClient.GetGroupByIdSiteByIdInformationProtectionSensitivityLabelByIdSublabelById`

```go
ctx := context.TODO()
id := groupsiteinformationprotectionsensitivitylabelsublabel.NewGroupSiteInformationProtectionSensitivityLabelSublabelID("groupIdValue", "siteIdValue", "sensitivityLabelIdValue", "sensitivityLabelId1Value")

read, err := client.GetGroupByIdSiteByIdInformationProtectionSensitivityLabelByIdSublabelById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteInformationProtectionSensitivityLabelSublabelClient.GetGroupByIdSiteByIdInformationProtectionSensitivityLabelByIdSublabelCount`

```go
ctx := context.TODO()
id := groupsiteinformationprotectionsensitivitylabelsublabel.NewGroupSiteInformationProtectionSensitivityLabelID("groupIdValue", "siteIdValue", "sensitivityLabelIdValue")

read, err := client.GetGroupByIdSiteByIdInformationProtectionSensitivityLabelByIdSublabelCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteInformationProtectionSensitivityLabelSublabelClient.ListGroupByIdSiteByIdInformationProtectionSensitivityLabelByIdSublabels`

```go
ctx := context.TODO()
id := groupsiteinformationprotectionsensitivitylabelsublabel.NewGroupSiteInformationProtectionSensitivityLabelID("groupIdValue", "siteIdValue", "sensitivityLabelIdValue")

// alternatively `client.ListGroupByIdSiteByIdInformationProtectionSensitivityLabelByIdSublabels(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdSiteByIdInformationProtectionSensitivityLabelByIdSublabelsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupSiteInformationProtectionSensitivityLabelSublabelClient.UpdateGroupByIdSiteByIdInformationProtectionSensitivityLabelByIdSublabelById`

```go
ctx := context.TODO()
id := groupsiteinformationprotectionsensitivitylabelsublabel.NewGroupSiteInformationProtectionSensitivityLabelSublabelID("groupIdValue", "siteIdValue", "sensitivityLabelIdValue", "sensitivityLabelId1Value")

payload := groupsiteinformationprotectionsensitivitylabelsublabel.SensitivityLabel{
	// ...
}


read, err := client.UpdateGroupByIdSiteByIdInformationProtectionSensitivityLabelByIdSublabelById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
