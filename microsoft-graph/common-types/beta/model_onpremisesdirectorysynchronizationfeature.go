package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesDirectorySynchronizationFeature struct {
	BlockCloudObjectTakeoverThroughHardMatchEnabled  *bool   `json:"blockCloudObjectTakeoverThroughHardMatchEnabled,omitempty"`
	BlockSoftMatchEnabled                            *bool   `json:"blockSoftMatchEnabled,omitempty"`
	BypassDirSyncOverridesEnabled                    *bool   `json:"bypassDirSyncOverridesEnabled,omitempty"`
	CloudPasswordPolicyForPasswordSyncedUsersEnabled *bool   `json:"cloudPasswordPolicyForPasswordSyncedUsersEnabled,omitempty"`
	ConcurrentCredentialUpdateEnabled                *bool   `json:"concurrentCredentialUpdateEnabled,omitempty"`
	ConcurrentOrgIdProvisioningEnabled               *bool   `json:"concurrentOrgIdProvisioningEnabled,omitempty"`
	DeviceWritebackEnabled                           *bool   `json:"deviceWritebackEnabled,omitempty"`
	DirectoryExtensionsEnabled                       *bool   `json:"directoryExtensionsEnabled,omitempty"`
	FopeConflictResolutionEnabled                    *bool   `json:"fopeConflictResolutionEnabled,omitempty"`
	GroupWriteBackEnabled                            *bool   `json:"groupWriteBackEnabled,omitempty"`
	ODataType                                        *string `json:"@odata.type,omitempty"`
	PasswordSyncEnabled                              *bool   `json:"passwordSyncEnabled,omitempty"`
	PasswordWritebackEnabled                         *bool   `json:"passwordWritebackEnabled,omitempty"`
	QuarantineUponProxyAddressesConflictEnabled      *bool   `json:"quarantineUponProxyAddressesConflictEnabled,omitempty"`
	QuarantineUponUpnConflictEnabled                 *bool   `json:"quarantineUponUpnConflictEnabled,omitempty"`
	SoftMatchOnUpnEnabled                            *bool   `json:"softMatchOnUpnEnabled,omitempty"`
	SynchronizeUpnForManagedUsersEnabled             *bool   `json:"synchronizeUpnForManagedUsersEnabled,omitempty"`
	UnifiedGroupWritebackEnabled                     *bool   `json:"unifiedGroupWritebackEnabled,omitempty"`
	UserForcePasswordChangeOnLogonEnabled            *bool   `json:"userForcePasswordChangeOnLogonEnabled,omitempty"`
	UserWritebackEnabled                             *bool   `json:"userWritebackEnabled,omitempty"`
}
