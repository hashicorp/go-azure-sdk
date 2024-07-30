
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotection` Documentation

The `siteinformationprotection` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotection"
```


### Client Initialization

```go
client := siteinformationprotection.NewSiteInformationProtectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteInformationProtectionClient.CreateSiteInformationProtectionDecryptBuffer`

```go
ctx := context.TODO()
id := siteinformationprotection.NewGroupIdSiteID("groupIdValue", "siteIdValue")

payload := siteinformationprotection.CreateSiteInformationProtectionDecryptBufferRequest{
	// ...
}


read, err := client.CreateSiteInformationProtectionDecryptBuffer(ctx, id, payload)
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
id := siteinformationprotection.NewGroupIdSiteID("groupIdValue", "siteIdValue")

payload := siteinformationprotection.CreateSiteInformationProtectionEncryptBufferRequest{
	// ...
}


read, err := client.CreateSiteInformationProtectionEncryptBuffer(ctx, id, payload)
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
id := siteinformationprotection.NewGroupIdSiteID("groupIdValue", "siteIdValue")

payload := siteinformationprotection.CreateSiteInformationProtectionSignDigestRequest{
	// ...
}


read, err := client.CreateSiteInformationProtectionSignDigest(ctx, id, payload)
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
id := siteinformationprotection.NewGroupIdSiteID("groupIdValue", "siteIdValue")

payload := siteinformationprotection.CreateSiteInformationProtectionVerifySignatureRequest{
	// ...
}


read, err := client.CreateSiteInformationProtectionVerifySignature(ctx, id, payload)
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
id := siteinformationprotection.NewGroupIdSiteID("groupIdValue", "siteIdValue")

read, err := client.DeleteSiteInformationProtection(ctx, id)
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
id := siteinformationprotection.NewGroupIdSiteID("groupIdValue", "siteIdValue")

read, err := client.GetSiteInformationProtection(ctx, id)
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
id := siteinformationprotection.NewGroupIdSiteID("groupIdValue", "siteIdValue")

payload := siteinformationprotection.InformationProtection{
	// ...
}


read, err := client.UpdateSiteInformationProtection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
