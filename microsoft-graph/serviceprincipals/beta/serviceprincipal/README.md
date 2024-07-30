
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipal` Documentation

The `serviceprincipal` SDK allows for interaction with the Azure Resource Manager Service `serviceprincipals` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipal"
```


### Client Initialization

```go
client := serviceprincipal.NewServicePrincipalClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ServicePrincipalClient.AddServicePrincipalTokenSigningCertificate`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.AddServicePrincipalTokenSigningCertificateRequest{
	// ...
}


read, err := client.AddServicePrincipalTokenSigningCertificate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.CheckServicePrincipalMemberGroup`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.CheckServicePrincipalMemberGroupRequest{
	// ...
}


read, err := client.CheckServicePrincipalMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.CheckServicePrincipalMemberObject`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.CheckServicePrincipalMemberObjectRequest{
	// ...
}


read, err := client.CheckServicePrincipalMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.CreateCreatePasswordSingleSignOnCredential`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.CreateCreatePasswordSingleSignOnCredentialRequest{
	// ...
}


read, err := client.CreateCreatePasswordSingleSignOnCredential(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.CreateDeletePasswordSingleSignOnCredential`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.CreateDeletePasswordSingleSignOnCredentialRequest{
	// ...
}


read, err := client.CreateDeletePasswordSingleSignOnCredential(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.CreateServicePrincipal`

```go
ctx := context.TODO()

payload := serviceprincipal.ServicePrincipal{
	// ...
}


read, err := client.CreateServicePrincipal(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.DeleteServicePrincipal`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

read, err := client.DeleteServicePrincipal(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.GetCount`

```go
ctx := context.TODO()


read, err := client.GetCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.GetServicePrincipal`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

read, err := client.GetServicePrincipal(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.GetServicePrincipalMemberGroup`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.GetServicePrincipalMemberGroupRequest{
	// ...
}


read, err := client.GetServicePrincipalMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.GetServicePrincipalMemberObject`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.GetServicePrincipalMemberObjectRequest{
	// ...
}


read, err := client.GetServicePrincipalMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.GetServicePrincipalPasswordSingleSignOnCredential`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.GetServicePrincipalPasswordSingleSignOnCredentialRequest{
	// ...
}


read, err := client.GetServicePrincipalPasswordSingleSignOnCredential(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.GetServicePrincipalsByIds`

```go
ctx := context.TODO()

payload := serviceprincipal.GetServicePrincipalsByIdsRequest{
	// ...
}


// alternatively `client.GetServicePrincipalsByIds(ctx, payload)` can be used to do batched pagination
items, err := client.GetServicePrincipalsByIdsComplete(ctx, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ServicePrincipalClient.GetServicePrincipalsUserOwnedObject`

```go
ctx := context.TODO()

payload := serviceprincipal.GetServicePrincipalsUserOwnedObjectRequest{
	// ...
}


read, err := client.GetServicePrincipalsUserOwnedObject(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.ListServicePrincipals`

```go
ctx := context.TODO()


// alternatively `client.ListServicePrincipals(ctx)` can be used to do batched pagination
items, err := client.ListServicePrincipalsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ServicePrincipalClient.RestoreServicePrincipal`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

read, err := client.RestoreServicePrincipal(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.UpdateServicePrincipal`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.ServicePrincipal{
	// ...
}


read, err := client.UpdateServicePrincipal(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.UpdateServicePrincipalPasswordSingleSignOnCredential`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.UpdateServicePrincipalPasswordSingleSignOnCredentialRequest{
	// ...
}


read, err := client.UpdateServicePrincipalPasswordSingleSignOnCredential(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.ValidateServicePrincipalsProperty`

```go
ctx := context.TODO()

payload := serviceprincipal.ValidateServicePrincipalsPropertyRequest{
	// ...
}


read, err := client.ValidateServicePrincipalsProperty(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
