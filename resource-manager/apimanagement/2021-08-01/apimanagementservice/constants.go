package apimanagementservice

import "strings"

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

func parseAccessType(input string) (*AccessType, error) {
	vals := map[string]AccessType{
		"accesskey":                     AccessTypeAccessKey,
		"systemassignedmanagedidentity": AccessTypeSystemAssignedManagedIdentity,
		"userassignedmanagedidentity":   AccessTypeUserAssignedManagedIdentity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessType(input)
	return &out, nil
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

func parseCertificateSource(input string) (*CertificateSource, error) {
	vals := map[string]CertificateSource{
		"builtin":  CertificateSourceBuiltIn,
		"custom":   CertificateSourceCustom,
		"keyvault": CertificateSourceKeyVault,
		"managed":  CertificateSourceManaged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CertificateSource(input)
	return &out, nil
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

func parseCertificateStatus(input string) (*CertificateStatus, error) {
	vals := map[string]CertificateStatus{
		"completed":  CertificateStatusCompleted,
		"failed":     CertificateStatusFailed,
		"inprogress": CertificateStatusInProgress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CertificateStatus(input)
	return &out, nil
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

func parseHostnameType(input string) (*HostnameType, error) {
	vals := map[string]HostnameType{
		"developerportal": HostnameTypeDeveloperPortal,
		"management":      HostnameTypeManagement,
		"portal":          HostnameTypePortal,
		"proxy":           HostnameTypeProxy,
		"scm":             HostnameTypeScm,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HostnameType(input)
	return &out, nil
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

func parseNameAvailabilityReason(input string) (*NameAvailabilityReason, error) {
	vals := map[string]NameAvailabilityReason{
		"alreadyexists": NameAvailabilityReasonAlreadyExists,
		"invalid":       NameAvailabilityReasonInvalid,
		"valid":         NameAvailabilityReasonValid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NameAvailabilityReason(input)
	return &out, nil
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

func parsePlatformVersion(input string) (*PlatformVersion, error) {
	vals := map[string]PlatformVersion{
		"mtv1":         PlatformVersionMtvOne,
		"stv1":         PlatformVersionStvOne,
		"stv2":         PlatformVersionStvTwo,
		"undetermined": PlatformVersionUndetermined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlatformVersion(input)
	return &out, nil
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

func parsePrivateEndpointServiceConnectionStatus(input string) (*PrivateEndpointServiceConnectionStatus, error) {
	vals := map[string]PrivateEndpointServiceConnectionStatus{
		"approved": PrivateEndpointServiceConnectionStatusApproved,
		"pending":  PrivateEndpointServiceConnectionStatusPending,
		"rejected": PrivateEndpointServiceConnectionStatusRejected,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivateEndpointServiceConnectionStatus(input)
	return &out, nil
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

func parsePublicNetworkAccess(input string) (*PublicNetworkAccess, error) {
	vals := map[string]PublicNetworkAccess{
		"disabled": PublicNetworkAccessDisabled,
		"enabled":  PublicNetworkAccessEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PublicNetworkAccess(input)
	return &out, nil
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

func parseSkuType(input string) (*SkuType, error) {
	vals := map[string]SkuType{
		"basic":       SkuTypeBasic,
		"consumption": SkuTypeConsumption,
		"developer":   SkuTypeDeveloper,
		"isolated":    SkuTypeIsolated,
		"premium":     SkuTypePremium,
		"standard":    SkuTypeStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuType(input)
	return &out, nil
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

func parseStoreName(input string) (*StoreName, error) {
	vals := map[string]StoreName{
		"certificateauthority": StoreNameCertificateAuthority,
		"root":                 StoreNameRoot,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StoreName(input)
	return &out, nil
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

func parseVirtualNetworkType(input string) (*VirtualNetworkType, error) {
	vals := map[string]VirtualNetworkType{
		"external": VirtualNetworkTypeExternal,
		"internal": VirtualNetworkTypeInternal,
		"none":     VirtualNetworkTypeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualNetworkType(input)
	return &out, nil
}
