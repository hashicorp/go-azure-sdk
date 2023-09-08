package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSLobChildAppOperationPredicate struct {
	BuildNumber   *string
	BundleId      *string
	ODataType     *string
	VersionNumber *string
}

func (p MacOSLobChildAppOperationPredicate) Matches(input MacOSLobChildApp) bool {

	if p.BuildNumber != nil && (input.BuildNumber == nil || *p.BuildNumber != *input.BuildNumber) {
		return false
	}

	if p.BundleId != nil && (input.BundleId == nil || *p.BundleId != *input.BundleId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.VersionNumber != nil && (input.VersionNumber == nil || *p.VersionNumber != *input.VersionNumber) {
		return false
	}

	return true
}
