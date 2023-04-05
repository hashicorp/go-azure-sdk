package storageaccounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessTier string

const (
	AccessTierCool    AccessTier = "Cool"
	AccessTierHot     AccessTier = "Hot"
	AccessTierPremium AccessTier = "Premium"
)

func PossibleValuesForAccessTier() []string {
	return []string{
		string(AccessTierCool),
		string(AccessTierHot),
		string(AccessTierPremium),
	}
}

type AccountImmutabilityPolicyState string

const (
	AccountImmutabilityPolicyStateDisabled AccountImmutabilityPolicyState = "Disabled"
	AccountImmutabilityPolicyStateLocked   AccountImmutabilityPolicyState = "Locked"
	AccountImmutabilityPolicyStateUnlocked AccountImmutabilityPolicyState = "Unlocked"
)

func PossibleValuesForAccountImmutabilityPolicyState() []string {
	return []string{
		string(AccountImmutabilityPolicyStateDisabled),
		string(AccountImmutabilityPolicyStateLocked),
		string(AccountImmutabilityPolicyStateUnlocked),
	}
}

type AccountStatus string

const (
	AccountStatusAvailable   AccountStatus = "available"
	AccountStatusUnavailable AccountStatus = "unavailable"
)

func PossibleValuesForAccountStatus() []string {
	return []string{
		string(AccountStatusAvailable),
		string(AccountStatusUnavailable),
	}
}

type AccountType string

const (
	AccountTypeComputer AccountType = "Computer"
	AccountTypeUser     AccountType = "User"
)

func PossibleValuesForAccountType() []string {
	return []string{
		string(AccountTypeComputer),
		string(AccountTypeUser),
	}
}

type Action string

const (
	ActionAllow Action = "Allow"
)

func PossibleValuesForAction() []string {
	return []string{
		string(ActionAllow),
	}
}

type AllowedCopyScope string

const (
	AllowedCopyScopeAAD         AllowedCopyScope = "AAD"
	AllowedCopyScopePrivateLink AllowedCopyScope = "PrivateLink"
)

func PossibleValuesForAllowedCopyScope() []string {
	return []string{
		string(AllowedCopyScopeAAD),
		string(AllowedCopyScopePrivateLink),
	}
}

type BlobRestoreProgressStatus string

const (
	BlobRestoreProgressStatusComplete   BlobRestoreProgressStatus = "Complete"
	BlobRestoreProgressStatusFailed     BlobRestoreProgressStatus = "Failed"
	BlobRestoreProgressStatusInProgress BlobRestoreProgressStatus = "InProgress"
)

func PossibleValuesForBlobRestoreProgressStatus() []string {
	return []string{
		string(BlobRestoreProgressStatusComplete),
		string(BlobRestoreProgressStatusFailed),
		string(BlobRestoreProgressStatusInProgress),
	}
}

type Bypass string

const (
	BypassAzureServices Bypass = "AzureServices"
	BypassLogging       Bypass = "Logging"
	BypassMetrics       Bypass = "Metrics"
	BypassNone          Bypass = "None"
)

func PossibleValuesForBypass() []string {
	return []string{
		string(BypassAzureServices),
		string(BypassLogging),
		string(BypassMetrics),
		string(BypassNone),
	}
}

type DefaultAction string

const (
	DefaultActionAllow DefaultAction = "Allow"
	DefaultActionDeny  DefaultAction = "Deny"
)

func PossibleValuesForDefaultAction() []string {
	return []string{
		string(DefaultActionAllow),
		string(DefaultActionDeny),
	}
}

type DefaultSharePermission string

const (
	DefaultSharePermissionNone                                       DefaultSharePermission = "None"
	DefaultSharePermissionStorageFileDataSmbShareContributor         DefaultSharePermission = "StorageFileDataSmbShareContributor"
	DefaultSharePermissionStorageFileDataSmbShareElevatedContributor DefaultSharePermission = "StorageFileDataSmbShareElevatedContributor"
	DefaultSharePermissionStorageFileDataSmbShareReader              DefaultSharePermission = "StorageFileDataSmbShareReader"
)

func PossibleValuesForDefaultSharePermission() []string {
	return []string{
		string(DefaultSharePermissionNone),
		string(DefaultSharePermissionStorageFileDataSmbShareContributor),
		string(DefaultSharePermissionStorageFileDataSmbShareElevatedContributor),
		string(DefaultSharePermissionStorageFileDataSmbShareReader),
	}
}

type DirectoryServiceOptions string

const (
	DirectoryServiceOptionsAADDS   DirectoryServiceOptions = "AADDS"
	DirectoryServiceOptionsAADKERB DirectoryServiceOptions = "AADKERB"
	DirectoryServiceOptionsAD      DirectoryServiceOptions = "AD"
	DirectoryServiceOptionsNone    DirectoryServiceOptions = "None"
)

func PossibleValuesForDirectoryServiceOptions() []string {
	return []string{
		string(DirectoryServiceOptionsAADDS),
		string(DirectoryServiceOptionsAADKERB),
		string(DirectoryServiceOptionsAD),
		string(DirectoryServiceOptionsNone),
	}
}

type DnsEndpointType string

const (
	DnsEndpointTypeAzureDnsZone DnsEndpointType = "AzureDnsZone"
	DnsEndpointTypeStandard     DnsEndpointType = "Standard"
)

func PossibleValuesForDnsEndpointType() []string {
	return []string{
		string(DnsEndpointTypeAzureDnsZone),
		string(DnsEndpointTypeStandard),
	}
}

type ExpirationAction string

const (
	ExpirationActionLog ExpirationAction = "Log"
)

func PossibleValuesForExpirationAction() []string {
	return []string{
		string(ExpirationActionLog),
	}
}

type GeoReplicationStatus string

const (
	GeoReplicationStatusBootstrap   GeoReplicationStatus = "Bootstrap"
	GeoReplicationStatusLive        GeoReplicationStatus = "Live"
	GeoReplicationStatusUnavailable GeoReplicationStatus = "Unavailable"
)

func PossibleValuesForGeoReplicationStatus() []string {
	return []string{
		string(GeoReplicationStatusBootstrap),
		string(GeoReplicationStatusLive),
		string(GeoReplicationStatusUnavailable),
	}
}

type HTTPProtocol string

const (
	HTTPProtocolHTTPS     HTTPProtocol = "https"
	HTTPProtocolHTTPSHttp HTTPProtocol = "https,http"
)

func PossibleValuesForHTTPProtocol() []string {
	return []string{
		string(HTTPProtocolHTTPS),
		string(HTTPProtocolHTTPSHttp),
	}
}

type KeyPermission string

const (
	KeyPermissionFull KeyPermission = "Full"
	KeyPermissionRead KeyPermission = "Read"
)

func PossibleValuesForKeyPermission() []string {
	return []string{
		string(KeyPermissionFull),
		string(KeyPermissionRead),
	}
}

type KeySource string

const (
	KeySourceMicrosoftPointKeyvault KeySource = "Microsoft.Keyvault"
	KeySourceMicrosoftPointStorage  KeySource = "Microsoft.Storage"
)

func PossibleValuesForKeySource() []string {
	return []string{
		string(KeySourceMicrosoftPointKeyvault),
		string(KeySourceMicrosoftPointStorage),
	}
}

type KeyType string

const (
	KeyTypeAccount KeyType = "Account"
	KeyTypeService KeyType = "Service"
)

func PossibleValuesForKeyType() []string {
	return []string{
		string(KeyTypeAccount),
		string(KeyTypeService),
	}
}

type Kind string

const (
	KindBlobStorage      Kind = "BlobStorage"
	KindBlockBlobStorage Kind = "BlockBlobStorage"
	KindFileStorage      Kind = "FileStorage"
	KindStorage          Kind = "Storage"
	KindStorageVTwo      Kind = "StorageV2"
)

func PossibleValuesForKind() []string {
	return []string{
		string(KindBlobStorage),
		string(KindBlockBlobStorage),
		string(KindFileStorage),
		string(KindStorage),
		string(KindStorageVTwo),
	}
}

type LargeFileSharesState string

const (
	LargeFileSharesStateDisabled LargeFileSharesState = "Disabled"
	LargeFileSharesStateEnabled  LargeFileSharesState = "Enabled"
)

func PossibleValuesForLargeFileSharesState() []string {
	return []string{
		string(LargeFileSharesStateDisabled),
		string(LargeFileSharesStateEnabled),
	}
}

type ListKeyExpand string

const (
	ListKeyExpandKerb ListKeyExpand = "kerb"
)

func PossibleValuesForListKeyExpand() []string {
	return []string{
		string(ListKeyExpandKerb),
	}
}

type MinimumTlsVersion string

const (
	MinimumTlsVersionTLSOneOne  MinimumTlsVersion = "TLS1_1"
	MinimumTlsVersionTLSOneTwo  MinimumTlsVersion = "TLS1_2"
	MinimumTlsVersionTLSOneZero MinimumTlsVersion = "TLS1_0"
)

func PossibleValuesForMinimumTlsVersion() []string {
	return []string{
		string(MinimumTlsVersionTLSOneOne),
		string(MinimumTlsVersionTLSOneTwo),
		string(MinimumTlsVersionTLSOneZero),
	}
}

type Permissions string

const (
	PermissionsA Permissions = "a"
	PermissionsC Permissions = "c"
	PermissionsD Permissions = "d"
	PermissionsL Permissions = "l"
	PermissionsP Permissions = "p"
	PermissionsR Permissions = "r"
	PermissionsU Permissions = "u"
	PermissionsW Permissions = "w"
)

func PossibleValuesForPermissions() []string {
	return []string{
		string(PermissionsA),
		string(PermissionsC),
		string(PermissionsD),
		string(PermissionsL),
		string(PermissionsP),
		string(PermissionsR),
		string(PermissionsU),
		string(PermissionsW),
	}
}

type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = "Creating"
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
)

func PossibleValuesForPrivateEndpointConnectionProvisioningState() []string {
	return []string{
		string(PrivateEndpointConnectionProvisioningStateCreating),
		string(PrivateEndpointConnectionProvisioningStateDeleting),
		string(PrivateEndpointConnectionProvisioningStateFailed),
		string(PrivateEndpointConnectionProvisioningStateSucceeded),
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

type ProvisioningState string

const (
	ProvisioningStateCreating     ProvisioningState = "Creating"
	ProvisioningStateResolvingDNS ProvisioningState = "ResolvingDNS"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCreating),
		string(ProvisioningStateResolvingDNS),
		string(ProvisioningStateSucceeded),
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

type Reason string

const (
	ReasonAccountNameInvalid Reason = "AccountNameInvalid"
	ReasonAlreadyExists      Reason = "AlreadyExists"
)

func PossibleValuesForReason() []string {
	return []string{
		string(ReasonAccountNameInvalid),
		string(ReasonAlreadyExists),
	}
}

type RoutingChoice string

const (
	RoutingChoiceInternetRouting  RoutingChoice = "InternetRouting"
	RoutingChoiceMicrosoftRouting RoutingChoice = "MicrosoftRouting"
)

func PossibleValuesForRoutingChoice() []string {
	return []string{
		string(RoutingChoiceInternetRouting),
		string(RoutingChoiceMicrosoftRouting),
	}
}

type Services string

const (
	ServicesB Services = "b"
	ServicesF Services = "f"
	ServicesQ Services = "q"
	ServicesT Services = "t"
)

func PossibleValuesForServices() []string {
	return []string{
		string(ServicesB),
		string(ServicesF),
		string(ServicesQ),
		string(ServicesT),
	}
}

type SignedResource string

const (
	SignedResourceB SignedResource = "b"
	SignedResourceC SignedResource = "c"
	SignedResourceF SignedResource = "f"
	SignedResourceS SignedResource = "s"
)

func PossibleValuesForSignedResource() []string {
	return []string{
		string(SignedResourceB),
		string(SignedResourceC),
		string(SignedResourceF),
		string(SignedResourceS),
	}
}

type SignedResourceTypes string

const (
	SignedResourceTypesC SignedResourceTypes = "c"
	SignedResourceTypesO SignedResourceTypes = "o"
	SignedResourceTypesS SignedResourceTypes = "s"
)

func PossibleValuesForSignedResourceTypes() []string {
	return []string{
		string(SignedResourceTypesC),
		string(SignedResourceTypesO),
		string(SignedResourceTypesS),
	}
}

type SkuConversionStatus string

const (
	SkuConversionStatusFailed     SkuConversionStatus = "Failed"
	SkuConversionStatusInProgress SkuConversionStatus = "InProgress"
	SkuConversionStatusSucceeded  SkuConversionStatus = "Succeeded"
)

func PossibleValuesForSkuConversionStatus() []string {
	return []string{
		string(SkuConversionStatusFailed),
		string(SkuConversionStatusInProgress),
		string(SkuConversionStatusSucceeded),
	}
}

type SkuName string

const (
	SkuNamePremiumLRS     SkuName = "Premium_LRS"
	SkuNamePremiumZRS     SkuName = "Premium_ZRS"
	SkuNameStandardGRS    SkuName = "Standard_GRS"
	SkuNameStandardGZRS   SkuName = "Standard_GZRS"
	SkuNameStandardLRS    SkuName = "Standard_LRS"
	SkuNameStandardRAGRS  SkuName = "Standard_RAGRS"
	SkuNameStandardRAGZRS SkuName = "Standard_RAGZRS"
	SkuNameStandardZRS    SkuName = "Standard_ZRS"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNamePremiumLRS),
		string(SkuNamePremiumZRS),
		string(SkuNameStandardGRS),
		string(SkuNameStandardGZRS),
		string(SkuNameStandardLRS),
		string(SkuNameStandardRAGRS),
		string(SkuNameStandardRAGZRS),
		string(SkuNameStandardZRS),
	}
}

type SkuTier string

const (
	SkuTierPremium  SkuTier = "Premium"
	SkuTierStandard SkuTier = "Standard"
)

func PossibleValuesForSkuTier() []string {
	return []string{
		string(SkuTierPremium),
		string(SkuTierStandard),
	}
}

type State string

const (
	StateDeprovisioning       State = "Deprovisioning"
	StateFailed               State = "Failed"
	StateNetworkSourceDeleted State = "NetworkSourceDeleted"
	StateProvisioning         State = "Provisioning"
	StateSucceeded            State = "Succeeded"
)

func PossibleValuesForState() []string {
	return []string{
		string(StateDeprovisioning),
		string(StateFailed),
		string(StateNetworkSourceDeleted),
		string(StateProvisioning),
		string(StateSucceeded),
	}
}

type StorageAccountExpand string

const (
	StorageAccountExpandBlobRestoreStatus   StorageAccountExpand = "blobRestoreStatus"
	StorageAccountExpandGeoReplicationStats StorageAccountExpand = "geoReplicationStats"
)

func PossibleValuesForStorageAccountExpand() []string {
	return []string{
		string(StorageAccountExpandBlobRestoreStatus),
		string(StorageAccountExpandGeoReplicationStats),
	}
}

type Type string

const (
	TypeMicrosoftPointStorageStorageAccounts Type = "Microsoft.Storage/storageAccounts"
)

func PossibleValuesForType() []string {
	return []string{
		string(TypeMicrosoftPointStorageStorageAccounts),
	}
}
