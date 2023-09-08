package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenShiftItemOperationPredicate struct {
	DisplayName   *string
	EndDateTime   *string
	Notes         *string
	ODataType     *string
	OpenSlotCount *int64
	StartDateTime *string
}

func (p OpenShiftItemOperationPredicate) Matches(input OpenShiftItem) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.EndDateTime != nil && (input.EndDateTime == nil || *p.EndDateTime != *input.EndDateTime) {
		return false
	}

	if p.Notes != nil && (input.Notes == nil || *p.Notes != *input.Notes) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OpenSlotCount != nil && (input.OpenSlotCount == nil || *p.OpenSlotCount != *input.OpenSlotCount) {
		return false
	}

	if p.StartDateTime != nil && (input.StartDateTime == nil || *p.StartDateTime != *input.StartDateTime) {
		return false
	}

	return true
}
