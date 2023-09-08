package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DlpEvaluationWindowsDevicesInput struct {
	ContentProperties        *ContentProperties         `json:"contentProperties,omitempty"`
	CurrentLabel             *CurrentLabel              `json:"currentLabel,omitempty"`
	DiscoveredSensitiveTypes *[]DiscoveredSensitiveType `json:"discoveredSensitiveTypes,omitempty"`
	ODataType                *string                    `json:"@odata.type,omitempty"`
	SharedBy                 *string                    `json:"sharedBy,omitempty"`
}
