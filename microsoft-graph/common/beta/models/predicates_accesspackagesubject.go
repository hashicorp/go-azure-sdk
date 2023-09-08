package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageSubjectOperationPredicate struct {
	AltSecId                     *string
	CleanupScheduledDateTime     *string
	ConnectedOrganizationId      *string
	DisplayName                  *string
	Email                        *string
	Id                           *string
	ODataType                    *string
	ObjectId                     *string
	OnPremisesSecurityIdentifier *string
	PrincipalName                *string
	Type                         *string
}

func (p AccessPackageSubjectOperationPredicate) Matches(input AccessPackageSubject) bool {

	if p.AltSecId != nil && (input.AltSecId == nil || *p.AltSecId != *input.AltSecId) {
		return false
	}

	if p.CleanupScheduledDateTime != nil && (input.CleanupScheduledDateTime == nil || *p.CleanupScheduledDateTime != *input.CleanupScheduledDateTime) {
		return false
	}

	if p.ConnectedOrganizationId != nil && (input.ConnectedOrganizationId == nil || *p.ConnectedOrganizationId != *input.ConnectedOrganizationId) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Email != nil && (input.Email == nil || *p.Email != *input.Email) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ObjectId != nil && (input.ObjectId == nil || *p.ObjectId != *input.ObjectId) {
		return false
	}

	if p.OnPremisesSecurityIdentifier != nil && (input.OnPremisesSecurityIdentifier == nil || *p.OnPremisesSecurityIdentifier != *input.OnPremisesSecurityIdentifier) {
		return false
	}

	if p.PrincipalName != nil && (input.PrincipalName == nil || *p.PrincipalName != *input.PrincipalName) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}
