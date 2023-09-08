package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInOperationPredicate struct {
	AppDisplayName      *string
	AppId               *string
	ClientAppUsed       *string
	CorrelationId       *string
	CreatedDateTime     *string
	Id                  *string
	IpAddress           *string
	IsInteractive       *bool
	ODataType           *string
	ResourceDisplayName *string
	ResourceId          *string
	UserDisplayName     *string
	UserId              *string
	UserPrincipalName   *string
}

func (p SignInOperationPredicate) Matches(input SignIn) bool {

	if p.AppDisplayName != nil && (input.AppDisplayName == nil || *p.AppDisplayName != *input.AppDisplayName) {
		return false
	}

	if p.AppId != nil && (input.AppId == nil || *p.AppId != *input.AppId) {
		return false
	}

	if p.ClientAppUsed != nil && (input.ClientAppUsed == nil || *p.ClientAppUsed != *input.ClientAppUsed) {
		return false
	}

	if p.CorrelationId != nil && (input.CorrelationId == nil || *p.CorrelationId != *input.CorrelationId) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IpAddress != nil && (input.IpAddress == nil || *p.IpAddress != *input.IpAddress) {
		return false
	}

	if p.IsInteractive != nil && (input.IsInteractive == nil || *p.IsInteractive != *input.IsInteractive) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ResourceDisplayName != nil && (input.ResourceDisplayName == nil || *p.ResourceDisplayName != *input.ResourceDisplayName) {
		return false
	}

	if p.ResourceId != nil && (input.ResourceId == nil || *p.ResourceId != *input.ResourceId) {
		return false
	}

	if p.UserDisplayName != nil && (input.UserDisplayName == nil || *p.UserDisplayName != *input.UserDisplayName) {
		return false
	}

	if p.UserId != nil && (input.UserId == nil || *p.UserId != *input.UserId) {
		return false
	}

	if p.UserPrincipalName != nil && (input.UserPrincipalName == nil || *p.UserPrincipalName != *input.UserPrincipalName) {
		return false
	}

	return true
}
