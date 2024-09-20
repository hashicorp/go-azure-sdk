
## `github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-03-01/workloadnetworks` Documentation

The `workloadnetworks` SDK allows for interaction with Azure Resource Manager `vmware` (API Version `2023-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-03-01/workloadnetworks"
```


### Client Initialization

```go
client := workloadnetworks.NewWorkloadNetworksClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkloadNetworksClient.CreateDhcp`

```go
ctx := context.TODO()
id := workloadnetworks.NewDhcpConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "dhcpId")

payload := workloadnetworks.WorkloadNetworkDhcp{
	// ...
}


if err := client.CreateDhcpThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.CreateDnsService`

```go
ctx := context.TODO()
id := workloadnetworks.NewDnsServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "dnsServiceId")

payload := workloadnetworks.WorkloadNetworkDnsService{
	// ...
}


if err := client.CreateDnsServiceThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.CreateDnsZone`

```go
ctx := context.TODO()
id := workloadnetworks.NewDnsZoneID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "dnsZoneId")

payload := workloadnetworks.WorkloadNetworkDnsZone{
	// ...
}


if err := client.CreateDnsZoneThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.CreatePortMirroring`

```go
ctx := context.TODO()
id := workloadnetworks.NewPortMirroringProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "portMirroringId")

payload := workloadnetworks.WorkloadNetworkPortMirroring{
	// ...
}


if err := client.CreatePortMirroringThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.CreatePublicIP`

```go
ctx := context.TODO()
id := workloadnetworks.NewPublicIPID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "publicIPId")

payload := workloadnetworks.WorkloadNetworkPublicIP{
	// ...
}


if err := client.CreatePublicIPThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.CreateSegments`

```go
ctx := context.TODO()
id := workloadnetworks.NewSegmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "segmentId")

payload := workloadnetworks.WorkloadNetworkSegment{
	// ...
}


if err := client.CreateSegmentsThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.CreateVMGroup`

```go
ctx := context.TODO()
id := workloadnetworks.NewVMGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "vmGroupId")

payload := workloadnetworks.WorkloadNetworkVMGroup{
	// ...
}


if err := client.CreateVMGroupThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.DeleteDhcp`

```go
ctx := context.TODO()
id := workloadnetworks.NewDhcpConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "dhcpId")

if err := client.DeleteDhcpThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.DeleteDnsService`

```go
ctx := context.TODO()
id := workloadnetworks.NewDnsServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "dnsServiceId")

if err := client.DeleteDnsServiceThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.DeleteDnsZone`

```go
ctx := context.TODO()
id := workloadnetworks.NewDnsZoneID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "dnsZoneId")

if err := client.DeleteDnsZoneThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.DeletePortMirroring`

```go
ctx := context.TODO()
id := workloadnetworks.NewPortMirroringProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "portMirroringId")

if err := client.DeletePortMirroringThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.DeletePublicIP`

```go
ctx := context.TODO()
id := workloadnetworks.NewPublicIPID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "publicIPId")

if err := client.DeletePublicIPThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.DeleteSegment`

```go
ctx := context.TODO()
id := workloadnetworks.NewSegmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "segmentId")

if err := client.DeleteSegmentThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.DeleteVMGroup`

```go
ctx := context.TODO()
id := workloadnetworks.NewVMGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "vmGroupId")

if err := client.DeleteVMGroupThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.Get`

```go
ctx := context.TODO()
id := workloadnetworks.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkloadNetworksClient.GetDhcp`

```go
ctx := context.TODO()
id := workloadnetworks.NewDhcpConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "dhcpId")

read, err := client.GetDhcp(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkloadNetworksClient.GetDnsService`

```go
ctx := context.TODO()
id := workloadnetworks.NewDnsServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "dnsServiceId")

read, err := client.GetDnsService(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkloadNetworksClient.GetDnsZone`

```go
ctx := context.TODO()
id := workloadnetworks.NewDnsZoneID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "dnsZoneId")

read, err := client.GetDnsZone(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkloadNetworksClient.GetGateway`

```go
ctx := context.TODO()
id := workloadnetworks.NewGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "gatewayId")

read, err := client.GetGateway(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkloadNetworksClient.GetPortMirroring`

```go
ctx := context.TODO()
id := workloadnetworks.NewPortMirroringProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "portMirroringId")

read, err := client.GetPortMirroring(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkloadNetworksClient.GetPublicIP`

```go
ctx := context.TODO()
id := workloadnetworks.NewPublicIPID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "publicIPId")

read, err := client.GetPublicIP(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkloadNetworksClient.GetSegment`

```go
ctx := context.TODO()
id := workloadnetworks.NewSegmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "segmentId")

read, err := client.GetSegment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkloadNetworksClient.GetVMGroup`

```go
ctx := context.TODO()
id := workloadnetworks.NewVMGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "vmGroupId")

read, err := client.GetVMGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkloadNetworksClient.GetVirtualMachine`

```go
ctx := context.TODO()
id := workloadnetworks.NewDefaultVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "virtualMachineId")

read, err := client.GetVirtualMachine(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkloadNetworksClient.List`

```go
ctx := context.TODO()
id := workloadnetworks.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkloadNetworksClient.ListDhcp`

```go
ctx := context.TODO()
id := workloadnetworks.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName")

// alternatively `client.ListDhcp(ctx, id)` can be used to do batched pagination
items, err := client.ListDhcpComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkloadNetworksClient.ListDnsServices`

```go
ctx := context.TODO()
id := workloadnetworks.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName")

// alternatively `client.ListDnsServices(ctx, id)` can be used to do batched pagination
items, err := client.ListDnsServicesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkloadNetworksClient.ListDnsZones`

```go
ctx := context.TODO()
id := workloadnetworks.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName")

// alternatively `client.ListDnsZones(ctx, id)` can be used to do batched pagination
items, err := client.ListDnsZonesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkloadNetworksClient.ListGateways`

```go
ctx := context.TODO()
id := workloadnetworks.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName")

// alternatively `client.ListGateways(ctx, id)` can be used to do batched pagination
items, err := client.ListGatewaysComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkloadNetworksClient.ListPortMirroring`

```go
ctx := context.TODO()
id := workloadnetworks.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName")

// alternatively `client.ListPortMirroring(ctx, id)` can be used to do batched pagination
items, err := client.ListPortMirroringComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkloadNetworksClient.ListPublicIPs`

```go
ctx := context.TODO()
id := workloadnetworks.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName")

// alternatively `client.ListPublicIPs(ctx, id)` can be used to do batched pagination
items, err := client.ListPublicIPsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkloadNetworksClient.ListSegments`

```go
ctx := context.TODO()
id := workloadnetworks.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName")

// alternatively `client.ListSegments(ctx, id)` can be used to do batched pagination
items, err := client.ListSegmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkloadNetworksClient.ListVMGroups`

```go
ctx := context.TODO()
id := workloadnetworks.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName")

// alternatively `client.ListVMGroups(ctx, id)` can be used to do batched pagination
items, err := client.ListVMGroupsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkloadNetworksClient.ListVirtualMachines`

```go
ctx := context.TODO()
id := workloadnetworks.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName")

// alternatively `client.ListVirtualMachines(ctx, id)` can be used to do batched pagination
items, err := client.ListVirtualMachinesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkloadNetworksClient.UpdateDhcp`

```go
ctx := context.TODO()
id := workloadnetworks.NewDhcpConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "dhcpId")

payload := workloadnetworks.WorkloadNetworkDhcp{
	// ...
}


if err := client.UpdateDhcpThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.UpdateDnsService`

```go
ctx := context.TODO()
id := workloadnetworks.NewDnsServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "dnsServiceId")

payload := workloadnetworks.WorkloadNetworkDnsService{
	// ...
}


if err := client.UpdateDnsServiceThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.UpdateDnsZone`

```go
ctx := context.TODO()
id := workloadnetworks.NewDnsZoneID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "dnsZoneId")

payload := workloadnetworks.WorkloadNetworkDnsZone{
	// ...
}


if err := client.UpdateDnsZoneThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.UpdatePortMirroring`

```go
ctx := context.TODO()
id := workloadnetworks.NewPortMirroringProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "portMirroringId")

payload := workloadnetworks.WorkloadNetworkPortMirroring{
	// ...
}


if err := client.UpdatePortMirroringThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.UpdateSegments`

```go
ctx := context.TODO()
id := workloadnetworks.NewSegmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "segmentId")

payload := workloadnetworks.WorkloadNetworkSegment{
	// ...
}


if err := client.UpdateSegmentsThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.UpdateVMGroup`

```go
ctx := context.TODO()
id := workloadnetworks.NewVMGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "vmGroupId")

payload := workloadnetworks.WorkloadNetworkVMGroup{
	// ...
}


if err := client.UpdateVMGroupThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
