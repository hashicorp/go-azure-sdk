package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCAuditActorOperationPredicate struct {
	ApplicationDisplayName *string
	ApplicationId          *string
	IpAddress              *string
	ODataType              *string
	RemoteTenantId         *string
	RemoteUserId           *string
	ServicePrincipalName   *string
	UserId                 *string
	UserPrincipalName      *string
}

func (p CloudPCAuditActorOperationPredicate) Matches(input CloudPCAuditActor) bool {

	if p.ApplicationDisplayName != nil && (input.ApplicationDisplayName == nil || *p.ApplicationDisplayName != *input.ApplicationDisplayName) {
		return false
	}

	if p.ApplicationId != nil && (input.ApplicationId == nil || *p.ApplicationId != *input.ApplicationId) {
		return false
	}

	if p.IpAddress != nil && (input.IpAddress == nil || *p.IpAddress != *input.IpAddress) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RemoteTenantId != nil && (input.RemoteTenantId == nil || *p.RemoteTenantId != *input.RemoteTenantId) {
		return false
	}

	if p.RemoteUserId != nil && (input.RemoteUserId == nil || *p.RemoteUserId != *input.RemoteUserId) {
		return false
	}

	if p.ServicePrincipalName != nil && (input.ServicePrincipalName == nil || *p.ServicePrincipalName != *input.ServicePrincipalName) {
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
