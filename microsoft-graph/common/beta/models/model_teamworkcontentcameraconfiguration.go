package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkContentCameraConfiguration struct {
	IsContentCameraInverted     *bool   `json:"isContentCameraInverted,omitempty"`
	IsContentCameraOptional     *bool   `json:"isContentCameraOptional,omitempty"`
	IsContentEnhancementEnabled *bool   `json:"isContentEnhancementEnabled,omitempty"`
	ODataType                   *string `json:"@odata.type,omitempty"`
}
