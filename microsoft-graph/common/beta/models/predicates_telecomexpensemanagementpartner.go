package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TelecomExpenseManagementPartnerOperationPredicate struct {
	AppAuthorized          *bool
	DisplayName            *string
	Enabled                *bool
	Id                     *string
	LastConnectionDateTime *string
	ODataType              *string
	Url                    *string
}

func (p TelecomExpenseManagementPartnerOperationPredicate) Matches(input TelecomExpenseManagementPartner) bool {

	if p.AppAuthorized != nil && (input.AppAuthorized == nil || *p.AppAuthorized != *input.AppAuthorized) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Enabled != nil && (input.Enabled == nil || *p.Enabled != *input.Enabled) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastConnectionDateTime != nil && (input.LastConnectionDateTime == nil || *p.LastConnectionDateTime != *input.LastConnectionDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Url != nil && (input.Url == nil || *p.Url != *input.Url) {
		return false
	}

	return true
}
