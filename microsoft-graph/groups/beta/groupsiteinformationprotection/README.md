
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupsiteinformationprotection` Documentation

The `groupsiteinformationprotection` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupsiteinformationprotection"
```


### Client Initialization

```go
client := groupsiteinformationprotection.NewGroupSiteInformationProtectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupSiteInformationProtectionClient.CreateGroupByIdSiteByIdInformationProtectionDecryptBuffer`

```go
ctx := context.TODO()
id := groupsiteinformationprotection.NewGroupSiteID("groupIdValue", "siteIdValue")

payload := groupsiteinformationprotection.CreateGroupByIdSiteByIdInformationProtectionDecryptBufferRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdInformationProtectionDecryptBuffer(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteInformationProtectionClient.CreateGroupByIdSiteByIdInformationProtectionEncryptBuffer`

```go
ctx := context.TODO()
id := groupsiteinformationprotection.NewGroupSiteID("groupIdValue", "siteIdValue")

payload := groupsiteinformationprotection.CreateGroupByIdSiteByIdInformationProtectionEncryptBufferRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdInformationProtectionEncryptBuffer(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteInformationProtectionClient.CreateGroupByIdSiteByIdInformationProtectionSignDigest`

```go
ctx := context.TODO()
id := groupsiteinformationprotection.NewGroupSiteID("groupIdValue", "siteIdValue")

payload := groupsiteinformationprotection.CreateGroupByIdSiteByIdInformationProtectionSignDigestRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdInformationProtectionSignDigest(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteInformationProtectionClient.CreateGroupByIdSiteByIdInformationProtectionVerifySignature`

```go
ctx := context.TODO()
id := groupsiteinformationprotection.NewGroupSiteID("groupIdValue", "siteIdValue")

payload := groupsiteinformationprotection.CreateGroupByIdSiteByIdInformationProtectionVerifySignatureRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdInformationProtectionVerifySignature(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteInformationProtectionClient.DeleteGroupByIdSiteByIdInformationProtection`

```go
ctx := context.TODO()
id := groupsiteinformationprotection.NewGroupSiteID("groupIdValue", "siteIdValue")

read, err := client.DeleteGroupByIdSiteByIdInformationProtection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteInformationProtectionClient.GetGroupByIdSiteByIdInformationProtection`

```go
ctx := context.TODO()
id := groupsiteinformationprotection.NewGroupSiteID("groupIdValue", "siteIdValue")

read, err := client.GetGroupByIdSiteByIdInformationProtection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteInformationProtectionClient.UpdateGroupByIdSiteByIdInformationProtection`

```go
ctx := context.TODO()
id := groupsiteinformationprotection.NewGroupSiteID("groupIdValue", "siteIdValue")

payload := groupsiteinformationprotection.InformationProtection{
	// ...
}


read, err := client.UpdateGroupByIdSiteByIdInformationProtection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
