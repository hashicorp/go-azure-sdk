
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/serviceprincipal` Documentation

The `serviceprincipal` SDK allows for interaction with the Azure Resource Manager Service `serviceprincipals` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/serviceprincipal"
```


### Client Initialization

```go
client := serviceprincipal.NewServicePrincipalClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ServicePrincipalClient.AddServicePrincipalKey`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.AddServicePrincipalKeyRequest{
	// ...
}


read, err := client.AddServicePrincipalKey(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.AddServicePrincipalPassword`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.AddServicePrincipalPasswordRequest{
	// ...
}


read, err := client.AddServicePrincipalPassword(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
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


### Example Usage: `ServicePrincipalClient.GetServicePrincipalsAvailableExtensionProperties`

```go
ctx := context.TODO()

payload := serviceprincipal.GetServicePrincipalsAvailableExtensionPropertiesRequest{
	// ...
}


// alternatively `client.GetServicePrincipalsAvailableExtensionProperties(ctx, payload)` can be used to do batched pagination
items, err := client.GetServicePrincipalsAvailableExtensionPropertiesComplete(ctx, payload)
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


### Example Usage: `ServicePrincipalClient.RemoveServicePrincipalKey`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.RemoveServicePrincipalKeyRequest{
	// ...
}


read, err := client.RemoveServicePrincipalKey(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.RemoveServicePrincipalPassword`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipal.RemoveServicePrincipalPasswordRequest{
	// ...
}


read, err := client.RemoveServicePrincipalPassword(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
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
