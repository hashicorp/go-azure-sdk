package apimanagementservice

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessType string

const (
	AccessTypeAccessKey                     AccessType = "AccessKey"
	AccessTypeSystemAssignedManagedIdentity AccessType = "SystemAssignedManagedIdentity"
	AccessTypeUserAssignedManagedIdentity   AccessType = "UserAssignedManagedIdentity"
)

func PossibleValuesForAccessType() []string {
	return []string{
		string(AccessTypeAccessKey),
		string(AccessTypeSystemAssignedManagedIdentity),
		string(AccessTypeUserAssignedManagedIdentity),
	}
}

type CertificateSource string

const (
	CertificateSourceBuiltIn  CertificateSource = "BuiltIn"
	CertificateSourceCustom   CertificateSource = "Custom"
	CertificateSourceKeyVault CertificateSource = "KeyVault"
	CertificateSourceManaged  CertificateSource = "Managed"
)

func PossibleValuesForCertificateSource() []string {
	return []string{
		string(CertificateSourceBuiltIn),
		string(CertificateSourceCustom),
		string(CertificateSourceKeyVault),
		string(CertificateSourceManaged),
	}
}

type CertificateStatus string

const (
	CertificateStatusCompleted  CertificateStatus = "Completed"
	CertificateStatusFailed     CertificateStatus = "Failed"
	CertificateStatusInProgress CertificateStatus = "InProgress"
)

func PossibleValuesForCertificateStatus() []string {
	return []string{
		string(CertificateStatusCompleted),
		string(CertificateStatusFailed),
		string(CertificateStatusInProgress),
	}
}

type HostnameType string

const (
	HostnameTypeDeveloperPortal HostnameType = "DeveloperPortal"
	HostnameTypeManagement      HostnameType = "Management"
	HostnameTypePortal          HostnameType = "Portal"
	HostnameTypeProxy           HostnameType = "Proxy"
	HostnameTypeScm             HostnameType = "Scm"
)

func PossibleValuesForHostnameType() []string {
	return []string{
		string(HostnameTypeDeveloperPortal),
		string(HostnameTypeManagement),
		string(HostnameTypePortal),
		string(HostnameTypeProxy),
		string(HostnameTypeScm),
	}
}

type NameAvailabilityReason string

const (
	NameAvailabilityReasonAlreadyExists NameAvailabilityReason = "AlreadyExists"
	NameAvailabilityReasonInvalid       NameAvailabilityReason = "Invalid"
	NameAvailabilityReasonValid         NameAvailabilityReason = "Valid"
)

func PossibleValuesForNameAvailabilityReason() []string {
	return []string{
		string(NameAvailabilityReasonAlreadyExists),
		string(NameAvailabilityReasonInvalid),
		string(NameAvailabilityReasonValid),
	}
}

type PlatformVersion string

const (
	PlatformVersionMtvOne       PlatformVersion = "mtv1"
	PlatformVersionStvOne       PlatformVersion = "stv1"
	PlatformVersionStvTwo       PlatformVersion = "stv2"
	PlatformVersionUndetermined PlatformVersion = "undetermined"
)

func PossibleValuesForPlatformVersion() []string {
	return []string{
		string(PlatformVersionMtvOne),
		string(PlatformVersionStvOne),
		string(PlatformVersionStvTwo),
		string(PlatformVersionUndetermined),
	}
}

type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = "Approved"
	PrivateEndpointServiceConnectionStatusPending  PrivateEndpointServiceConnectionStatus = "Pending"
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = "Rejected"
)

func PossibleValuesForPrivateEndpointServiceConnectionStatus() []string {
	return []string{
		string(PrivateEndpointServiceConnectionStatusApproved),
		string(PrivateEndpointServiceConnectionStatusPending),
		string(PrivateEndpointServiceConnectionStatusRejected),
	}
}

type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

func PossibleValuesForPublicNetworkAccess() []string {
	return []string{
		string(PublicNetworkAccessDisabled),
		string(PublicNetworkAccessEnabled),
	}
}

type SkuType string

const (
	SkuTypeBasic       SkuType = "Basic"
	SkuTypeConsumption SkuType = "Consumption"
	SkuTypeDeveloper   SkuType = "Developer"
	SkuTypeIsolated    SkuType = "Isolated"
	SkuTypePremium     SkuType = "Premium"
	SkuTypeStandard    SkuType = "Standard"
)

func PossibleValuesForSkuType() []string {
	return []string{
		string(SkuTypeBasic),
		string(SkuTypeConsumption),
		string(SkuTypeDeveloper),
		string(SkuTypeIsolated),
		string(SkuTypePremium),
		string(SkuTypeStandard),
	}
}

type StoreName string

const (
	StoreNameCertificateAuthority StoreName = "CertificateAuthority"
	StoreNameRoot                 StoreName = "Root"
)

func PossibleValuesForStoreName() []string {
	return []string{
		string(StoreNameCertificateAuthority),
		string(StoreNameRoot),
	}
}

type VirtualNetworkType string

const (
	VirtualNetworkTypeExternal VirtualNetworkType = "External"
	VirtualNetworkTypeInternal VirtualNetworkType = "Internal"
	VirtualNetworkTypeNone     VirtualNetworkType = "None"
)

func PossibleValuesForVirtualNetworkType() []string {
	return []string{
		string(VirtualNetworkTypeExternal),
		string(VirtualNetworkTypeInternal),
		string(VirtualNetworkTypeNone),
	}
}
