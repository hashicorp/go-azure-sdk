package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BuildVersionDetailsOperationPredicate struct {
	BuildNumber         *int64
	MajorVersion        *int64
	MinorVersion        *int64
	ODataType           *string
	UpdateBuildRevision *int64
}

func (p BuildVersionDetailsOperationPredicate) Matches(input BuildVersionDetails) bool {

	if p.BuildNumber != nil && (input.BuildNumber == nil || *p.BuildNumber != *input.BuildNumber) {
		return false
	}

	if p.MajorVersion != nil && (input.MajorVersion == nil || *p.MajorVersion != *input.MajorVersion) {
		return false
	}

	if p.MinorVersion != nil && (input.MinorVersion == nil || *p.MinorVersion != *input.MinorVersion) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UpdateBuildRevision != nil && (input.UpdateBuildRevision == nil || *p.UpdateBuildRevision != *input.UpdateBuildRevision) {
		return false
	}

	return true
}
