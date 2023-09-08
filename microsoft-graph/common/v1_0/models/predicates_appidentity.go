package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppIdentityOperationPredicate struct {
	AppId                *string
	DisplayName          *string
	ODataType            *string
	ServicePrincipalId   *string
	ServicePrincipalName *string
}

func (p AppIdentityOperationPredicate) Matches(input AppIdentity) bool {

	if p.AppId != nil && (input.AppId == nil || *p.AppId != *input.AppId) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ServicePrincipalId != nil && (input.ServicePrincipalId == nil || *p.ServicePrincipalId != *input.ServicePrincipalId) {
		return false
	}

	if p.ServicePrincipalName != nil && (input.ServicePrincipalName == nil || *p.ServicePrincipalName != *input.ServicePrincipalName) {
		return false
	}

	return true
}
