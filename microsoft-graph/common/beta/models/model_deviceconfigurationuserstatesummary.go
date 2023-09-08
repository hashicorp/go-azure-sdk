package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationUserStateSummary struct {
	CompliantUserCount     *int64  `json:"compliantUserCount,omitempty"`
	ConflictUserCount      *int64  `json:"conflictUserCount,omitempty"`
	ErrorUserCount         *int64  `json:"errorUserCount,omitempty"`
	Id                     *string `json:"id,omitempty"`
	NonCompliantUserCount  *int64  `json:"nonCompliantUserCount,omitempty"`
	NotApplicableUserCount *int64  `json:"notApplicableUserCount,omitempty"`
	ODataType              *string `json:"@odata.type,omitempty"`
	RemediatedUserCount    *int64  `json:"remediatedUserCount,omitempty"`
	UnknownUserCount       *int64  `json:"unknownUserCount,omitempty"`
}
