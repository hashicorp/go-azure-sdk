
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


### Example Usage: `ServicePrincipalClient.AddTokenSigningCertificate`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.AddTokenSigningCertificateRequest{
	// ...
}


read, err := client.AddTokenSigningCertificate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.CheckMemberGroups`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.CheckMemberGroupsRequest{
	// ...
}


// alternatively `client.CheckMemberGroups(ctx, id, payload, serviceprincipal.DefaultCheckMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberGroupsComplete(ctx, id, payload, serviceprincipal.DefaultCheckMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ServicePrincipalClient.CheckMemberObjects`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.CheckMemberObjectsRequest{
	// ...
}


// alternatively `client.CheckMemberObjects(ctx, id, payload, serviceprincipal.DefaultCheckMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberObjectsComplete(ctx, id, payload, serviceprincipal.DefaultCheckMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ServicePrincipalClient.CreatePasswordSingleSignOnCredential`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.CreatePasswordSingleSignOnCredentialRequest{
	// ...
}


read, err := client.CreatePasswordSingleSignOnCredential(ctx, id, payload)
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


### Example Usage: `ServicePrincipalClient.DeletePasswordSingleSignOnCredential`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.DeletePasswordSingleSignOnCredentialRequest{
	// ...
}


read, err := client.DeletePasswordSingleSignOnCredential(ctx, id, payload)
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

read, err := client.DeleteServicePrincipal(ctx, id, serviceprincipal.DefaultDeleteServicePrincipalOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.GetByIds`

```go
ctx := context.TODO()

payload := serviceprincipal.GetByIdsRequest{
	// ...
}


// alternatively `client.GetByIds(ctx, payload, serviceprincipal.DefaultGetByIdsOperationOptions())` can be used to do batched pagination
items, err := client.GetByIdsComplete(ctx, payload, serviceprincipal.DefaultGetByIdsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ServicePrincipalClient.GetCount`

```go
ctx := context.TODO()


read, err := client.GetCount(ctx, serviceprincipal.DefaultGetCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.GetMemberGroups`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.GetMemberGroupsRequest{
	// ...
}


// alternatively `client.GetMemberGroups(ctx, id, payload, serviceprincipal.DefaultGetMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberGroupsComplete(ctx, id, payload, serviceprincipal.DefaultGetMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ServicePrincipalClient.GetMemberObjects`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.GetMemberObjectsRequest{
	// ...
}


// alternatively `client.GetMemberObjects(ctx, id, payload, serviceprincipal.DefaultGetMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberObjectsComplete(ctx, id, payload, serviceprincipal.DefaultGetMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ServicePrincipalClient.GetPasswordSingleSignOnCredential`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.GetPasswordSingleSignOnCredentialRequest{
	// ...
}


read, err := client.GetPasswordSingleSignOnCredential(ctx, id, payload)
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

read, err := client.GetServicePrincipal(ctx, id, serviceprincipal.DefaultGetServicePrincipalOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.GetUserOwnedObject`

```go
ctx := context.TODO()

payload := serviceprincipal.GetUserOwnedObjectRequest{
	// ...
}


read, err := client.GetUserOwnedObject(ctx, payload)
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


// alternatively `client.ListServicePrincipals(ctx, serviceprincipal.DefaultListServicePrincipalsOperationOptions())` can be used to do batched pagination
items, err := client.ListServicePrincipalsComplete(ctx, serviceprincipal.DefaultListServicePrincipalsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ServicePrincipalClient.Restore`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.RestoreRequest{
	// ...
}


read, err := client.Restore(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.UpdatePasswordSingleSignOnCredential`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.UpdatePasswordSingleSignOnCredentialRequest{
	// ...
}


read, err := client.UpdatePasswordSingleSignOnCredential(ctx, id, payload)
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


### Example Usage: `ServicePrincipalClient.ValidateProperty`

```go
ctx := context.TODO()

payload := serviceprincipal.ValidatePropertyRequest{
	// ...
}


read, err := client.ValidateProperty(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
