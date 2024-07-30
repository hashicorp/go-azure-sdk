package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPrivacyDataAccessControlItem struct {
	AccessLevel          *WindowsPrivacyDataAccessControlItemAccessLevel  `json:"accessLevel,omitempty"`
	AppDisplayName       *string                                          `json:"appDisplayName,omitempty"`
	AppPackageFamilyName *string                                          `json:"appPackageFamilyName,omitempty"`
	DataCategory         *WindowsPrivacyDataAccessControlItemDataCategory `json:"dataCategory,omitempty"`
	Id                   *string                                          `json:"id,omitempty"`
	ODataType            *string                                          `json:"@odata.type,omitempty"`
}
