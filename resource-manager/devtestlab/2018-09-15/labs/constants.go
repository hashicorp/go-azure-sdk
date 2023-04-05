package labs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnableStatus string

const (
	EnableStatusDisabled EnableStatus = "Disabled"
	EnableStatusEnabled  EnableStatus = "Enabled"
)

func PossibleValuesForEnableStatus() []string {
	return []string{
		string(EnableStatusDisabled),
		string(EnableStatusEnabled),
	}
}

type EnvironmentPermission string

const (
	EnvironmentPermissionContributor EnvironmentPermission = "Contributor"
	EnvironmentPermissionReader      EnvironmentPermission = "Reader"
)

func PossibleValuesForEnvironmentPermission() []string {
	return []string{
		string(EnvironmentPermissionContributor),
		string(EnvironmentPermissionReader),
	}
}

type HostCachingOptions string

const (
	HostCachingOptionsNone      HostCachingOptions = "None"
	HostCachingOptionsReadOnly  HostCachingOptions = "ReadOnly"
	HostCachingOptionsReadWrite HostCachingOptions = "ReadWrite"
)

func PossibleValuesForHostCachingOptions() []string {
	return []string{
		string(HostCachingOptionsNone),
		string(HostCachingOptionsReadOnly),
		string(HostCachingOptionsReadWrite),
	}
}

type PremiumDataDisk string

const (
	PremiumDataDiskDisabled PremiumDataDisk = "Disabled"
	PremiumDataDiskEnabled  PremiumDataDisk = "Enabled"
)

func PossibleValuesForPremiumDataDisk() []string {
	return []string{
		string(PremiumDataDiskDisabled),
		string(PremiumDataDiskEnabled),
	}
}

type StorageType string

const (
	StorageTypePremium     StorageType = "Premium"
	StorageTypeStandard    StorageType = "Standard"
	StorageTypeStandardSSD StorageType = "StandardSSD"
)

func PossibleValuesForStorageType() []string {
	return []string{
		string(StorageTypePremium),
		string(StorageTypeStandard),
		string(StorageTypeStandardSSD),
	}
}

type TransportProtocol string

const (
	TransportProtocolTcp TransportProtocol = "Tcp"
	TransportProtocolUdp TransportProtocol = "Udp"
)

func PossibleValuesForTransportProtocol() []string {
	return []string{
		string(TransportProtocolTcp),
		string(TransportProtocolUdp),
	}
}
