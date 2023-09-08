
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupsiteinformationprotectionsensitivitylabel` Documentation

The `groupsiteinformationprotectionsensitivitylabel` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupsiteinformationprotectionsensitivitylabel"
```


### Client Initialization

```go
client := groupsiteinformationprotectionsensitivitylabel.NewGroupSiteInformationProtectionSensitivityLabelClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupSiteInformationProtectionSensitivityLabelClient.CreateGroupByIdSiteByIdInformationProtectionSensitivityLabel`

```go
ctx := context.TODO()
id := groupsiteinformationprotectionsensitivitylabel.NewGroupSiteID("groupIdValue", "siteIdValue")

payload := groupsiteinformationprotectionsensitivitylabel.SensitivityLabel{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdInformationProtectionSensitivityLabel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteInformationProtectionSensitivityLabelClient.CreateGroupByIdSiteByIdInformationProtectionSensitivityLabelEvaluate`

```go
ctx := context.TODO()
id := groupsiteinformationprotectionsensitivitylabel.NewGroupSiteID("groupIdValue", "siteIdValue")

payload := groupsiteinformationprotectionsensitivitylabel.CreateGroupByIdSiteByIdInformationProtectionSensitivityLabelEvaluateRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdInformationProtectionSensitivityLabelEvaluate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteInformationProtectionSensitivityLabelClient.DeleteGroupByIdSiteByIdInformationProtectionSensitivityLabelById`

```go
ctx := context.TODO()
id := groupsiteinformationprotectionsensitivitylabel.NewGroupSiteInformationProtectionSensitivityLabelID("groupIdValue", "siteIdValue", "sensitivityLabelIdValue")

read, err := client.DeleteGroupByIdSiteByIdInformationProtectionSensitivityLabelById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteInformationProtectionSensitivityLabelClient.GetGroupByIdSiteByIdInformationProtectionSensitivityLabelById`

```go
ctx := context.TODO()
id := groupsiteinformationprotectionsensitivitylabel.NewGroupSiteInformationProtectionSensitivityLabelID("groupIdValue", "siteIdValue", "sensitivityLabelIdValue")

read, err := client.GetGroupByIdSiteByIdInformationProtectionSensitivityLabelById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteInformationProtectionSensitivityLabelClient.GetGroupByIdSiteByIdInformationProtectionSensitivityLabelCount`

```go
ctx := context.TODO()
id := groupsiteinformationprotectionsensitivitylabel.NewGroupSiteID("groupIdValue", "siteIdValue")

read, err := client.GetGroupByIdSiteByIdInformationProtectionSensitivityLabelCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteInformationProtectionSensitivityLabelClient.ListGroupByIdSiteByIdInformationProtectionSensitivityLabels`

```go
ctx := context.TODO()
id := groupsiteinformationprotectionsensitivitylabel.NewGroupSiteID("groupIdValue", "siteIdValue")

// alternatively `client.ListGroupByIdSiteByIdInformationProtectionSensitivityLabels(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdSiteByIdInformationProtectionSensitivityLabelsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupSiteInformationProtectionSensitivityLabelClient.UpdateGroupByIdSiteByIdInformationProtectionSensitivityLabelById`

```go
ctx := context.TODO()
id := groupsiteinformationprotectionsensitivitylabel.NewGroupSiteInformationProtectionSensitivityLabelID("groupIdValue", "siteIdValue", "sensitivityLabelIdValue")

payload := groupsiteinformationprotectionsensitivitylabel.SensitivityLabel{
	// ...
}


read, err := client.UpdateGroupByIdSiteByIdInformationProtectionSensitivityLabelById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
