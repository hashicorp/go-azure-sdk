package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkOnPremisesCalendarSyncConfigurationOperationPredicate struct {
	Domain         *string
	DomainUserName *string
	ODataType      *string
	SmtpAddress    *string
}

func (p TeamworkOnPremisesCalendarSyncConfigurationOperationPredicate) Matches(input TeamworkOnPremisesCalendarSyncConfiguration) bool {

	if p.Domain != nil && (input.Domain == nil || *p.Domain != *input.Domain) {
		return false
	}

	if p.DomainUserName != nil && (input.DomainUserName == nil || *p.DomainUserName != *input.DomainUserName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SmtpAddress != nil && (input.SmtpAddress == nil || *p.SmtpAddress != *input.SmtpAddress) {
		return false
	}

	return true
}
