package virtualmachines

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

type VirtualMachineCreationSource string

const (
	VirtualMachineCreationSourceFromCustomImage        VirtualMachineCreationSource = "FromCustomImage"
	VirtualMachineCreationSourceFromGalleryImage       VirtualMachineCreationSource = "FromGalleryImage"
	VirtualMachineCreationSourceFromSharedGalleryImage VirtualMachineCreationSource = "FromSharedGalleryImage"
)

func PossibleValuesForVirtualMachineCreationSource() []string {
	return []string{
		string(VirtualMachineCreationSourceFromCustomImage),
		string(VirtualMachineCreationSourceFromGalleryImage),
		string(VirtualMachineCreationSourceFromSharedGalleryImage),
	}
}
