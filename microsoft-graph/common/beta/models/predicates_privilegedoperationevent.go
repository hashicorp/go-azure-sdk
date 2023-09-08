package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedOperationEventOperationPredicate struct {
	AdditionalInformation *string
	CreationDateTime      *string
	ExpirationDateTime    *string
	Id                    *string
	ODataType             *string
	ReferenceKey          *string
	ReferenceSystem       *string
	RequestType           *string
	RequestorId           *string
	RequestorName         *string
	RoleId                *string
	RoleName              *string
	TenantId              *string
	UserId                *string
	UserMail              *string
	UserName              *string
}

func (p PrivilegedOperationEventOperationPredicate) Matches(input PrivilegedOperationEvent) bool {

	if p.AdditionalInformation != nil && (input.AdditionalInformation == nil || *p.AdditionalInformation != *input.AdditionalInformation) {
		return false
	}

	if p.CreationDateTime != nil && (input.CreationDateTime == nil || *p.CreationDateTime != *input.CreationDateTime) {
		return false
	}

	if p.ExpirationDateTime != nil && (input.ExpirationDateTime == nil || *p.ExpirationDateTime != *input.ExpirationDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ReferenceKey != nil && (input.ReferenceKey == nil || *p.ReferenceKey != *input.ReferenceKey) {
		return false
	}

	if p.ReferenceSystem != nil && (input.ReferenceSystem == nil || *p.ReferenceSystem != *input.ReferenceSystem) {
		return false
	}

	if p.RequestType != nil && (input.RequestType == nil || *p.RequestType != *input.RequestType) {
		return false
	}

	if p.RequestorId != nil && (input.RequestorId == nil || *p.RequestorId != *input.RequestorId) {
		return false
	}

	if p.RequestorName != nil && (input.RequestorName == nil || *p.RequestorName != *input.RequestorName) {
		return false
	}

	if p.RoleId != nil && (input.RoleId == nil || *p.RoleId != *input.RoleId) {
		return false
	}

	if p.RoleName != nil && (input.RoleName == nil || *p.RoleName != *input.RoleName) {
		return false
	}

	if p.TenantId != nil && (input.TenantId == nil || *p.TenantId != *input.TenantId) {
		return false
	}

	if p.UserId != nil && (input.UserId == nil || *p.UserId != *input.UserId) {
		return false
	}

	if p.UserMail != nil && (input.UserMail == nil || *p.UserMail != *input.UserMail) {
		return false
	}

	if p.UserName != nil && (input.UserName == nil || *p.UserName != *input.UserName) {
		return false
	}

	return true
}
