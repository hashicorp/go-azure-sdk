
## `github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-09-01/vmwares` Documentation

The `vmwares` SDK allows for interaction with Azure Resource Manager `vmware` (API Version `2023-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-09-01/vmwares"
```


### Client Initialization

```go
client := vmwares.NewVMwaresClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VMwaresClient.WorkloadNetworksCreateDhcp`

```go
ctx := context.TODO()
id := vmwares.NewDhcpConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "dhcpIdValue")

payload := vmwares.WorkloadNetworkDhcp{
	// ...
}


if err := client.WorkloadNetworksCreateDhcpThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksCreateDnsService`

```go
ctx := context.TODO()
id := vmwares.NewDnsServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "dnsServiceIdValue")

payload := vmwares.WorkloadNetworkDnsService{
	// ...
}


if err := client.WorkloadNetworksCreateDnsServiceThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksCreateDnsZone`

```go
ctx := context.TODO()
id := vmwares.NewDnsZoneID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "dnsZoneIdValue")

payload := vmwares.WorkloadNetworkDnsZone{
	// ...
}


if err := client.WorkloadNetworksCreateDnsZoneThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksCreatePortMirroring`

```go
ctx := context.TODO()
id := vmwares.NewPortMirroringProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "portMirroringIdValue")

payload := vmwares.WorkloadNetworkPortMirroring{
	// ...
}


if err := client.WorkloadNetworksCreatePortMirroringThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksCreatePublicIP`

```go
ctx := context.TODO()
id := vmwares.NewPublicIPID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "publicIPIdValue")

payload := vmwares.WorkloadNetworkPublicIP{
	// ...
}


if err := client.WorkloadNetworksCreatePublicIPThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksCreateVMGroup`

```go
ctx := context.TODO()
id := vmwares.NewVMGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "vmGroupIdValue")

payload := vmwares.WorkloadNetworkVMGroup{
	// ...
}


if err := client.WorkloadNetworksCreateVMGroupThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksDeleteDhcp`

```go
ctx := context.TODO()
id := vmwares.NewDhcpConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "dhcpIdValue")

if err := client.WorkloadNetworksDeleteDhcpThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksDeleteDnsService`

```go
ctx := context.TODO()
id := vmwares.NewDnsServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "dnsServiceIdValue")

if err := client.WorkloadNetworksDeleteDnsServiceThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksDeleteDnsZone`

```go
ctx := context.TODO()
id := vmwares.NewDnsZoneID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "dnsZoneIdValue")

if err := client.WorkloadNetworksDeleteDnsZoneThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksDeletePortMirroring`

```go
ctx := context.TODO()
id := vmwares.NewPortMirroringProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "portMirroringIdValue")

if err := client.WorkloadNetworksDeletePortMirroringThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksDeletePublicIP`

```go
ctx := context.TODO()
id := vmwares.NewPublicIPID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "publicIPIdValue")

if err := client.WorkloadNetworksDeletePublicIPThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksDeleteVMGroup`

```go
ctx := context.TODO()
id := vmwares.NewVMGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "vmGroupIdValue")

if err := client.WorkloadNetworksDeleteVMGroupThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksGetDhcp`

```go
ctx := context.TODO()
id := vmwares.NewDhcpConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "dhcpIdValue")

read, err := client.WorkloadNetworksGetDhcp(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksGetDnsService`

```go
ctx := context.TODO()
id := vmwares.NewDnsServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "dnsServiceIdValue")

read, err := client.WorkloadNetworksGetDnsService(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksGetDnsZone`

```go
ctx := context.TODO()
id := vmwares.NewDnsZoneID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "dnsZoneIdValue")

read, err := client.WorkloadNetworksGetDnsZone(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksGetPortMirroring`

```go
ctx := context.TODO()
id := vmwares.NewPortMirroringProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "portMirroringIdValue")

read, err := client.WorkloadNetworksGetPortMirroring(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksGetPublicIP`

```go
ctx := context.TODO()
id := vmwares.NewPublicIPID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "publicIPIdValue")

read, err := client.WorkloadNetworksGetPublicIP(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksGetVMGroup`

```go
ctx := context.TODO()
id := vmwares.NewVMGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "vmGroupIdValue")

read, err := client.WorkloadNetworksGetVMGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksListDhcp`

```go
ctx := context.TODO()
id := vmwares.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue")

// alternatively `client.WorkloadNetworksListDhcp(ctx, id)` can be used to do batched pagination
items, err := client.WorkloadNetworksListDhcpComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksListDnsServices`

```go
ctx := context.TODO()
id := vmwares.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue")

// alternatively `client.WorkloadNetworksListDnsServices(ctx, id)` can be used to do batched pagination
items, err := client.WorkloadNetworksListDnsServicesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksListDnsZones`

```go
ctx := context.TODO()
id := vmwares.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue")

// alternatively `client.WorkloadNetworksListDnsZones(ctx, id)` can be used to do batched pagination
items, err := client.WorkloadNetworksListDnsZonesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksListPortMirroring`

```go
ctx := context.TODO()
id := vmwares.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue")

// alternatively `client.WorkloadNetworksListPortMirroring(ctx, id)` can be used to do batched pagination
items, err := client.WorkloadNetworksListPortMirroringComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksListPublicIPs`

```go
ctx := context.TODO()
id := vmwares.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue")

// alternatively `client.WorkloadNetworksListPublicIPs(ctx, id)` can be used to do batched pagination
items, err := client.WorkloadNetworksListPublicIPsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksListVMGroups`

```go
ctx := context.TODO()
id := vmwares.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue")

// alternatively `client.WorkloadNetworksListVMGroups(ctx, id)` can be used to do batched pagination
items, err := client.WorkloadNetworksListVMGroupsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksUpdateDhcp`

```go
ctx := context.TODO()
id := vmwares.NewDhcpConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "dhcpIdValue")

payload := vmwares.WorkloadNetworkDhcp{
	// ...
}


if err := client.WorkloadNetworksUpdateDhcpThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksUpdateDnsService`

```go
ctx := context.TODO()
id := vmwares.NewDnsServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "dnsServiceIdValue")

payload := vmwares.WorkloadNetworkDnsService{
	// ...
}


if err := client.WorkloadNetworksUpdateDnsServiceThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksUpdateDnsZone`

```go
ctx := context.TODO()
id := vmwares.NewDnsZoneID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "dnsZoneIdValue")

payload := vmwares.WorkloadNetworkDnsZone{
	// ...
}


if err := client.WorkloadNetworksUpdateDnsZoneThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksUpdatePortMirroring`

```go
ctx := context.TODO()
id := vmwares.NewPortMirroringProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "portMirroringIdValue")

payload := vmwares.WorkloadNetworkPortMirroring{
	// ...
}


if err := client.WorkloadNetworksUpdatePortMirroringThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksUpdateVMGroup`

```go
ctx := context.TODO()
id := vmwares.NewVMGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "vmGroupIdValue")

payload := vmwares.WorkloadNetworkVMGroup{
	// ...
}


if err := client.WorkloadNetworksUpdateVMGroupThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
