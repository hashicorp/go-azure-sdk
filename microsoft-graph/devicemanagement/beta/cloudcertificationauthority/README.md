
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/cloudcertificationauthority` Documentation

The `cloudcertificationauthority` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/cloudcertificationauthority"
```


### Client Initialization

```go
client := cloudcertificationauthority.NewCloudCertificationAuthorityClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CloudCertificationAuthorityClient.ChangeCloudCertificationAuthorityStatus`

```go
ctx := context.TODO()
id := cloudcertificationauthority.NewDeviceManagementCloudCertificationAuthorityID("cloudCertificationAuthorityId")

payload := cloudcertificationauthority.ChangeCloudCertificationAuthorityStatusRequest{
	// ...
}


read, err := client.ChangeCloudCertificationAuthorityStatus(ctx, id, payload, cloudcertificationauthority.DefaultChangeCloudCertificationAuthorityStatusOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudCertificationAuthorityClient.CreateCloudCertificationAuthority`

```go
ctx := context.TODO()

payload := cloudcertificationauthority.CloudCertificationAuthority{
	// ...
}


read, err := client.CreateCloudCertificationAuthority(ctx, payload, cloudcertificationauthority.DefaultCreateCloudCertificationAuthorityOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudCertificationAuthorityClient.CreateCloudCertificationAuthorityPatchCloudCertificationAuthority`

```go
ctx := context.TODO()
id := cloudcertificationauthority.NewDeviceManagementCloudCertificationAuthorityID("cloudCertificationAuthorityId")

read, err := client.CreateCloudCertificationAuthorityPatchCloudCertificationAuthority(ctx, id, cloudcertificationauthority.DefaultCreateCloudCertificationAuthorityPatchCloudCertificationAuthorityOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudCertificationAuthorityClient.CreateCloudCertificationAuthoritySearchLeafCertificateBySerialNumber`

```go
ctx := context.TODO()
id := cloudcertificationauthority.NewDeviceManagementCloudCertificationAuthorityID("cloudCertificationAuthorityId")

payload := cloudcertificationauthority.CreateCloudCertificationAuthoritySearchLeafCertificateBySerialNumberRequest{
	// ...
}


read, err := client.CreateCloudCertificationAuthoritySearchLeafCertificateBySerialNumber(ctx, id, payload, cloudcertificationauthority.DefaultCreateCloudCertificationAuthoritySearchLeafCertificateBySerialNumberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudCertificationAuthorityClient.CreateCloudCertificationAuthorityUploadExternallySignedCertificate`

```go
ctx := context.TODO()
id := cloudcertificationauthority.NewDeviceManagementCloudCertificationAuthorityID("cloudCertificationAuthorityId")

payload := cloudcertificationauthority.CreateCloudCertificationAuthorityUploadExternallySignedCertificateRequest{
	// ...
}


read, err := client.CreateCloudCertificationAuthorityUploadExternallySignedCertificate(ctx, id, payload, cloudcertificationauthority.DefaultCreateCloudCertificationAuthorityUploadExternallySignedCertificateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudCertificationAuthorityClient.DeleteCloudCertificationAuthority`

```go
ctx := context.TODO()
id := cloudcertificationauthority.NewDeviceManagementCloudCertificationAuthorityID("cloudCertificationAuthorityId")

read, err := client.DeleteCloudCertificationAuthority(ctx, id, cloudcertificationauthority.DefaultDeleteCloudCertificationAuthorityOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudCertificationAuthorityClient.GetCloudCertificationAuthority`

```go
ctx := context.TODO()
id := cloudcertificationauthority.NewDeviceManagementCloudCertificationAuthorityID("cloudCertificationAuthorityId")

read, err := client.GetCloudCertificationAuthority(ctx, id, cloudcertificationauthority.DefaultGetCloudCertificationAuthorityOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudCertificationAuthorityClient.GetCloudCertificationAuthorityAllCloudCertificationAuthorities`

```go
ctx := context.TODO()
id := cloudcertificationauthority.NewDeviceManagementCloudCertificationAuthorityID("cloudCertificationAuthorityId")

// alternatively `client.GetCloudCertificationAuthorityAllCloudCertificationAuthorities(ctx, id, cloudcertificationauthority.DefaultGetCloudCertificationAuthorityAllCloudCertificationAuthoritiesOperationOptions())` can be used to do batched pagination
items, err := client.GetCloudCertificationAuthorityAllCloudCertificationAuthoritiesComplete(ctx, id, cloudcertificationauthority.DefaultGetCloudCertificationAuthorityAllCloudCertificationAuthoritiesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CloudCertificationAuthorityClient.GetCloudCertificationAuthorityAllLeafCertificates`

```go
ctx := context.TODO()
id := cloudcertificationauthority.NewDeviceManagementCloudCertificationAuthorityID("cloudCertificationAuthorityId")

// alternatively `client.GetCloudCertificationAuthorityAllLeafCertificates(ctx, id, cloudcertificationauthority.DefaultGetCloudCertificationAuthorityAllLeafCertificatesOperationOptions())` can be used to do batched pagination
items, err := client.GetCloudCertificationAuthorityAllLeafCertificatesComplete(ctx, id, cloudcertificationauthority.DefaultGetCloudCertificationAuthorityAllLeafCertificatesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CloudCertificationAuthorityClient.GetCloudCertificationAuthorityCloudCertificationAuthority`

```go
ctx := context.TODO()
id := cloudcertificationauthority.NewDeviceManagementCloudCertificationAuthorityID("cloudCertificationAuthorityId")

read, err := client.GetCloudCertificationAuthorityCloudCertificationAuthority(ctx, id, cloudcertificationauthority.DefaultGetCloudCertificationAuthorityCloudCertificationAuthorityOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudCertificationAuthorityClient.GetCloudCertificationAuthorityCount`

```go
ctx := context.TODO()


read, err := client.GetCloudCertificationAuthorityCount(ctx, cloudcertificationauthority.DefaultGetCloudCertificationAuthorityCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudCertificationAuthorityClient.ListCloudCertificationAuthorities`

```go
ctx := context.TODO()


// alternatively `client.ListCloudCertificationAuthorities(ctx, cloudcertificationauthority.DefaultListCloudCertificationAuthoritiesOperationOptions())` can be used to do batched pagination
items, err := client.ListCloudCertificationAuthoritiesComplete(ctx, cloudcertificationauthority.DefaultListCloudCertificationAuthoritiesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CloudCertificationAuthorityClient.ListCloudCertificationAuthorityPostCloudCertificationAuthorities`

```go
ctx := context.TODO()
id := cloudcertificationauthority.NewDeviceManagementCloudCertificationAuthorityID("cloudCertificationAuthorityId")

// alternatively `client.ListCloudCertificationAuthorityPostCloudCertificationAuthorities(ctx, id, cloudcertificationauthority.DefaultListCloudCertificationAuthorityPostCloudCertificationAuthoritiesOperationOptions())` can be used to do batched pagination
items, err := client.ListCloudCertificationAuthorityPostCloudCertificationAuthoritiesComplete(ctx, id, cloudcertificationauthority.DefaultListCloudCertificationAuthorityPostCloudCertificationAuthoritiesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CloudCertificationAuthorityClient.RevokeCloudCertificationAuthorityCertificate`

```go
ctx := context.TODO()
id := cloudcertificationauthority.NewDeviceManagementCloudCertificationAuthorityID("cloudCertificationAuthorityId")

payload := cloudcertificationauthority.RevokeCloudCertificationAuthorityCertificateRequest{
	// ...
}


read, err := client.RevokeCloudCertificationAuthorityCertificate(ctx, id, payload, cloudcertificationauthority.DefaultRevokeCloudCertificationAuthorityCertificateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudCertificationAuthorityClient.RevokeCloudCertificationAuthorityLeafCertificate`

```go
ctx := context.TODO()
id := cloudcertificationauthority.NewDeviceManagementCloudCertificationAuthorityID("cloudCertificationAuthorityId")

payload := cloudcertificationauthority.RevokeCloudCertificationAuthorityLeafCertificateRequest{
	// ...
}


read, err := client.RevokeCloudCertificationAuthorityLeafCertificate(ctx, id, payload, cloudcertificationauthority.DefaultRevokeCloudCertificationAuthorityLeafCertificateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudCertificationAuthorityClient.RevokeCloudCertificationAuthorityLeafCertificateBySerialNumber`

```go
ctx := context.TODO()
id := cloudcertificationauthority.NewDeviceManagementCloudCertificationAuthorityID("cloudCertificationAuthorityId")

payload := cloudcertificationauthority.RevokeCloudCertificationAuthorityLeafCertificateBySerialNumberRequest{
	// ...
}


read, err := client.RevokeCloudCertificationAuthorityLeafCertificateBySerialNumber(ctx, id, payload, cloudcertificationauthority.DefaultRevokeCloudCertificationAuthorityLeafCertificateBySerialNumberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudCertificationAuthorityClient.UpdateCloudCertificationAuthority`

```go
ctx := context.TODO()
id := cloudcertificationauthority.NewDeviceManagementCloudCertificationAuthorityID("cloudCertificationAuthorityId")

payload := cloudcertificationauthority.CloudCertificationAuthority{
	// ...
}


read, err := client.UpdateCloudCertificationAuthority(ctx, id, payload, cloudcertificationauthority.DefaultUpdateCloudCertificationAuthorityOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
