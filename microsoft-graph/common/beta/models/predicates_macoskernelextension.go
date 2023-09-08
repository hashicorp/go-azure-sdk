package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSKernelExtensionOperationPredicate struct {
	BundleId       *string
	ODataType      *string
	TeamIdentifier *string
}

func (p MacOSKernelExtensionOperationPredicate) Matches(input MacOSKernelExtension) bool {

	if p.BundleId != nil && (input.BundleId == nil || *p.BundleId != *input.BundleId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TeamIdentifier != nil && (input.TeamIdentifier == nil || *p.TeamIdentifier != *input.TeamIdentifier) {
		return false
	}

	return true
}
