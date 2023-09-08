package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationConditionsApplicationsOperationPredicate struct {
	IncludeAllApplications *bool
	ODataType              *string
}

func (p AuthenticationConditionsApplicationsOperationPredicate) Matches(input AuthenticationConditionsApplications) bool {

	if p.IncludeAllApplications != nil && (input.IncludeAllApplications == nil || *p.IncludeAllApplications != *input.IncludeAllApplications) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
