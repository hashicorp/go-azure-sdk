package vaults

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPolicyUpdateKind string

const (
	AccessPolicyUpdateKindAdd     AccessPolicyUpdateKind = "add"
	AccessPolicyUpdateKindRemove  AccessPolicyUpdateKind = "remove"
	AccessPolicyUpdateKindReplace AccessPolicyUpdateKind = "replace"
)

func PossibleValuesForAccessPolicyUpdateKind() []string {
	return []string{
		string(AccessPolicyUpdateKindAdd),
		string(AccessPolicyUpdateKindRemove),
		string(AccessPolicyUpdateKindReplace),
	}
}

type ActionsRequired string

const (
	ActionsRequiredNone ActionsRequired = "None"
)

func PossibleValuesForActionsRequired() []string {
	return []string{
		string(ActionsRequiredNone),
	}
}

type CertificatePermissions string

const (
	CertificatePermissionsAll            CertificatePermissions = "all"
	CertificatePermissionsBackup         CertificatePermissions = "backup"
	CertificatePermissionsCreate         CertificatePermissions = "create"
	CertificatePermissionsDelete         CertificatePermissions = "delete"
	CertificatePermissionsDeleteissuers  CertificatePermissions = "deleteissuers"
	CertificatePermissionsGet            CertificatePermissions = "get"
	CertificatePermissionsGetissuers     CertificatePermissions = "getissuers"
	CertificatePermissionsImport         CertificatePermissions = "import"
	CertificatePermissionsList           CertificatePermissions = "list"
	CertificatePermissionsListissuers    CertificatePermissions = "listissuers"
	CertificatePermissionsManagecontacts CertificatePermissions = "managecontacts"
	CertificatePermissionsManageissuers  CertificatePermissions = "manageissuers"
	CertificatePermissionsPurge          CertificatePermissions = "purge"
	CertificatePermissionsRecover        CertificatePermissions = "recover"
	CertificatePermissionsRestore        CertificatePermissions = "restore"
	CertificatePermissionsSetissuers     CertificatePermissions = "setissuers"
	CertificatePermissionsUpdate         CertificatePermissions = "update"
)

func PossibleValuesForCertificatePermissions() []string {
	return []string{
		string(CertificatePermissionsAll),
		string(CertificatePermissionsBackup),
		string(CertificatePermissionsCreate),
		string(CertificatePermissionsDelete),
		string(CertificatePermissionsDeleteissuers),
		string(CertificatePermissionsGet),
		string(CertificatePermissionsGetissuers),
		string(CertificatePermissionsImport),
		string(CertificatePermissionsList),
		string(CertificatePermissionsListissuers),
		string(CertificatePermissionsManagecontacts),
		string(CertificatePermissionsManageissuers),
		string(CertificatePermissionsPurge),
		string(CertificatePermissionsRecover),
		string(CertificatePermissionsRestore),
		string(CertificatePermissionsSetissuers),
		string(CertificatePermissionsUpdate),
	}
}

type CreateMode string

const (
	CreateModeDefault CreateMode = "default"
	CreateModeRecover CreateMode = "recover"
)

func PossibleValuesForCreateMode() []string {
	return []string{
		string(CreateModeDefault),
		string(CreateModeRecover),
	}
}

type KeyPermissions string

const (
	KeyPermissionsAll               KeyPermissions = "all"
	KeyPermissionsBackup            KeyPermissions = "backup"
	KeyPermissionsCreate            KeyPermissions = "create"
	KeyPermissionsDecrypt           KeyPermissions = "decrypt"
	KeyPermissionsDelete            KeyPermissions = "delete"
	KeyPermissionsEncrypt           KeyPermissions = "encrypt"
	KeyPermissionsGet               KeyPermissions = "get"
	KeyPermissionsGetrotationpolicy KeyPermissions = "getrotationpolicy"
	KeyPermissionsImport            KeyPermissions = "import"
	KeyPermissionsList              KeyPermissions = "list"
	KeyPermissionsPurge             KeyPermissions = "purge"
	KeyPermissionsRecover           KeyPermissions = "recover"
	KeyPermissionsRelease           KeyPermissions = "release"
	KeyPermissionsRestore           KeyPermissions = "restore"
	KeyPermissionsRotate            KeyPermissions = "rotate"
	KeyPermissionsSetrotationpolicy KeyPermissions = "setrotationpolicy"
	KeyPermissionsSign              KeyPermissions = "sign"
	KeyPermissionsUnwrapKey         KeyPermissions = "unwrapKey"
	KeyPermissionsUpdate            KeyPermissions = "update"
	KeyPermissionsVerify            KeyPermissions = "verify"
	KeyPermissionsWrapKey           KeyPermissions = "wrapKey"
)

func PossibleValuesForKeyPermissions() []string {
	return []string{
		string(KeyPermissionsAll),
		string(KeyPermissionsBackup),
		string(KeyPermissionsCreate),
		string(KeyPermissionsDecrypt),
		string(KeyPermissionsDelete),
		string(KeyPermissionsEncrypt),
		string(KeyPermissionsGet),
		string(KeyPermissionsGetrotationpolicy),
		string(KeyPermissionsImport),
		string(KeyPermissionsList),
		string(KeyPermissionsPurge),
		string(KeyPermissionsRecover),
		string(KeyPermissionsRelease),
		string(KeyPermissionsRestore),
		string(KeyPermissionsRotate),
		string(KeyPermissionsSetrotationpolicy),
		string(KeyPermissionsSign),
		string(KeyPermissionsUnwrapKey),
		string(KeyPermissionsUpdate),
		string(KeyPermissionsVerify),
		string(KeyPermissionsWrapKey),
	}
}

type NetworkRuleAction string

const (
	NetworkRuleActionAllow NetworkRuleAction = "Allow"
	NetworkRuleActionDeny  NetworkRuleAction = "Deny"
)

func PossibleValuesForNetworkRuleAction() []string {
	return []string{
		string(NetworkRuleActionAllow),
		string(NetworkRuleActionDeny),
	}
}

type NetworkRuleBypassOptions string

const (
	NetworkRuleBypassOptionsAzureServices NetworkRuleBypassOptions = "AzureServices"
	NetworkRuleBypassOptionsNone          NetworkRuleBypassOptions = "None"
)

func PossibleValuesForNetworkRuleBypassOptions() []string {
	return []string{
		string(NetworkRuleBypassOptionsAzureServices),
		string(NetworkRuleBypassOptionsNone),
	}
}

type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCreating     PrivateEndpointConnectionProvisioningState = "Creating"
	PrivateEndpointConnectionProvisioningStateDeleting     PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateDisconnected PrivateEndpointConnectionProvisioningState = "Disconnected"
	PrivateEndpointConnectionProvisioningStateFailed       PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateSucceeded    PrivateEndpointConnectionProvisioningState = "Succeeded"
	PrivateEndpointConnectionProvisioningStateUpdating     PrivateEndpointConnectionProvisioningState = "Updating"
)

func PossibleValuesForPrivateEndpointConnectionProvisioningState() []string {
	return []string{
		string(PrivateEndpointConnectionProvisioningStateCreating),
		string(PrivateEndpointConnectionProvisioningStateDeleting),
		string(PrivateEndpointConnectionProvisioningStateDisconnected),
		string(PrivateEndpointConnectionProvisioningStateFailed),
		string(PrivateEndpointConnectionProvisioningStateSucceeded),
		string(PrivateEndpointConnectionProvisioningStateUpdating),
	}
}

type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusApproved     PrivateEndpointServiceConnectionStatus = "Approved"
	PrivateEndpointServiceConnectionStatusDisconnected PrivateEndpointServiceConnectionStatus = "Disconnected"
	PrivateEndpointServiceConnectionStatusPending      PrivateEndpointServiceConnectionStatus = "Pending"
	PrivateEndpointServiceConnectionStatusRejected     PrivateEndpointServiceConnectionStatus = "Rejected"
)

func PossibleValuesForPrivateEndpointServiceConnectionStatus() []string {
	return []string{
		string(PrivateEndpointServiceConnectionStatusApproved),
		string(PrivateEndpointServiceConnectionStatusDisconnected),
		string(PrivateEndpointServiceConnectionStatusPending),
		string(PrivateEndpointServiceConnectionStatusRejected),
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

type SecretPermissions string

const (
	SecretPermissionsAll     SecretPermissions = "all"
	SecretPermissionsBackup  SecretPermissions = "backup"
	SecretPermissionsDelete  SecretPermissions = "delete"
	SecretPermissionsGet     SecretPermissions = "get"
	SecretPermissionsList    SecretPermissions = "list"
	SecretPermissionsPurge   SecretPermissions = "purge"
	SecretPermissionsRecover SecretPermissions = "recover"
	SecretPermissionsRestore SecretPermissions = "restore"
	SecretPermissionsSet     SecretPermissions = "set"
)

func PossibleValuesForSecretPermissions() []string {
	return []string{
		string(SecretPermissionsAll),
		string(SecretPermissionsBackup),
		string(SecretPermissionsDelete),
		string(SecretPermissionsGet),
		string(SecretPermissionsList),
		string(SecretPermissionsPurge),
		string(SecretPermissionsRecover),
		string(SecretPermissionsRestore),
		string(SecretPermissionsSet),
	}
}

type SkuFamily string

const (
	SkuFamilyA SkuFamily = "A"
)

func PossibleValuesForSkuFamily() []string {
	return []string{
		string(SkuFamilyA),
	}
}

type SkuName string

const (
	SkuNamePremium  SkuName = "premium"
	SkuNameStandard SkuName = "standard"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNamePremium),
		string(SkuNameStandard),
	}
}

type StoragePermissions string

const (
	StoragePermissionsAll           StoragePermissions = "all"
	StoragePermissionsBackup        StoragePermissions = "backup"
	StoragePermissionsDelete        StoragePermissions = "delete"
	StoragePermissionsDeletesas     StoragePermissions = "deletesas"
	StoragePermissionsGet           StoragePermissions = "get"
	StoragePermissionsGetsas        StoragePermissions = "getsas"
	StoragePermissionsList          StoragePermissions = "list"
	StoragePermissionsListsas       StoragePermissions = "listsas"
	StoragePermissionsPurge         StoragePermissions = "purge"
	StoragePermissionsRecover       StoragePermissions = "recover"
	StoragePermissionsRegeneratekey StoragePermissions = "regeneratekey"
	StoragePermissionsRestore       StoragePermissions = "restore"
	StoragePermissionsSet           StoragePermissions = "set"
	StoragePermissionsSetsas        StoragePermissions = "setsas"
	StoragePermissionsUpdate        StoragePermissions = "update"
)

func PossibleValuesForStoragePermissions() []string {
	return []string{
		string(StoragePermissionsAll),
		string(StoragePermissionsBackup),
		string(StoragePermissionsDelete),
		string(StoragePermissionsDeletesas),
		string(StoragePermissionsGet),
		string(StoragePermissionsGetsas),
		string(StoragePermissionsList),
		string(StoragePermissionsListsas),
		string(StoragePermissionsPurge),
		string(StoragePermissionsRecover),
		string(StoragePermissionsRegeneratekey),
		string(StoragePermissionsRestore),
		string(StoragePermissionsSet),
		string(StoragePermissionsSetsas),
		string(StoragePermissionsUpdate),
	}
}

type Type string

const (
	TypeMicrosoftPointKeyVaultVaults Type = "Microsoft.KeyVault/vaults"
)

func PossibleValuesForType() []string {
	return []string{
		string(TypeMicrosoftPointKeyVaultVaults),
	}
}

type VaultListFilterTypes string

const (
	VaultListFilterTypesResourceTypeEqMicrosoftPointKeyVaultVaults VaultListFilterTypes = "resourceType eq 'Microsoft.KeyVault/vaults'"
)

func PossibleValuesForVaultListFilterTypes() []string {
	return []string{
		string(VaultListFilterTypesResourceTypeEqMicrosoftPointKeyVaultVaults),
	}
}

type VaultProvisioningState string

const (
	VaultProvisioningStateRegisteringDns VaultProvisioningState = "RegisteringDns"
	VaultProvisioningStateSucceeded      VaultProvisioningState = "Succeeded"
)

func PossibleValuesForVaultProvisioningState() []string {
	return []string{
		string(VaultProvisioningStateRegisteringDns),
		string(VaultProvisioningStateSucceeded),
	}
}
