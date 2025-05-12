
## `github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/vmwares` Documentation

The `vmwares` SDK allows for interaction with Azure Resource Manager `vmware` (API Version `2024-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/vmwares"
```


### Client Initialization

```go
client := vmwares.NewVMwaresClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VMwaresClient.SkusList`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.SkusList(ctx, id)` can be used to do batched pagination
items, err := client.SkusListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksCreateDhcp`

```go
ctx := context.TODO()
id := vmwares.NewDhcpConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "dhcpId")

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
id := vmwares.NewDnsServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "dnsServiceId")

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
id := vmwares.NewDnsZoneID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "dnsZoneId")

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
id := vmwares.NewPortMirroringProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "portMirroringId")

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
id := vmwares.NewPublicIPID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "publicIPId")

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
id := vmwares.NewVMGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "vmGroupId")

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
id := vmwares.NewDhcpConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "dhcpId")

if err := client.WorkloadNetworksDeleteDhcpThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksDeleteDnsService`

```go
ctx := context.TODO()
id := vmwares.NewDnsServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "dnsServiceId")

if err := client.WorkloadNetworksDeleteDnsServiceThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksDeleteDnsZone`

```go
ctx := context.TODO()
id := vmwares.NewDnsZoneID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "dnsZoneId")

if err := client.WorkloadNetworksDeleteDnsZoneThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksDeletePortMirroring`

```go
ctx := context.TODO()
id := vmwares.NewPortMirroringProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "portMirroringId")

if err := client.WorkloadNetworksDeletePortMirroringThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksDeletePublicIP`

```go
ctx := context.TODO()
id := vmwares.NewPublicIPID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "publicIPId")

if err := client.WorkloadNetworksDeletePublicIPThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksDeleteVMGroup`

```go
ctx := context.TODO()
id := vmwares.NewVMGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "vmGroupId")

if err := client.WorkloadNetworksDeleteVMGroupThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VMwaresClient.WorkloadNetworksGetDhcp`

```go
ctx := context.TODO()
id := vmwares.NewDhcpConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "dhcpId")

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
id := vmwares.NewDnsServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "dnsServiceId")

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
id := vmwares.NewDnsZoneID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "dnsZoneId")

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
id := vmwares.NewPortMirroringProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "portMirroringId")

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
id := vmwares.NewPublicIPID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "publicIPId")

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
id := vmwares.NewVMGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "vmGroupId")

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
id := vmwares.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName")

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
id := vmwares.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName")

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
id := vmwares.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName")

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
id := vmwares.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName")

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
id := vmwares.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName")

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
id := vmwares.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName")

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
id := vmwares.NewDhcpConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "dhcpId")

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
id := vmwares.NewDnsServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "dnsServiceId")

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
id := vmwares.NewDnsZoneID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "dnsZoneId")

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
id := vmwares.NewPortMirroringProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "portMirroringId")

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
id := vmwares.NewVMGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "vmGroupId")

payload := vmwares.WorkloadNetworkVMGroup{
	// ...
}


if err := client.WorkloadNetworksUpdateVMGroupThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
