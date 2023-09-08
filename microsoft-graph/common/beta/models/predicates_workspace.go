package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceOperationPredicate struct {
	Building               *string
	Capacity               *int64
	DisplayName            *string
	EmailAddress           *string
	FloorLabel             *string
	FloorNumber            *int64
	Id                     *string
	IsWheelChairAccessible *bool
	Label                  *string
	Nickname               *string
	ODataType              *string
	Phone                  *string
}

func (p WorkspaceOperationPredicate) Matches(input Workspace) bool {

	if p.Building != nil && (input.Building == nil || *p.Building != *input.Building) {
		return false
	}

	if p.Capacity != nil && (input.Capacity == nil || *p.Capacity != *input.Capacity) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.EmailAddress != nil && (input.EmailAddress == nil || *p.EmailAddress != *input.EmailAddress) {
		return false
	}

	if p.FloorLabel != nil && (input.FloorLabel == nil || *p.FloorLabel != *input.FloorLabel) {
		return false
	}

	if p.FloorNumber != nil && (input.FloorNumber == nil || *p.FloorNumber != *input.FloorNumber) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsWheelChairAccessible != nil && (input.IsWheelChairAccessible == nil || *p.IsWheelChairAccessible != *input.IsWheelChairAccessible) {
		return false
	}

	if p.Label != nil && (input.Label == nil || *p.Label != *input.Label) {
		return false
	}

	if p.Nickname != nil && (input.Nickname == nil || *p.Nickname != *input.Nickname) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Phone != nil && (input.Phone == nil || *p.Phone != *input.Phone) {
		return false
	}

	return true
}
