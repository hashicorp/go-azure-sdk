
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotection` Documentation

The `siteinformationprotection` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotection"
```


### Client Initialization

```go
client := siteinformationprotection.NewSiteInformationProtectionClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteInformationProtectionClient.CreateSiteInformationProtectionDecryptBuffer`

```go
ctx := context.TODO()
id := siteinformationprotection.NewGroupIdSiteID("groupId", "siteId")

payload := siteinformationprotection.CreateSiteInformationProtectionDecryptBufferRequest{
	// ...
}


read, err := client.CreateSiteInformationProtectionDecryptBuffer(ctx, id, payload, siteinformationprotection.DefaultCreateSiteInformationProtectionDecryptBufferOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionClient.CreateSiteInformationProtectionEncryptBuffer`

```go
ctx := context.TODO()
id := siteinformationprotection.NewGroupIdSiteID("groupId", "siteId")

payload := siteinformationprotection.CreateSiteInformationProtectionEncryptBufferRequest{
	// ...
}


read, err := client.CreateSiteInformationProtectionEncryptBuffer(ctx, id, payload, siteinformationprotection.DefaultCreateSiteInformationProtectionEncryptBufferOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionClient.CreateSiteInformationProtectionSignDigest`

```go
ctx := context.TODO()
id := siteinformationprotection.NewGroupIdSiteID("groupId", "siteId")

payload := siteinformationprotection.CreateSiteInformationProtectionSignDigestRequest{
	// ...
}


read, err := client.CreateSiteInformationProtectionSignDigest(ctx, id, payload, siteinformationprotection.DefaultCreateSiteInformationProtectionSignDigestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionClient.CreateSiteInformationProtectionVerifySignature`

```go
ctx := context.TODO()
id := siteinformationprotection.NewGroupIdSiteID("groupId", "siteId")

payload := siteinformationprotection.CreateSiteInformationProtectionVerifySignatureRequest{
	// ...
}


read, err := client.CreateSiteInformationProtectionVerifySignature(ctx, id, payload, siteinformationprotection.DefaultCreateSiteInformationProtectionVerifySignatureOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionClient.DeleteSiteInformationProtection`

```go
ctx := context.TODO()
id := siteinformationprotection.NewGroupIdSiteID("groupId", "siteId")

read, err := client.DeleteSiteInformationProtection(ctx, id, siteinformationprotection.DefaultDeleteSiteInformationProtectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionClient.GetSiteInformationProtection`

```go
ctx := context.TODO()
id := siteinformationprotection.NewGroupIdSiteID("groupId", "siteId")

read, err := client.GetSiteInformationProtection(ctx, id, siteinformationprotection.DefaultGetSiteInformationProtectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionClient.UpdateSiteInformationProtection`

```go
ctx := context.TODO()
id := siteinformationprotection.NewGroupIdSiteID("groupId", "siteId")

payload := siteinformationprotection.InformationProtection{
	// ...
}


read, err := client.UpdateSiteInformationProtection(ctx, id, payload, siteinformationprotection.DefaultUpdateSiteInformationProtectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
