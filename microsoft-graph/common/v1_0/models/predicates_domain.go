package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainOperationPredicate struct {
	AuthenticationType               *string
	AvailabilityStatus               *string
	Id                               *string
	IsAdminManaged                   *bool
	IsDefault                        *bool
	IsInitial                        *bool
	IsRoot                           *bool
	IsVerified                       *bool
	Manufacturer                     *string
	Model                            *string
	ODataType                        *string
	PasswordNotificationWindowInDays *int64
	PasswordValidityPeriodInDays     *int64
}

func (p DomainOperationPredicate) Matches(input Domain) bool {

	if p.AuthenticationType != nil && (input.AuthenticationType == nil || *p.AuthenticationType != *input.AuthenticationType) {
		return false
	}

	if p.AvailabilityStatus != nil && (input.AvailabilityStatus == nil || *p.AvailabilityStatus != *input.AvailabilityStatus) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsAdminManaged != nil && (input.IsAdminManaged == nil || *p.IsAdminManaged != *input.IsAdminManaged) {
		return false
	}

	if p.IsDefault != nil && (input.IsDefault == nil || *p.IsDefault != *input.IsDefault) {
		return false
	}

	if p.IsInitial != nil && (input.IsInitial == nil || *p.IsInitial != *input.IsInitial) {
		return false
	}

	if p.IsRoot != nil && (input.IsRoot == nil || *p.IsRoot != *input.IsRoot) {
		return false
	}

	if p.IsVerified != nil && (input.IsVerified == nil || *p.IsVerified != *input.IsVerified) {
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

	if p.PasswordNotificationWindowInDays != nil && (input.PasswordNotificationWindowInDays == nil || *p.PasswordNotificationWindowInDays != *input.PasswordNotificationWindowInDays) {
		return false
	}

	if p.PasswordValidityPeriodInDays != nil && (input.PasswordValidityPeriodInDays == nil || *p.PasswordValidityPeriodInDays != *input.PasswordValidityPeriodInDays) {
		return false
	}

	return true
}
