package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterShareOperationPredicate struct {
	AllowAllUsers   *bool
	CreatedDateTime *string
	DisplayName     *string
	Id              *string
	IsAcceptingJobs *bool
	Manufacturer    *string
	Model           *string
	ODataType       *string
}

func (p PrinterShareOperationPredicate) Matches(input PrinterShare) bool {

	if p.AllowAllUsers != nil && (input.AllowAllUsers == nil || *p.AllowAllUsers != *input.AllowAllUsers) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsAcceptingJobs != nil && (input.IsAcceptingJobs == nil || *p.IsAcceptingJobs != *input.IsAcceptingJobs) {
		return false
	}

	if p.Manufacturer != nil && (input.Manufacturer == nil || *p.Manufacturer != *input.Manufacturer) {
		return false
	}

	if p.Model != nil && (input.Model == nil || *p.Model != *input.Model) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
