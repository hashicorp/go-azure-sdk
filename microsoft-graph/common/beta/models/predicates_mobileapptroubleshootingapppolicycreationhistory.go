package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppTroubleshootingAppPolicyCreationHistoryOperationPredicate struct {
	ErrorCode          *string
	ODataType          *string
	OccurrenceDateTime *string
}

func (p MobileAppTroubleshootingAppPolicyCreationHistoryOperationPredicate) Matches(input MobileAppTroubleshootingAppPolicyCreationHistory) bool {

	if p.ErrorCode != nil && (input.ErrorCode == nil || *p.ErrorCode != *input.ErrorCode) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OccurrenceDateTime != nil && (input.OccurrenceDateTime == nil || *p.OccurrenceDateTime != *input.OccurrenceDateTime) {
		return false
	}

	return true
}
