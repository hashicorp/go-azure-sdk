package computerps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineScaleSetExtensionProperties struct {
	AutoUpgradeMinorVersion       *bool                    `json:"autoUpgradeMinorVersion,omitempty"`
	EnableAutomaticUpgrade        *bool                    `json:"enableAutomaticUpgrade,omitempty"`
	ForceUpdateTag                *string                  `json:"forceUpdateTag,omitempty"`
	ProtectedSettings             *interface{}             `json:"protectedSettings,omitempty"`
	ProtectedSettingsFromKeyVault *KeyVaultSecretReference `json:"protectedSettingsFromKeyVault,omitempty"`
	ProvisionAfterExtensions      *[]string                `json:"provisionAfterExtensions,omitempty"`
	ProvisioningState             *string                  `json:"provisioningState,omitempty"`
	Publisher                     *string                  `json:"publisher,omitempty"`
	Settings                      *interface{}             `json:"settings,omitempty"`
	SuppressFailures              *bool                    `json:"suppressFailures,omitempty"`
	Type                          *string                  `json:"type,omitempty"`
	TypeHandlerVersion            *string                  `json:"typeHandlerVersion,omitempty"`
}
