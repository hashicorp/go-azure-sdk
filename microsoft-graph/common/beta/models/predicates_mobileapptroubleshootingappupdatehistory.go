package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppTroubleshootingAppUpdateHistoryOperationPredicate struct {
	ODataType          *string
	OccurrenceDateTime *string
}

func (p MobileAppTroubleshootingAppUpdateHistoryOperationPredicate) Matches(input MobileAppTroubleshootingAppUpdateHistory) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OccurrenceDateTime != nil && (input.OccurrenceDateTime == nil || *p.OccurrenceDateTime != *input.OccurrenceDateTime) {
		return false
	}

	return true
}
