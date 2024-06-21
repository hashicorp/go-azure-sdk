package galleryapplications

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GalleryApplicationProperties struct {
	Description         *string              `json:"description,omitempty"`
	EndOfLifeDate       *string              `json:"endOfLifeDate,omitempty"`
	Eula                *string              `json:"eula,omitempty"`
	PrivacyStatementUri *string              `json:"privacyStatementUri,omitempty"`
	ReleaseNoteUri      *string              `json:"releaseNoteUri,omitempty"`
	SupportedOSType     OperatingSystemTypes `json:"supportedOSType"`
}
