package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamRenamedEventMessageDetailOperationPredicate struct {
	ODataType       *string
	TeamDisplayName *string
	TeamId          *string
}

func (p TeamRenamedEventMessageDetailOperationPredicate) Matches(input TeamRenamedEventMessageDetail) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TeamDisplayName != nil && (input.TeamDisplayName == nil || *p.TeamDisplayName != *input.TeamDisplayName) {
		return false
	}

	if p.TeamId != nil && (input.TeamId == nil || *p.TeamId != *input.TeamId) {
		return false
	}

	return true
}
