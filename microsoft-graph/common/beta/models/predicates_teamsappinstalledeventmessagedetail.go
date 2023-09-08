package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppInstalledEventMessageDetailOperationPredicate struct {
	ODataType           *string
	TeamsAppDisplayName *string
	TeamsAppId          *string
}

func (p TeamsAppInstalledEventMessageDetailOperationPredicate) Matches(input TeamsAppInstalledEventMessageDetail) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TeamsAppDisplayName != nil && (input.TeamsAppDisplayName == nil || *p.TeamsAppDisplayName != *input.TeamsAppDisplayName) {
		return false
	}

	if p.TeamsAppId != nil && (input.TeamsAppId == nil || *p.TeamsAppId != *input.TeamsAppId) {
		return false
	}

	return true
}
