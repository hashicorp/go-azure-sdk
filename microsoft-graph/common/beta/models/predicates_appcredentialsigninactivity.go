package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppCredentialSignInActivityOperationPredicate struct {
	AppId                    *string
	AppObjectId              *string
	CreatedDateTime          *string
	ExpirationDateTime       *string
	Id                       *string
	KeyId                    *string
	ODataType                *string
	ResourceId               *string
	ServicePrincipalObjectId *string
}

func (p AppCredentialSignInActivityOperationPredicate) Matches(input AppCredentialSignInActivity) bool {

	if p.AppId != nil && (input.AppId == nil || *p.AppId != *input.AppId) {
		return false
	}

	if p.AppObjectId != nil && (input.AppObjectId == nil || *p.AppObjectId != *input.AppObjectId) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.ExpirationDateTime != nil && (input.ExpirationDateTime == nil || *p.ExpirationDateTime != *input.ExpirationDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.KeyId != nil && (input.KeyId == nil || *p.KeyId != *input.KeyId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ResourceId != nil && (input.ResourceId == nil || *p.ResourceId != *input.ResourceId) {
		return false
	}

	if p.ServicePrincipalObjectId != nil && (input.ServicePrincipalObjectId == nil || *p.ServicePrincipalObjectId != *input.ServicePrincipalObjectId) {
		return false
	}

	return true
}
