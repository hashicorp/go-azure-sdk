
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/publickeyinfrastructurecertificatebasedauthconfiguration` Documentation

The `publickeyinfrastructurecertificatebasedauthconfiguration` SDK allows for interaction with Microsoft Graph `directory` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/publickeyinfrastructurecertificatebasedauthconfiguration"
```


### Client Initialization

```go
client := publickeyinfrastructurecertificatebasedauthconfiguration.NewPublicKeyInfrastructureCertificateBasedAuthConfigurationClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PublicKeyInfrastructureCertificateBasedAuthConfigurationClient.CreatePublicKeyInfrastructureCertificateBasedAuthConfiguration`

```go
ctx := context.TODO()

payload := publickeyinfrastructurecertificatebasedauthconfiguration.CertificateBasedAuthPki{
	// ...
}


read, err := client.CreatePublicKeyInfrastructureCertificateBasedAuthConfiguration(ctx, payload, publickeyinfrastructurecertificatebasedauthconfiguration.DefaultCreatePublicKeyInfrastructureCertificateBasedAuthConfigurationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PublicKeyInfrastructureCertificateBasedAuthConfigurationClient.CreatePublicKeyInfrastructureCertificateBasedAuthConfigurationUpload`

```go
ctx := context.TODO()
id := publickeyinfrastructurecertificatebasedauthconfiguration.NewDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationID("certificateBasedAuthPkiId")

payload := publickeyinfrastructurecertificatebasedauthconfiguration.CreatePublicKeyInfrastructureCertificateBasedAuthConfigurationUploadRequest{
	// ...
}


read, err := client.CreatePublicKeyInfrastructureCertificateBasedAuthConfigurationUpload(ctx, id, payload, publickeyinfrastructurecertificatebasedauthconfiguration.DefaultCreatePublicKeyInfrastructureCertificateBasedAuthConfigurationUploadOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PublicKeyInfrastructureCertificateBasedAuthConfigurationClient.DeletePublicKeyInfrastructureCertificateBasedAuthConfiguration`

```go
ctx := context.TODO()
id := publickeyinfrastructurecertificatebasedauthconfiguration.NewDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationID("certificateBasedAuthPkiId")

read, err := client.DeletePublicKeyInfrastructureCertificateBasedAuthConfiguration(ctx, id, publickeyinfrastructurecertificatebasedauthconfiguration.DefaultDeletePublicKeyInfrastructureCertificateBasedAuthConfigurationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PublicKeyInfrastructureCertificateBasedAuthConfigurationClient.GetPublicKeyInfrastructureCertificateBasedAuthConfiguration`

```go
ctx := context.TODO()
id := publickeyinfrastructurecertificatebasedauthconfiguration.NewDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationID("certificateBasedAuthPkiId")

read, err := client.GetPublicKeyInfrastructureCertificateBasedAuthConfiguration(ctx, id, publickeyinfrastructurecertificatebasedauthconfiguration.DefaultGetPublicKeyInfrastructureCertificateBasedAuthConfigurationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PublicKeyInfrastructureCertificateBasedAuthConfigurationClient.GetPublicKeyInfrastructureCertificateBasedAuthConfigurationsCount`

```go
ctx := context.TODO()


read, err := client.GetPublicKeyInfrastructureCertificateBasedAuthConfigurationsCount(ctx, publickeyinfrastructurecertificatebasedauthconfiguration.DefaultGetPublicKeyInfrastructureCertificateBasedAuthConfigurationsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PublicKeyInfrastructureCertificateBasedAuthConfigurationClient.ListPublicKeyInfrastructureCertificateBasedAuthConfigurations`

```go
ctx := context.TODO()


// alternatively `client.ListPublicKeyInfrastructureCertificateBasedAuthConfigurations(ctx, publickeyinfrastructurecertificatebasedauthconfiguration.DefaultListPublicKeyInfrastructureCertificateBasedAuthConfigurationsOperationOptions())` can be used to do batched pagination
items, err := client.ListPublicKeyInfrastructureCertificateBasedAuthConfigurationsComplete(ctx, publickeyinfrastructurecertificatebasedauthconfiguration.DefaultListPublicKeyInfrastructureCertificateBasedAuthConfigurationsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PublicKeyInfrastructureCertificateBasedAuthConfigurationClient.UpdatePublicKeyInfrastructureCertificateBasedAuthConfiguration`

```go
ctx := context.TODO()
id := publickeyinfrastructurecertificatebasedauthconfiguration.NewDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationID("certificateBasedAuthPkiId")

payload := publickeyinfrastructurecertificatebasedauthconfiguration.CertificateBasedAuthPki{
	// ...
}


read, err := client.UpdatePublicKeyInfrastructureCertificateBasedAuthConfiguration(ctx, id, payload, publickeyinfrastructurecertificatebasedauthconfiguration.DefaultUpdatePublicKeyInfrastructureCertificateBasedAuthConfigurationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
