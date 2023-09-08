package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSAssociatedDomainsItemOperationPredicate struct {
	ApplicationIdentifier  *string
	DirectDownloadsEnabled *bool
	ODataType              *string
}

func (p MacOSAssociatedDomainsItemOperationPredicate) Matches(input MacOSAssociatedDomainsItem) bool {

	if p.ApplicationIdentifier != nil && (input.ApplicationIdentifier == nil || *p.ApplicationIdentifier != *input.ApplicationIdentifier) {
		return false
	}

	if p.DirectDownloadsEnabled != nil && (input.DirectDownloadsEnabled == nil || *p.DirectDownloadsEnabled != *input.DirectDownloadsEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
