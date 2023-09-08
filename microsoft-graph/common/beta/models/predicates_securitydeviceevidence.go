package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDeviceEvidenceOperationPredicate struct {
	AzureAdDeviceId          *string
	CreatedDateTime          *string
	DeviceDnsName            *string
	FirstSeenDateTime        *string
	MdeDeviceId              *string
	ODataType                *string
	OsBuild                  *int64
	OsPlatform               *string
	RbacGroupId              *int64
	RbacGroupName            *string
	RemediationStatusDetails *string
	Version                  *string
}

func (p SecurityDeviceEvidenceOperationPredicate) Matches(input SecurityDeviceEvidence) bool {

	if p.AzureAdDeviceId != nil && (input.AzureAdDeviceId == nil || *p.AzureAdDeviceId != *input.AzureAdDeviceId) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DeviceDnsName != nil && (input.DeviceDnsName == nil || *p.DeviceDnsName != *input.DeviceDnsName) {
		return false
	}

	if p.FirstSeenDateTime != nil && (input.FirstSeenDateTime == nil || *p.FirstSeenDateTime != *input.FirstSeenDateTime) {
		return false
	}

	if p.MdeDeviceId != nil && (input.MdeDeviceId == nil || *p.MdeDeviceId != *input.MdeDeviceId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OsBuild != nil && (input.OsBuild == nil || *p.OsBuild != *input.OsBuild) {
		return false
	}

	if p.OsPlatform != nil && (input.OsPlatform == nil || *p.OsPlatform != *input.OsPlatform) {
		return false
	}

	if p.RbacGroupId != nil && (input.RbacGroupId == nil || *p.RbacGroupId != *input.RbacGroupId) {
		return false
	}

	if p.RbacGroupName != nil && (input.RbacGroupName == nil || *p.RbacGroupName != *input.RbacGroupName) {
		return false
	}

	if p.RemediationStatusDetails != nil && (input.RemediationStatusDetails == nil || *p.RemediationStatusDetails != *input.RemediationStatusDetails) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
