package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosAzureAdSingleSignOnExtension struct {
	BundleIdAccessControlList *[]string            `json:"bundleIdAccessControlList,omitempty"`
	Configurations            *[]KeyTypedValuePair `json:"configurations,omitempty"`
	EnableSharedDeviceMode    *bool                `json:"enableSharedDeviceMode,omitempty"`
	ODataType                 *string              `json:"@odata.type,omitempty"`
}
