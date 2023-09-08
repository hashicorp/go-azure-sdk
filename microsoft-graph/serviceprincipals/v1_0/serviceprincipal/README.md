
## `github.com/hashicorp/go-azure-sdk/resource-manager/serviceprincipals/v1.0/serviceprincipal` Documentation

The `serviceprincipal` SDK allows for interaction with the Azure Resource Manager Service `serviceprincipals` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/serviceprincipals/v1.0/serviceprincipal"
```


### Client Initialization

```go
client := serviceprincipal.NewServicePrincipalClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ServicePrincipalClient.AddServicePrincipalByIdKey`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.AddServicePrincipalByIdKeyRequest{
	// ...
}


read, err := client.AddServicePrincipalByIdKey(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.AddServicePrincipalByIdPassword`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.AddServicePrincipalByIdPasswordRequest{
	// ...
}


read, err := client.AddServicePrincipalByIdPassword(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.AddServicePrincipalByIdTokenSigningCertificate`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.AddServicePrincipalByIdTokenSigningCertificateRequest{
	// ...
}


read, err := client.AddServicePrincipalByIdTokenSigningCertificate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.CheckServicePrincipalByIdMemberGroup`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.CheckServicePrincipalByIdMemberGroupRequest{
	// ...
}


read, err := client.CheckServicePrincipalByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.CheckServicePrincipalByIdMemberObject`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.CheckServicePrincipalByIdMemberObjectRequest{
	// ...
}


read, err := client.CheckServicePrincipalByIdMemberObject(ctx, id, payload)
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


### Example Usage: `ServicePrincipalClient.DeleteServicePrincipalById`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

read, err := client.DeleteServicePrincipalById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.GetServicePrincipalById`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

read, err := client.GetServicePrincipalById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.GetServicePrincipalByIdMemberGroup`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.GetServicePrincipalByIdMemberGroupRequest{
	// ...
}


read, err := client.GetServicePrincipalByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.GetServicePrincipalByIdMemberObject`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.GetServicePrincipalByIdMemberObjectRequest{
	// ...
}


read, err := client.GetServicePrincipalByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.GetServicePrincipalCount`

```go
ctx := context.TODO()


read, err := client.GetServicePrincipalCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.GetServicePrincipalsAvailableExtensionProperties`

```go
ctx := context.TODO()


// alternatively `client.GetServicePrincipalsAvailableExtensionProperties(ctx)` can be used to do batched pagination
items, err := client.GetServicePrincipalsAvailableExtensionPropertiesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ServicePrincipalClient.GetServicePrincipalsByIds`

```go
ctx := context.TODO()


// alternatively `client.GetServicePrincipalsByIds(ctx)` can be used to do batched pagination
items, err := client.GetServicePrincipalsByIdsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
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


### Example Usage: `ServicePrincipalClient.RemoveServicePrincipalByIdKey`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.RemoveServicePrincipalByIdKeyRequest{
	// ...
}


read, err := client.RemoveServicePrincipalByIdKey(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.RemoveServicePrincipalByIdPassword`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.RemoveServicePrincipalByIdPasswordRequest{
	// ...
}


read, err := client.RemoveServicePrincipalByIdPassword(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.RestoreServicePrincipalById`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

read, err := client.RestoreServicePrincipalById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.UpdateServicePrincipalById`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.ServicePrincipal{
	// ...
}


read, err := client.UpdateServicePrincipalById(ctx, id, payload)
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
