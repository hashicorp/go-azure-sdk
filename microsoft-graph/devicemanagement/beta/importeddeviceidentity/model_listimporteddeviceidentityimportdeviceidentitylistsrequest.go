package importeddeviceidentity

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListImportedDeviceIdentityImportDeviceIdentityListsRequest struct {
	ImportedDeviceIdentities          *[]ImportedDeviceIdentity `json:"importedDeviceIdentities,omitempty"`
	OverwriteImportedDeviceIdentities *bool                     `json:"overwriteImportedDeviceIdentities,omitempty"`
}
