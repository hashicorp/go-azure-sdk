package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkDisplayConfiguration struct {
	ConfiguredDisplays                *[]TeamworkConfiguredPeripheral     `json:"configuredDisplays,omitempty"`
	DisplayCount                      *int64                              `json:"displayCount,omitempty"`
	InBuiltDisplayScreenConfiguration *TeamworkDisplayScreenConfiguration `json:"inBuiltDisplayScreenConfiguration,omitempty"`
	IsContentDuplicationAllowed       *bool                               `json:"isContentDuplicationAllowed,omitempty"`
	IsDualDisplayModeEnabled          *bool                               `json:"isDualDisplayModeEnabled,omitempty"`
	ODataType                         *string                             `json:"@odata.type,omitempty"`
}
