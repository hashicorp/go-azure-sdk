package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesAccidentalDeletionPreventionOperationPredicate struct {
	AlertThreshold *int64
	ODataType      *string
}

func (p OnPremisesAccidentalDeletionPreventionOperationPredicate) Matches(input OnPremisesAccidentalDeletionPrevention) bool {

	if p.AlertThreshold != nil && (input.AlertThreshold == nil || *p.AlertThreshold != *input.AlertThreshold) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
