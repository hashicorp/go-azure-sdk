package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCOnPremisesConnectionStatusDetailsOperationPredicate struct {
	EndDateTime   *string
	ODataType     *string
	StartDateTime *string
}

func (p CloudPCOnPremisesConnectionStatusDetailsOperationPredicate) Matches(input CloudPCOnPremisesConnectionStatusDetails) bool {

	if p.EndDateTime != nil && (input.EndDateTime == nil || *p.EndDateTime != *input.EndDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.StartDateTime != nil && (input.StartDateTime == nil || *p.StartDateTime != *input.StartDateTime) {
		return false
	}

	return true
}
