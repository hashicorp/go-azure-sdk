package regulatorycompliance

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegulatoryComplianceStandardProperties struct {
	FailedControls      *int64 `json:"failedControls,omitempty"`
	PassedControls      *int64 `json:"passedControls,omitempty"`
	SkippedControls     *int64 `json:"skippedControls,omitempty"`
	State               *State `json:"state,omitempty"`
	UnsupportedControls *int64 `json:"unsupportedControls,omitempty"`
}
