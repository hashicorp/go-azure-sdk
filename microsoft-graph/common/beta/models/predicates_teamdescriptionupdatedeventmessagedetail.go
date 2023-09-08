package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamDescriptionUpdatedEventMessageDetailOperationPredicate struct {
	ODataType       *string
	TeamDescription *string
	TeamId          *string
}

func (p TeamDescriptionUpdatedEventMessageDetailOperationPredicate) Matches(input TeamDescriptionUpdatedEventMessageDetail) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TeamDescription != nil && (input.TeamDescription == nil || *p.TeamDescription != *input.TeamDescription) {
		return false
	}

	if p.TeamId != nil && (input.TeamId == nil || *p.TeamId != *input.TeamId) {
		return false
	}

	return true
}
