
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupsiteinformationprotectionpolicylabel` Documentation

The `groupsiteinformationprotectionpolicylabel` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupsiteinformationprotectionpolicylabel"
```


### Client Initialization

```go
client := groupsiteinformationprotectionpolicylabel.NewGroupSiteInformationProtectionPolicyLabelClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupSiteInformationProtectionPolicyLabelClient.CreateGroupByIdSiteByIdInformationProtectionPolicyLabel`

```go
ctx := context.TODO()
id := groupsiteinformationprotectionpolicylabel.NewGroupSiteID("groupIdValue", "siteIdValue")

payload := groupsiteinformationprotectionpolicylabel.InformationProtectionLabel{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdInformationProtectionPolicyLabel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteInformationProtectionPolicyLabelClient.CreateGroupByIdSiteByIdInformationProtectionPolicyLabelEvaluateApplication`

```go
ctx := context.TODO()
id := groupsiteinformationprotectionpolicylabel.NewGroupSiteID("groupIdValue", "siteIdValue")

payload := groupsiteinformationprotectionpolicylabel.CreateGroupByIdSiteByIdInformationProtectionPolicyLabelEvaluateApplicationRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdInformationProtectionPolicyLabelEvaluateApplication(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteInformationProtectionPolicyLabelClient.CreateGroupByIdSiteByIdInformationProtectionPolicyLabelEvaluateClassificationResult`

```go
ctx := context.TODO()
id := groupsiteinformationprotectionpolicylabel.NewGroupSiteID("groupIdValue", "siteIdValue")

payload := groupsiteinformationprotectionpolicylabel.CreateGroupByIdSiteByIdInformationProtectionPolicyLabelEvaluateClassificationResultRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdInformationProtectionPolicyLabelEvaluateClassificationResult(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteInformationProtectionPolicyLabelClient.CreateGroupByIdSiteByIdInformationProtectionPolicyLabelEvaluateRemoval`

```go
ctx := context.TODO()
id := groupsiteinformationprotectionpolicylabel.NewGroupSiteID("groupIdValue", "siteIdValue")

payload := groupsiteinformationprotectionpolicylabel.CreateGroupByIdSiteByIdInformationProtectionPolicyLabelEvaluateRemovalRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdInformationProtectionPolicyLabelEvaluateRemoval(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteInformationProtectionPolicyLabelClient.CreateGroupByIdSiteByIdInformationProtectionPolicyLabelExtractLabel`

```go
ctx := context.TODO()
id := groupsiteinformationprotectionpolicylabel.NewGroupSiteID("groupIdValue", "siteIdValue")

payload := groupsiteinformationprotectionpolicylabel.CreateGroupByIdSiteByIdInformationProtectionPolicyLabelExtractLabelRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdInformationProtectionPolicyLabelExtractLabel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteInformationProtectionPolicyLabelClient.DeleteGroupByIdSiteByIdInformationProtectionPolicyLabelById`

```go
ctx := context.TODO()
id := groupsiteinformationprotectionpolicylabel.NewGroupSiteInformationProtectionPolicyLabelID("groupIdValue", "siteIdValue", "informationProtectionLabelIdValue")

read, err := client.DeleteGroupByIdSiteByIdInformationProtectionPolicyLabelById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteInformationProtectionPolicyLabelClient.GetGroupByIdSiteByIdInformationProtectionPolicyLabelById`

```go
ctx := context.TODO()
id := groupsiteinformationprotectionpolicylabel.NewGroupSiteInformationProtectionPolicyLabelID("groupIdValue", "siteIdValue", "informationProtectionLabelIdValue")

read, err := client.GetGroupByIdSiteByIdInformationProtectionPolicyLabelById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteInformationProtectionPolicyLabelClient.GetGroupByIdSiteByIdInformationProtectionPolicyLabelCount`

```go
ctx := context.TODO()
id := groupsiteinformationprotectionpolicylabel.NewGroupSiteID("groupIdValue", "siteIdValue")

read, err := client.GetGroupByIdSiteByIdInformationProtectionPolicyLabelCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteInformationProtectionPolicyLabelClient.ListGroupByIdSiteByIdInformationProtectionPolicyLabels`

```go
ctx := context.TODO()
id := groupsiteinformationprotectionpolicylabel.NewGroupSiteID("groupIdValue", "siteIdValue")

// alternatively `client.ListGroupByIdSiteByIdInformationProtectionPolicyLabels(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdSiteByIdInformationProtectionPolicyLabelsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupSiteInformationProtectionPolicyLabelClient.UpdateGroupByIdSiteByIdInformationProtectionPolicyLabelById`

```go
ctx := context.TODO()
id := groupsiteinformationprotectionpolicylabel.NewGroupSiteInformationProtectionPolicyLabelID("groupIdValue", "siteIdValue", "informationProtectionLabelIdValue")

payload := groupsiteinformationprotectionpolicylabel.InformationProtectionLabel{
	// ...
}


read, err := client.UpdateGroupByIdSiteByIdInformationProtectionPolicyLabelById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
