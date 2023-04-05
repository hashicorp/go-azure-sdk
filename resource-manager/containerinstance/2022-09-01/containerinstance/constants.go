package containerinstance

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerGroupIPAddressType string

const (
	ContainerGroupIPAddressTypePrivate ContainerGroupIPAddressType = "Private"
	ContainerGroupIPAddressTypePublic  ContainerGroupIPAddressType = "Public"
)

func PossibleValuesForContainerGroupIPAddressType() []string {
	return []string{
		string(ContainerGroupIPAddressTypePrivate),
		string(ContainerGroupIPAddressTypePublic),
	}
}

type ContainerGroupNetworkProtocol string

const (
	ContainerGroupNetworkProtocolTCP ContainerGroupNetworkProtocol = "TCP"
	ContainerGroupNetworkProtocolUDP ContainerGroupNetworkProtocol = "UDP"
)

func PossibleValuesForContainerGroupNetworkProtocol() []string {
	return []string{
		string(ContainerGroupNetworkProtocolTCP),
		string(ContainerGroupNetworkProtocolUDP),
	}
}

type ContainerGroupRestartPolicy string

const (
	ContainerGroupRestartPolicyAlways    ContainerGroupRestartPolicy = "Always"
	ContainerGroupRestartPolicyNever     ContainerGroupRestartPolicy = "Never"
	ContainerGroupRestartPolicyOnFailure ContainerGroupRestartPolicy = "OnFailure"
)

func PossibleValuesForContainerGroupRestartPolicy() []string {
	return []string{
		string(ContainerGroupRestartPolicyAlways),
		string(ContainerGroupRestartPolicyNever),
		string(ContainerGroupRestartPolicyOnFailure),
	}
}

type ContainerGroupSku string

const (
	ContainerGroupSkuDedicated ContainerGroupSku = "Dedicated"
	ContainerGroupSkuStandard  ContainerGroupSku = "Standard"
)

func PossibleValuesForContainerGroupSku() []string {
	return []string{
		string(ContainerGroupSkuDedicated),
		string(ContainerGroupSkuStandard),
	}
}

type ContainerNetworkProtocol string

const (
	ContainerNetworkProtocolTCP ContainerNetworkProtocol = "TCP"
	ContainerNetworkProtocolUDP ContainerNetworkProtocol = "UDP"
)

func PossibleValuesForContainerNetworkProtocol() []string {
	return []string{
		string(ContainerNetworkProtocolTCP),
		string(ContainerNetworkProtocolUDP),
	}
}

type DnsNameLabelReusePolicy string

const (
	DnsNameLabelReusePolicyNoreuse            DnsNameLabelReusePolicy = "Noreuse"
	DnsNameLabelReusePolicyResourceGroupReuse DnsNameLabelReusePolicy = "ResourceGroupReuse"
	DnsNameLabelReusePolicySubscriptionReuse  DnsNameLabelReusePolicy = "SubscriptionReuse"
	DnsNameLabelReusePolicyTenantReuse        DnsNameLabelReusePolicy = "TenantReuse"
	DnsNameLabelReusePolicyUnsecure           DnsNameLabelReusePolicy = "Unsecure"
)

func PossibleValuesForDnsNameLabelReusePolicy() []string {
	return []string{
		string(DnsNameLabelReusePolicyNoreuse),
		string(DnsNameLabelReusePolicyResourceGroupReuse),
		string(DnsNameLabelReusePolicySubscriptionReuse),
		string(DnsNameLabelReusePolicyTenantReuse),
		string(DnsNameLabelReusePolicyUnsecure),
	}
}

type GpuSku string

const (
	GpuSkuKEightZero  GpuSku = "K80"
	GpuSkuPOneHundred GpuSku = "P100"
	GpuSkuVOneHundred GpuSku = "V100"
)

func PossibleValuesForGpuSku() []string {
	return []string{
		string(GpuSkuKEightZero),
		string(GpuSkuPOneHundred),
		string(GpuSkuVOneHundred),
	}
}

type LogAnalyticsLogType string

const (
	LogAnalyticsLogTypeContainerInsights     LogAnalyticsLogType = "ContainerInsights"
	LogAnalyticsLogTypeContainerInstanceLogs LogAnalyticsLogType = "ContainerInstanceLogs"
)

func PossibleValuesForLogAnalyticsLogType() []string {
	return []string{
		string(LogAnalyticsLogTypeContainerInsights),
		string(LogAnalyticsLogTypeContainerInstanceLogs),
	}
}

type OperatingSystemTypes string

const (
	OperatingSystemTypesLinux   OperatingSystemTypes = "Linux"
	OperatingSystemTypesWindows OperatingSystemTypes = "Windows"
)

func PossibleValuesForOperatingSystemTypes() []string {
	return []string{
		string(OperatingSystemTypesLinux),
		string(OperatingSystemTypesWindows),
	}
}

type Scheme string

const (
	SchemeHTTP  Scheme = "http"
	SchemeHTTPS Scheme = "https"
)

func PossibleValuesForScheme() []string {
	return []string{
		string(SchemeHTTP),
		string(SchemeHTTPS),
	}
}
