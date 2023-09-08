package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsWorkloadStatusOperationPredicate struct {
	DisplayName        *string
	ODataType          *string
	OffboardedDateTime *string
	OnboardedDateTime  *string
}

func (p ManagedTenantsWorkloadStatusOperationPredicate) Matches(input ManagedTenantsWorkloadStatus) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OffboardedDateTime != nil && (input.OffboardedDateTime == nil || *p.OffboardedDateTime != *input.OffboardedDateTime) {
		return false
	}

	if p.OnboardedDateTime != nil && (input.OnboardedDateTime == nil || *p.OnboardedDateTime != *input.OnboardedDateTime) {
		return false
	}

	return true
}
