package deponboardingsettingimportedappledeviceidentity

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDepOnboardingSettingImportedAppleDeviceIdentityImportAppleDeviceIdentityListsRequest struct {
	ImportedAppleDeviceIdentities     *[]ImportedAppleDeviceIdentity `json:"importedAppleDeviceIdentities,omitempty"`
	OverwriteImportedDeviceIdentities *bool                          `json:"overwriteImportedDeviceIdentities,omitempty"`
}
