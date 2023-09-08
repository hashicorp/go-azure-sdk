package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCDeviceImageOperationPredicate struct {
	DisplayName           *string
	ExpirationDate        *string
	Id                    *string
	LastModifiedDateTime  *string
	ODataType             *string
	OperatingSystem       *string
	OsBuildNumber         *string
	SourceImageResourceId *string
	Version               *string
}

func (p CloudPCDeviceImageOperationPredicate) Matches(input CloudPCDeviceImage) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.ExpirationDate != nil && (input.ExpirationDate == nil || *p.ExpirationDate != *input.ExpirationDate) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OperatingSystem != nil && (input.OperatingSystem == nil || *p.OperatingSystem != *input.OperatingSystem) {
		return false
	}

	if p.OsBuildNumber != nil && (input.OsBuildNumber == nil || *p.OsBuildNumber != *input.OsBuildNumber) {
		return false
	}

	if p.SourceImageResourceId != nil && (input.SourceImageResourceId == nil || *p.SourceImageResourceId != *input.SourceImageResourceId) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
