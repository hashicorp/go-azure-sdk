package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftManagedDesktop struct {
	ODataType *string                      `json:"@odata.type,omitempty"`
	Profile   *string                      `json:"profile,omitempty"`
	Type      *MicrosoftManagedDesktopType `json:"type,omitempty"`
}
