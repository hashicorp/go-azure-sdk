package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectSiteOperationPredicate struct {
	ConditionalAccessProtectionLevelId *string
	Name                               *string
	ODataType                          *string
}

func (p ProtectSiteOperationPredicate) Matches(input ProtectSite) bool {

	if p.ConditionalAccessProtectionLevelId != nil && (input.ConditionalAccessProtectionLevelId == nil || *p.ConditionalAccessProtectionLevelId != *input.ConditionalAccessProtectionLevelId) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
