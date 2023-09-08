package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceLogCollectionResponseOperationPredicate struct {
	EnrolledByUser               *string
	ExpirationDateTimeUTC        *string
	Id                           *string
	InitiatedByUserPrincipalName *string
	ManagedDeviceId              *string
	ODataType                    *string
	ReceivedDateTimeUTC          *string
	RequestedDateTimeUTC         *string
}

func (p DeviceLogCollectionResponseOperationPredicate) Matches(input DeviceLogCollectionResponse) bool {

	if p.EnrolledByUser != nil && (input.EnrolledByUser == nil || *p.EnrolledByUser != *input.EnrolledByUser) {
		return false
	}

	if p.ExpirationDateTimeUTC != nil && (input.ExpirationDateTimeUTC == nil || *p.ExpirationDateTimeUTC != *input.ExpirationDateTimeUTC) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.InitiatedByUserPrincipalName != nil && (input.InitiatedByUserPrincipalName == nil || *p.InitiatedByUserPrincipalName != *input.InitiatedByUserPrincipalName) {
		return false
	}

	if p.ManagedDeviceId != nil && (input.ManagedDeviceId == nil || *p.ManagedDeviceId != *input.ManagedDeviceId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ReceivedDateTimeUTC != nil && (input.ReceivedDateTimeUTC == nil || *p.ReceivedDateTimeUTC != *input.ReceivedDateTimeUTC) {
		return false
	}

	if p.RequestedDateTimeUTC != nil && (input.RequestedDateTimeUTC == nil || *p.RequestedDateTimeUTC != *input.RequestedDateTimeUTC) {
		return false
	}

	return true
}
