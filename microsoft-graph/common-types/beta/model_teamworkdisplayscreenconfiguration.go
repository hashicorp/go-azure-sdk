package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkDisplayScreenConfiguration struct {
	BacklightBrightness   *int64  `json:"backlightBrightness,omitempty"`
	BacklightTimeout      *string `json:"backlightTimeout,omitempty"`
	IsHighContrastEnabled *bool   `json:"isHighContrastEnabled,omitempty"`
	IsScreensaverEnabled  *bool   `json:"isScreensaverEnabled,omitempty"`
	ODataType             *string `json:"@odata.type,omitempty"`
	ScreensaverTimeout    *string `json:"screensaverTimeout,omitempty"`
}
