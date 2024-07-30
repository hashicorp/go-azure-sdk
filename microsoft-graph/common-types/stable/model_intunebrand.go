package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntuneBrand struct {
	ContactITEmailAddress     *string      `json:"contactITEmailAddress,omitempty"`
	ContactITName             *string      `json:"contactITName,omitempty"`
	ContactITNotes            *string      `json:"contactITNotes,omitempty"`
	ContactITPhoneNumber      *string      `json:"contactITPhoneNumber,omitempty"`
	DarkBackgroundLogo        *MimeContent `json:"darkBackgroundLogo,omitempty"`
	DisplayName               *string      `json:"displayName,omitempty"`
	LightBackgroundLogo       *MimeContent `json:"lightBackgroundLogo,omitempty"`
	ODataType                 *string      `json:"@odata.type,omitempty"`
	OnlineSupportSiteName     *string      `json:"onlineSupportSiteName,omitempty"`
	OnlineSupportSiteUrl      *string      `json:"onlineSupportSiteUrl,omitempty"`
	PrivacyUrl                *string      `json:"privacyUrl,omitempty"`
	ShowDisplayNameNextToLogo *bool        `json:"showDisplayNameNextToLogo,omitempty"`
	ShowLogo                  *bool        `json:"showLogo,omitempty"`
	ShowNameNextToLogo        *bool        `json:"showNameNextToLogo,omitempty"`
	ThemeColor                *RgbColor    `json:"themeColor,omitempty"`
}
