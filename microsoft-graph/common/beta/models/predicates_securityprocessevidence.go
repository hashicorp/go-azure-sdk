package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityProcessEvidenceOperationPredicate struct {
	CreatedDateTime               *string
	MdeDeviceId                   *string
	ODataType                     *string
	ParentProcessCreationDateTime *string
	ParentProcessId               *int64
	ProcessCommandLine            *string
	ProcessCreationDateTime       *string
	ProcessId                     *int64
	RemediationStatusDetails      *string
}

func (p SecurityProcessEvidenceOperationPredicate) Matches(input SecurityProcessEvidence) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.MdeDeviceId != nil && (input.MdeDeviceId == nil || *p.MdeDeviceId != *input.MdeDeviceId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ParentProcessCreationDateTime != nil && (input.ParentProcessCreationDateTime == nil || *p.ParentProcessCreationDateTime != *input.ParentProcessCreationDateTime) {
		return false
	}

	if p.ParentProcessId != nil && (input.ParentProcessId == nil || *p.ParentProcessId != *input.ParentProcessId) {
		return false
	}

	if p.ProcessCommandLine != nil && (input.ProcessCommandLine == nil || *p.ProcessCommandLine != *input.ProcessCommandLine) {
		return false
	}

	if p.ProcessCreationDateTime != nil && (input.ProcessCreationDateTime == nil || *p.ProcessCreationDateTime != *input.ProcessCreationDateTime) {
		return false
	}

	if p.ProcessId != nil && (input.ProcessId == nil || *p.ProcessId != *input.ProcessId) {
		return false
	}

	if p.RemediationStatusDetails != nil && (input.RemediationStatusDetails == nil || *p.RemediationStatusDetails != *input.RemediationStatusDetails) {
		return false
	}

	return true
}
