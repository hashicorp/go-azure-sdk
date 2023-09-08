package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComplianceManagementPartnerOperationPredicate struct {
	AndroidOnboarded      *bool
	DisplayName           *string
	Id                    *string
	IosOnboarded          *bool
	LastHeartbeatDateTime *string
	MacOsOnboarded        *bool
	ODataType             *string
}

func (p ComplianceManagementPartnerOperationPredicate) Matches(input ComplianceManagementPartner) bool {

	if p.AndroidOnboarded != nil && (input.AndroidOnboarded == nil || *p.AndroidOnboarded != *input.AndroidOnboarded) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IosOnboarded != nil && (input.IosOnboarded == nil || *p.IosOnboarded != *input.IosOnboarded) {
		return false
	}

	if p.LastHeartbeatDateTime != nil && (input.LastHeartbeatDateTime == nil || *p.LastHeartbeatDateTime != *input.LastHeartbeatDateTime) {
		return false
	}

	if p.MacOsOnboarded != nil && (input.MacOsOnboarded == nil || *p.MacOsOnboarded != *input.MacOsOnboarded) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
