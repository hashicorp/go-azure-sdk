package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpeditedWindowsQualityUpdateSettings struct {
	DaysUntilForcedReboot *int64  `json:"daysUntilForcedReboot,omitempty"`
	ODataType             *string `json:"@odata.type,omitempty"`
	QualityUpdateRelease  *string `json:"qualityUpdateRelease,omitempty"`
}
