package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommandOperationPredicate struct {
	AppServiceName    *string
	Error             *string
	Id                *string
	ODataType         *string
	PackageFamilyName *string
	PermissionTicket  *string
	PostBackUri       *string
	Status            *string
	Type              *string
}

func (p CommandOperationPredicate) Matches(input Command) bool {

	if p.AppServiceName != nil && (input.AppServiceName == nil || *p.AppServiceName != *input.AppServiceName) {
		return false
	}

	if p.Error != nil && (input.Error == nil || *p.Error != *input.Error) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PackageFamilyName != nil && (input.PackageFamilyName == nil || *p.PackageFamilyName != *input.PackageFamilyName) {
		return false
	}

	if p.PermissionTicket != nil && (input.PermissionTicket == nil || *p.PermissionTicket != *input.PermissionTicket) {
		return false
	}

	if p.PostBackUri != nil && (input.PostBackUri == nil || *p.PostBackUri != *input.PostBackUri) {
		return false
	}

	if p.Status != nil && (input.Status == nil || *p.Status != *input.Status) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}
