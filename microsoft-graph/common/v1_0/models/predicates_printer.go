package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterOperationPredicate struct {
	DisplayName        *string
	HasPhysicalDevice  *bool
	Id                 *string
	IsAcceptingJobs    *bool
	IsShared           *bool
	LastSeenDateTime   *string
	Manufacturer       *string
	Model              *string
	ODataType          *string
	RegisteredDateTime *string
}

func (p PrinterOperationPredicate) Matches(input Printer) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.HasPhysicalDevice != nil && (input.HasPhysicalDevice == nil || *p.HasPhysicalDevice != *input.HasPhysicalDevice) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsAcceptingJobs != nil && (input.IsAcceptingJobs == nil || *p.IsAcceptingJobs != *input.IsAcceptingJobs) {
		return false
	}

	if p.IsShared != nil && (input.IsShared == nil || *p.IsShared != *input.IsShared) {
		return false
	}

	if p.LastSeenDateTime != nil && (input.LastSeenDateTime == nil || *p.LastSeenDateTime != *input.LastSeenDateTime) {
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

	if p.RegisteredDateTime != nil && (input.RegisteredDateTime == nil || *p.RegisteredDateTime != *input.RegisteredDateTime) {
		return false
	}

	return true
}
